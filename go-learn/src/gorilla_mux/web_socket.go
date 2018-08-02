package gorilla_mux

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func StartWebsocket()  {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		nextRequestID := func() string{
			return fmt.Sprintf("%d", time.Now().UnixNano())
		}

		requestID := r.Header.Get("X-Request-Id")

		if requestID == "" {
			requestID = nextRequestID()
		}

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
				//continue
				return
			}
			fmt.Printf("%s send: %s requestId: %s \n", conn.RemoteAddr(), string(msg), string(requestID))
			if err := conn.WriteMessage(msgType, msg); err != nil {
				log.Fatal(err)
				//continue
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "websockets.html")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}