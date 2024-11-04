package handler

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"mercanil/key-value/model"
	"net/http"
)

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	model.Store.RLock()
	defer model.Store.RUnlock()
	value, err := model.Get(key)
	if err != nil {
		if errors.Is(model.ErrNoSuchKey, err) {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	_, err = fmt.Fprintf(writer, value)
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func PutHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	value, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	defer req.Body.Close()

	err = model.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
