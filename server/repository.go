package main

import (
	"context"
	"database/sql"
	"time"
)

type CurrencyRepository interface {
	Save(bid string, quoteTime time.Time) error
}

type currencyRepository struct{}

func NewCurrencyRepository() CurrencyRepository {
	return currencyRepository{}
}

func (r currencyRepository) Save(bid string, quoteTime time.Time) error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/currency")
	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO currency_conversions (bid, quote_time) VALUES(?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	_, err = stmt.ExecContext(ctx, bid, quoteTime)
	if err != nil {
		return err
	}

	return nil
}
