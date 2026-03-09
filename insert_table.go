package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func insert_table(input Input, result Output) {
	connStr := "user=postgres password=dbpass dbname=calculator sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO calculator (first_number, sign, second_number, result) VALUES ($1, $2, $3, $4)", input.FirstNumber, input.Sign, input.SecondNumber, result.Result)
	if err != nil {
		panic(err)
	}
}
