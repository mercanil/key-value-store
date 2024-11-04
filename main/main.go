package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mercanil/key-value/handler"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/key/{key}", handler.PutHandler).Methods("PUT")
	r.HandleFunc("/v1/key/{key}", handler.GetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3030", r))
}

func helloMuxHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello mux")
}
