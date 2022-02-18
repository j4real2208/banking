package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4real2208/banking/service"
)
type Customer struct {
	Name string`json:"full_name" xml:"name"`
	City string	`json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip-code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers:= []Customer {
	// 	{"JOJO","New Delhi","2110002"},
	// 	{"Matt","New Delhi","2110002"},
	// 	{"Rob","New Delhi","2110002"},
	// }
	// Calling the customers
	customers ,err := ch.service.GetAllCustomer()
	if err != nil {
		 log.Default()
	}
	if r.Header.Get("Content-Type")=="application/xml"{
		w.Header().Add("Content-Type","application/xml")		
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers)getCustomers(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	id:= vars["customer_id"]
	customer , err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Default()
   }else{
		w.Header().Add("Content-Type","application/json")
	 	json.NewEncoder(w).Encode(customer)  
   }

}