//Experimental file. Not to be used in production
//Using GORM to interact with postgres database.

package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //implements the Postgres driver.
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your_password"
	dbname   = "photostorm_dev"
)

//User struct reps GORM model
type User struct {
	gorm.Model //has the basic fields that we will almost always want, like the unique ID for each resource and a timestamp for when the resources was created and updated.
	Name       string
	Email      string `gorm:"not null;unique_index"`
	Orders []Order
}
//Order struct 
type Order struct {
	gorm.Model
	UserID uint
	Amount int
	Description string
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID: user.ID,
		Amount: amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}

func main() {
	//Creating the connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Opening a database connection. Returns a *gorm.DB if successful
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	defer db.Close()

	db.LogMode(true) //enables logging so we can track what SQL are run behind the scenes

	db.AutoMigrate(&User{}, &Order{}) //creates corresponding 'users' and 'orders' table with its fields and gorm.Model fields.

	// //Creating a new record with GORM
	// name, email := getInfo() //get the name and email from the user
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil { // The db.Create() method returns a point to a gorm.DB object, so we can then grab the Error attribute to verify that the create happened successfully,
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)
//*******************************************************//
	//Querying for the first user
	var usr User
	db.First(&usr) //retrieves the first record and stores in usr. By passing in an address to a user object, GORM knows to query the users table, and knows to put the resulting record into the object passed in as the first argument.
	if db.Error != nil {
		panic(db.Error)
	}
	// // fmt.Println(usr)
	
	createOrder(db, usr, 1001, "Fake Description #1")
	createOrder(db, usr, 9999, "Fake Description #2")
	createOrder(db, usr, 8800, "Fake Description #3")

	db.Preload("Orders").First(&usr)
 	if db.Error != nil {
 		panic(db.Error)
 	}
	fmt.Println("Email:", usr.Email)
	fmt.Println("Number of orders:", len(usr.Orders))
	fmt.Println("Orders:", usr.Orders)
	
	//************************************************//
	//Querying with the Where() method
// 	var u User
// 	maxID := 3
// 	db.Where("id <= ?", maxID).First(&u)
// 	if db.Error != nil {
// 		panic(db.Error)
// 	}
// 	fmt.Println(u)
//****************************************************//
// //Querying for multiple users
// 	var users []User
// 	db.Find(&users)
// 	if db.Error != nil {
// 		panic(db.Error)
// 	}
// 	fmt.Println("Retrieved", len(users), "users.")
// 	fmt.Println(users)

} //end main

// func getInfo() (name, email string) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("What is your name?")
// 	name, _ = reader.ReadString('\n') //ReadString reads until the first occurrence of delim (\n) in the input, returning a string containing the data up to and including the delimiter.
// 	name = strings.TrimSpace(name)    // returns a slice of the string s, with all leading and trailing white space removed
// 	fmt.Println("What is your email?")
// 	email, _ = reader.ReadString('\n')
// 	email = strings.TrimSpace(email)
// 	return name, email
// }

//===========================================//
//*****USING RAW SQL (database/sql)