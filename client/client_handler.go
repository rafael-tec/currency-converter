package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CurrencyHandler struct {
	service CurrencyQuoteService
}

func NewCurrencyHandler(service CurrencyQuoteService) CurrencyHandler {
	return CurrencyHandler{service: service}
}

func (h CurrencyHandler) GetQuoteHandler(w http.ResponseWriter, r *http.Request) {
	InfoLoggger.Println("Request received")

	err := h.service.RegisterCurrencyQuote()
	if err != nil {
		ErrorLogger.Printf("Request failed: %s", err.Error())

		errorResponse, _ := json.Marshal(ErrorResponse{Message: err.Error()})

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	InfoLoggger.Printf("Request sucessfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w)
}
