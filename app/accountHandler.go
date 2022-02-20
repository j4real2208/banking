package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ch *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request)  {
	vars:= mux.Vars(r)
	customerId := vars["customer_id"]
	var req dto.NewAccountRequest
	err:= json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w,http.StatusBadRequest,err.Error())		
	}else{
		req.CustomerId = customerId
		
		// Log Statment
		//logger.Info("--------Inside account handler to create new account and goint to handler func----- CustomerId->"+customerId)

		account , appError := ch.service.NewAccount(req)
		
		if appError != nil {
			//logger.Error("Came back to handler to print account erorr")
			writeResponse(w, appError.Code, appError.Message)
		}else{
			writeResponse(w,http.StatusCreated,account)
		}

	}
}