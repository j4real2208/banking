package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4real2208/banking/service"
)


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
		fmt.Fprintf(w,err.Error())
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
		writeResponse(w,err.Code,err.AsMessage())
   }else{
	writeResponse(w,http.StatusOK,customer)
   }

}

func writeResponse(w http.ResponseWriter , code int , data interface{})  {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	if err:=json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	} 
}