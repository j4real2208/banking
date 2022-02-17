package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/service"
)

func Start() {
	//Our Server 
	router := mux.NewRouter()

	// Wiring and injecting dependecies
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	
	// Route Handling
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//Starting New Server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}