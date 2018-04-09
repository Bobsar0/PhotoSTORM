package main

import (
	"net/http"

	"github.com/bobsar0/PhotoSTORM/views"

	"github.com/gorilla/mux"
)

var homeView *views.View    //stores view for home page consisting of home.gohtml template and layouts
var contactView *views.View //stores view for contact page consisting of contact.gohtml template and layouts
var faqView *views.View     //stores view for faq page consisting of faq.gohtml template and layouts

//'home' handler function processes incoming web requests to access home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}

//'contact' handler function processes incoming web requests to access contact page
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}

//'faq' handler function processes incoming web requests to access faq page
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}
func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	r := mux.NewRouter()              //New gorilla/mux router
	r.HandleFunc("/", home)           // tells the router to call the home function when the user wants to visit home page - indicated by path '/'
	r.HandleFunc("/contact", contact) // tells the router to call the home function when the user wants to visit contact page - indicated by path '/contact'
	r.HandleFunc("/faq", faq)         // tells the router to call the home function when the user wants to visit faq page - indicated by path '/faq'

	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}