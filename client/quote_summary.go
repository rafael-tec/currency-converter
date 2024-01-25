package main

import (
	"fmt"
	"os"
	"time"
)

const (
	_fileNameSufix     = "currency_conversions"
	_currencyFieldName = "dollar"
	_quoteTime         = "time"
	filePermission     = 6
)

type CurrencyQuoteSummary interface {
	AppendQuote(dateSummary time.Time, quoteTime time.Time, fromCurrency string, bid string) error
}

type currencyQuoteSummary struct{}

func NewCurrencyQuoteSummary() CurrencyQuoteSummary {
	return currencyQuoteSummary{}
}

func (s currencyQuoteSummary) AppendQuote(
	dateTimeSummary time.Time,
	quoteTime time.Time,
	fromCurrency string,
	bid string,
) error {
	dateSummary := dateTimeSummary.Format(time.DateOnly)
	fileName := fmt.Sprintf("%s-%s", dateSummary, _fileNameSufix)

	var file *os.File
	var err error

	if notExists := s.fileExistsFor(dateSummary); notExists {
		file, err = os.Create(fileName)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, filePermission)
		if err != nil {
			return err
		}
	}

	defer file.Close()

	lineContent := fmt.Sprintf(
		"%s:%s %s:%s",
		_currencyFieldName,
		bid,
		_quoteTime,
		quoteTime.Format(time.DateTime),
	)

	_, err = file.WriteString(lineContent + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (s currencyQuoteSummary) fileExistsFor(dateSummary string) bool {
	fileName := fmt.Sprintf("%s-%s", dateSummary, _fileNameSufix)

	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}
