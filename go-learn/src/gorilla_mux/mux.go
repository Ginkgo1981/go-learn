package gorilla_mux

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

func SimpleMuxServer()  {
	r := mux.NewRouter()
	r.HandleFunc("/", DefaultHandler)
	r.HandleFunc("/products/{productId}", GetProductHandler).Methods("GET")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", GetArticleHandler)
	http.Handle("/", r)

	server := &http.Server{
		Addr: ":3000",
		Handler: r,
	}
	fmt.Println("=== start server ====")
	log.Fatal(server.ListenAndServe())
	//log.Fatal(http.ListenAndServe(":8080", r))
}

type Product struct {
	Name string
	Quantity int
	Price int
}

var products = []Product{
	Product{"product1", 100, 100},
}

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var p Product
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &p)
	if err != nil {
		log.Fatal(err)
	}
	products = append(products, p)
	fmt.Println("=== vars:", vars)
	fmt.Println("=== product: ", p)
	j, err := json.Marshal(p)
	w.Write(j)
	//fmt.Fprint(w,j)
}

func GetProductHandler(w http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)
	productId := vars["productId"]
	fmt.Fprintln(w, string(productId))
}


func GetProducts(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "==== GetProducts ====")
	j, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(j)
}


func DefaultHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "==== DefaultHandler ====")

}

func GetArticleHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	category := vars["category"]
	id := vars["id"]
	fmt.Fprintln(w, "==== GetArticleHandler ====", category, id)
}
