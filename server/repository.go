package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CurrencyRepository interface {
	Save(bid string, quoteTime time.Time) error
}

type currencyRepository struct{}

func NewCurrencyRepository() CurrencyRepository {
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sts := `
		DROP TABLE IF EXISTS currency_conversions;
	
		CREATE TABLE currency_conversions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bid DECIMAL(10,2) NOT NULL,
			quote_time VARCHAR(30) NOT NULL
		);

		INSERT INTO currency_conversions(
			bid,
			quote_time
		) VALUES(
			7.77,
			"2024-01-30 22:22:22"
		);

		INSERT INTO currency_conversions(
			bid,
			quote_time
		) VALUES(
			3.33,
			"2024-01-29 21:10:03"
		);
	`
	_, err = db.Exec(sts)

	if err != nil {
		ErrorLogger.Fatal(err)
	}

	return currencyRepository{}
}

func (r currencyRepository) Save(bid string, quoteTime time.Time) error {
	db, err := sql.Open("sqlite3", ":memory")
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

	dateTime := quoteTime.Format(time.DateTime)

	_, err = stmt.ExecContext(ctx, bid, dateTime)
	if err != nil {
		return err
	}

	return nil
}
