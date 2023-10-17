package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

const connectionString = "host=localhost port=5432 user=postgres password=Kailali@123 dbname=krane sslmode=disable"

var Db *goqu.Database
var Ds *goqu.Dataset

func DatabaseConnection() {

	var err error
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	Db = goqu.New("postgres", db)

	Ds = goqu.From("event")

	fmt.Println("database is ready for use")

	// You can now perform database operations using the 'db' object.
}
