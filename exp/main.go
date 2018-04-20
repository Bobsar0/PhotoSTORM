package main

import (
	"fmt"

	"github.com/bobsar0/photostorm/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your_password"
	dbname   = "photostorm_dev"
)

func main() {
	// Create a DB connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Create our model services
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close() // Defer closing database until our main function exits,
	us.DestructiveReset()
	//us.AutoMigrate() //Ensures that our database is migrated properly.

	user := models.User{
		Name:     "Steve Onyeneke",
		Email:    "bob@gmail.com",
		Password: "bob",
	}
	err = us.Create(&user)
	if err != nil {
		panic(err)
	}

	// Verify that the user has a Remember and RememberHash
	fmt.Printf("%+v\n", user)
	if user.Remember == "" {
		panic("Invalid remember token")
	}
	// Verify that we can lookup a user with that remember token
	user2, err := us.ByRemember(user.Remember)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)
}
