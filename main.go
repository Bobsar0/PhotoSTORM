package main

import (
	"fmt"
	"net/http"

	"github.com/bobsar0/PhotoSTORM/controllers"
	"github.com/bobsar0/PhotoSTORM/models"

	"github.com/gorilla/mux"
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
	//us.DestructiveReset()
	us.AutoMigrate() //Ensures that our database is migrated properly.

	// Initialize controllers
	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()
	staticG := controllers.NewGalleries()

	r := mux.NewRouter() //New gorilla/mux router
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.Handle("/gallery", staticG.Gallery).Methods("GET")

	r.HandleFunc("/signup", usersC.NewUserForm).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	// NOTE: We are using the Handle function, not HandleFunc
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}
