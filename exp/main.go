//experimental file. Not to be used in production
package main

import (
	"database/sql" //allows us access our SQL database
	"fmt"

	_ "github.com/lib/pq" //implements the Postgres driver used by the database/sql package.
	//we need this package to  be imported so that it’s init() function gets called which will register the
	//"postgres" driver for the database/sql package.
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Atib0b00"
	dbname   = "postgres"
)

func main() {
	//Creating the connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Opening a database connection. Returns an *sql.DB if successful
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//Pinging the database. Verifies that our code actually tries to talk to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	db.Close()
}
