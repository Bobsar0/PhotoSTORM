package main

import (
	"net/http"

	"github.com/bobsar0/PhotoSTORM/controllers"

	"github.com/gorilla/mux"
)

//Handler functions deleted. Home, contact and faq pages now served via the ServeHTTP handler method on the views.View type inside of our static controller.

func main() {
	//Views initializations also deleted as not needed
	
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	staticG := controllers.NewGalleries()

	r := mux.NewRouter() //New gorilla/mux router
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.Handle("/gallery", staticG.Gallery).Methods("GET")

	r.HandleFunc("/signup", usersC.NewUserForm).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}
