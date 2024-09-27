package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret1234"
	dbname   = "go_db"
)

var DB *sql.DB

func InitDB() {

	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Could not connect to database...")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nConnected......")

	createKeywordTable()
	createUserTable()
}

func createKeywordTable() {
	query := `CREATE TABLE IF NOT EXISTS keywords(
		id SERIAL PRIMARY KEY,
		keyword TEXT NOT NULL,
		user_id INT
	)`

	_, err := DB.Exec(query)

	if err != nil {
		panic("Could not create keywords table")
	}

}

func createUserTable() {
	query := `CREATE TABLE IF NOT EXISTS user_tbl(
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(query)

	if err != nil {
		panic("Could not create user tble")
	}
}
