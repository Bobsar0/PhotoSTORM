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
	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close() // Defer closing database until our main function exits,
	services.AutoMigrate() //Ensures that our database is migrated properly.

	// Initialize controllers
	usersC := controllers.NewUsers(services.User)
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries(services.Gallery)

	r := mux.NewRouter() //New gorilla/mux router
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	r.HandleFunc("/signup", usersC.NewUserForm).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	r.Handle("/galleries/new", galleriesC.New).Methods("GET")
	r.HandleFunc("/galleries", galleriesC.Create).Methods("POST")

	fmt.Println("Starting the server on :8000...")
	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}
