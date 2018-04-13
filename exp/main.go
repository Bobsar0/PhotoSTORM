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
	dbname   = "photostorm_dev"
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
	defer db.Close()

	// //Inserting our first user
	// _, err = db.Exec(`
	// 	INSERT INTO users(name,email)
	// 	VALUES($1, $2)`, //We are NOT trying to build this string ourselves, but instead we are using placeholders like $1 and we are letting the database/sql package handle creating the SQL statement.
	// 	"Chioma Onyeneke", "ccchioma60@gmail.com")
	// if err != nil {
	// 	log.Fatalln(err)
	//************************************************//

	// //Inserting and retrieving the ID of the new record
	// var id int

	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id`,
	// 	"Jon Calhoun", "jon@calhoun.io")
	// err = row.Scan(&id) // Scan copies the columns from the matched row into the values pointed at by id
	// if err != nil {
	// 	panic(err)
	// }
	//**************************************************//
	// //Querying a single record with database/sql
	// var name, email string
	// row1 := db.QueryRow(`
	// SELECT id, name, email
	// FROM users
	// WHERE id=$1`, 10)
	// err = row1.Scan(&id, &name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ID:", id, "Name:", name, "Email:", email)

	// //Querying multiple records with database/sql
	// //var name, email string
	// rows, err := db.Query(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE email = $1
	// 	OR ID > $2`,
	// 	"jon@calhoun.io", 3)
	// if err != nil {
	// 	panic(err)
	// }
	// for rows.Next() { //tells our rows object that we are ready to move on to the next row.
	// 	rows.Scan(&id, &name, &email) // retrieves the data from query and tell the method which parameters we want to store that data in.
	// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	// }

	var id int
	for i := 1; i < 6; i++ {
		//Create some fake data
		userID := 1
		if i > 3 {
			userID = 2
		}
		amount := 1000 * i
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		err = db.QueryRow(`
			INSERT INTO orders (user_id, amount, description)
			VALUES ($1, $2, $3)
			RETURNING id`,
			userID, amount, description).Scan(&id) //retrieve the ID from the returned sql.Row object
		if err != nil {
			panic(err)
		}
		fmt.Println("Created an order with the ID:", id)
	}

} //end main
