package base

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	databaseHost := "localhost"
	databaseUser := "student"

	connection :=
		fmt.Sprintf("postgresql://%s:qwerty@%s:5432/proxy_db?sslmode=disable", databaseUser, databaseHost)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil

}
