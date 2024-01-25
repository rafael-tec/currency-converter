package main

import "net/http"

func main() {
	InitLoggers()

	InfoLoggger.Println("Booting 'client' application")

	quoteSummary := NewCurrencyQuoteSummary()
	gateway := NewCurrencyGateway()
	conversionService := NewCurrencyConversionService(gateway, quoteSummary)
	handler := NewCurrencyHandler(conversionService)

	InfoLoggger.Println("Listening in 8081 port")

	http.HandleFunc("/client/quote", handler.GetQuoteHandler)
	http.ListenAndServe(":8081", nil)
}
