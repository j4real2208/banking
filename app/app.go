package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/logger"
	"github.com/j4real2208/banking/service"
	"github.com/jmoiron/sqlx"
)
func sanityCheck()  {
	if ( os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWD") == ""|| os.Getenv("DB_NAME") == "" ) {
		logger.Info("Enviornment variables not set or defined .....")		
	}
}
func Start() {
	sanityCheck()
	//Our Server 
	router := mux.NewRouter()

	// Wiring and injecting dependecies
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	//accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	//ch:= CustomerHandlers
	
	// Route Handling
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomers).Methods(http.MethodGet)

	// Retriving Address and Port
	addr := os.Getenv("SERVER_ADDRESS")
	prt := os.Getenv("SERVER_PORT")
	//Starting New Server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",addr,prt), router))
}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbName := os.Getenv("DB_NAME")
	addr := os.Getenv("SERVER_ADDRESS")
	dbPrt := os.Getenv("DB_PORT")
	
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPasswd,addr,dbPrt,dbName)	
	
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}