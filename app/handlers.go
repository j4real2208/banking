package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)
type Customer struct {
	Name string`json:"full_name" xml:"name"`
	City string	`json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip-code"`
}



func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers:= []Customer {
		{"JOJO","New Delhi","2110002"},
		{"Matt","New Delhi","2110002"},
		{"Rob","New Delhi","2110002"},
	}
	if r.Header.Get("Content-Type")=="application/xml"{
		w.Header().Add("Content-Type","application/xml")		
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(customers)
	}
}