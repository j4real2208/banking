package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//Our Server 
	mux := mux.NewRouter()
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World !! ")
	})
	mux.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}