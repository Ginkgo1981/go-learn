package simple_server

import (
	"flag"
	"log"
	"os"
	"net/http"
	"fmt"
	"time"
	"os/signal"
	"context"
	"sync/atomic"
)

//https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7
//A simple golang web server with basic logging, tracing, health check, graceful shutdown and zero dependencies

type key int
const requestIDKey key = 0
var listenAddr string

var healthy int32

func StartServer() {
	flag.StringVar(&listenAddr, "listen-addr", ":3000", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http:", log.LstdFlags)

	logger.Println("Server is stating...")

	router := http.NewServeMux()
	router.Handle("/", index())
	router.Handle("/healthz", heathz())

	nextRequestID := func() string{
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr: listenAddr,
		Handler: tracing(nextRequestID)(logging(logger)(router)),
		ErrorLog:logger,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)


	go func() {
		<-quit
		logger.Println("server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at: ", listenAddr)

	atomic.StoreInt32(&healthy, 1)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s:%v\n", listenAddr, err)
	}

	<-done
	logger.Println("server Stopped")
}
func logging(logger *log.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
				fmt.Println("Request-Id", requestID)
			}()
			next.ServeHTTP(w, r)
		})
	}
}


func tracing(nextRequestID func() string) func(handler http.Handler) http.Handler{

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			fmt.Println("Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		requestID, _ := r.Context().Value(requestIDKey).(string)
		fmt.Println("Request-Id", requestID)
		fmt.Fprintln(w, "Hello, index page")
	})
}

func heathz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Heathz")
	})
}

