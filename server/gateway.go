package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CurrencyConversion struct {
	USDBRL struct {
		BID string `json:"bid"`
	}
}

type CurrencyGateway interface {
	FetchCurrencyConversion(fromCurrency string, toCurrency string) (CurrencyConversion, error)
}

type economiaGateway struct{}

func NewCurrencyGateway() CurrencyGateway {
	return economiaGateway{}
}

func (g economiaGateway) FetchCurrencyConversion(fromCurrency string, toCurrency string) (CurrencyConversion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("create request with context failed: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("fetching currency conversion failed: %w", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("reading response body failed: %w", err)
	}

	var conversion CurrencyConversion
	err = json.Unmarshal(resBody, &conversion)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("parsing response body to struct failed %w", err)
	}

	return conversion, nil
}
