package main

import (
	"time"
)

const (
	_defaultFromCurrency = "USD"
	_defaultToCurrency   = "BRL"
)

type CurrencyConversionService interface {
	FetchCurrencyConversion() (string, error)
}

type currencyConversionService struct {
	gateway    CurrencyGateway
	repository CurrencyRepository
}

func NewCurrencyConversionService(
	gateway CurrencyGateway,
	repository CurrencyRepository,
) CurrencyConversionService {
	return currencyConversionService{gateway: gateway, repository: repository}
}

func (s currencyConversionService) FetchCurrencyConversion() (string, error) {
	conversion, err := s.gateway.FetchCurrencyConversion(_defaultFromCurrency, _defaultToCurrency)
	if err != nil {
		ErrorLogger.Printf("Fetching currency conversion occurred error: %s", err.Error())
		return "", err
	}

	err = s.repository.Save(conversion.USDBRL.BID, time.Now())
	if err != nil {
		ErrorLogger.Printf("Persist currency quote occurred error: %s", err.Error())
		return "", err
	}

	return conversion.USDBRL.BID, nil
}
