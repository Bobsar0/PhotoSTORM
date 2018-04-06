package main

import (
	"fmt"      //package used to format strings and output them to different places.
	"net/http" //package used to both create an app capable of responding to web requests and making web requests to other servers.

	"github.com/gorilla/mux" //The gorrila/mux router
)

//'home' handler function processes incoming web requests to access home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to PhotoSTORM.com!</h1>")
}

//'contact' handler function processes incoming web requests to access contact page
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@PhotoSTORM.com\">"+
		"support@PhotoSTORM.com</a>.")
}

//'faq' handler function processes incoming web requests to access faq page
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Check back for <b>Frequently Asked Questions.</b>")
}
func main() {
	r := mux.NewRouter()              //New gorilla/mux router
	r.HandleFunc("/", home)           // tells the router to call the home function when the user wants to visit home page - indicated by path '/'
	r.HandleFunc("/contact", contact) // tells the router to call the home function when the user wants to visit contact page - indicated by path '/contact'
	r.HandleFunc("/faq", faq)         // tells the router to call the home function when the user wants to visit faq page - indicated by path '/faq'

	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}
