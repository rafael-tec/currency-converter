package main

import "time"

const (
	_defaultFromCurrency = "USD"
	_defaultToCurrency   = "BRL"
)

type CurrencyQuoteService interface {
	RegisterCurrencyQuote() error
}

type currencyQuoteService struct {
	gateway      CurrencyGateway
	quoteSummary CurrencyQuoteSummary
}

func NewCurrencyConversionService(
	gateway CurrencyGateway,
	quoteSummary CurrencyQuoteSummary,
) CurrencyQuoteService {
	return currencyQuoteService{gateway: gateway, quoteSummary: quoteSummary}
}

func (s currencyQuoteService) RegisterCurrencyQuote() error {
	bid, err := s.gateway.FetchCurrencyQuote(_defaultFromCurrency, _defaultToCurrency)
	if err != nil {
		ErrorLogger.Printf("Fetch currency quote occurred error: %s", err.Error())
		return err
	}

	err = s.quoteSummary.AppendQuote(time.Now(), time.Now(), _defaultFromCurrency, bid.BID)
	if err != nil {
		ErrorLogger.Printf("Append quote in summary occurred error: %s", err.Error())
		return err
	}

	return nil
}
