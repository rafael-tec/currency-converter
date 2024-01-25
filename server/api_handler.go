package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CurrencyQuoteResponse struct {
	BID string `json:"bid"`
}

type CurrencyConversionHandler struct {
	service CurrencyConversionService
}

func NewCurrencyConversionHandler(service CurrencyConversionService) CurrencyConversionHandler {
	return CurrencyConversionHandler{service: service}
}

func (h CurrencyConversionHandler) GetCurrencyConversionHandler(w http.ResponseWriter, r *http.Request) {
	InfoLoggger.Println("Request received")

	bid, err := h.service.FetchCurrencyConversion()
	if err != nil {
		ErrorLogger.Printf("Request failed: %s", err)

		errorResponse, _ := json.Marshal(ErrorResponse{Message: err.Error()})

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resBody, err := json.Marshal(CurrencyQuoteResponse{BID: bid})
	if err != nil {
		ErrorLogger.Printf("Request failed: %s", err)

		errorResponse, _ := json.Marshal(ErrorResponse{Message: err.Error()})

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	InfoLoggger.Println("Request sucessfully")

	w.Write([]byte(resBody))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w)
}
