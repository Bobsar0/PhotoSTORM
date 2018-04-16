//Experimental file. Not to be used in production
//Using GORM to interact with a database.

package main

import (
	"fmt"

	"github.com/bobsar0/PhotoSTORM/models"

	_ "github.com/lib/pq" //implements the Postgres driver.
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your_password"
	dbname   = "photostorm_dev"
)

func main() {
	//Creating the connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Opening a database connection. Returns a *gorm.DB if successful
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	//Create a user
	user := models.User{
		Name: "Steve Onyeneke",
		Email: "bob@g",
	}
	if err := us.Create(&user); err!=nil{
		panic(err)
	}

	// Update a user
	user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	// Update the call to ByEmail
	foundUser, err := us.ByEmail("bob@g")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	if err := us.Delete(foundUser.ID); err != nil {
		panic(err)
	}
		// Verify the user is deleted
	_, err = us.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}
}//end main

