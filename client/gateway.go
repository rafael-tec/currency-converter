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
	BID string `json:"bid"`
}

type CurrencyGateway interface {
	FetchCurrencyQuote(fromCurrency string, toCurrency string) (CurrencyConversion, error)
}

type currencyGateway struct{}

func NewCurrencyGateway() CurrencyGateway {
	return currencyGateway{}
}

func (g currencyGateway) FetchCurrencyQuote(fromCurrency string, toCurrency string) (CurrencyConversion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/server/cotacao", nil)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("create request with context failed: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("fetch currency conversion failed: %w", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("read response body failed: %w", err)
	}

	if res.StatusCode != 200 {
		return CurrencyConversion{}, fmt.Errorf("status code: %d message: %s", res.StatusCode, resBody)
	}

	var conversion CurrencyConversion
	err = json.Unmarshal(resBody, &conversion)
	if err != nil {
		return CurrencyConversion{}, fmt.Errorf("parse response body to struct failed %w", err)
	}

	return conversion, nil
}
