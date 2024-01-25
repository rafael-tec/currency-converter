package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitLoggers()

	InfoLoggger.Println("Booting 'server' application")

	repository := NewCurrencyRepository()
	gateway := NewCurrencyGateway()
	service := NewCurrencyConversionService(gateway, repository)
	handler := NewCurrencyConversionHandler(service)

	InfoLoggger.Println("Listening in 8080 port")

	http.HandleFunc("/server/cotacao", handler.GetCurrencyConversionHandler)
	http.ListenAndServe(":8080", nil)
}
