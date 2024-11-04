package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	ErrNoSuchKey = errors.New("No such a key")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxHandler)

	log.Fatal(http.ListenAndServe(":3030", r))
}

func helloMuxHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello mux")
}

func Put(key, val string, store map[string]string) error {
	store[key] = val
	return nil
}

func Get(key string, store map[string]string) (string, error) {
	val, ok := store[key]
	if !ok {
		return "", ErrNoSuchKey
	}
	return val, nil

}

func Delete(key string, store map[string]string) error {
	delete(store, key)
	return nil
}
