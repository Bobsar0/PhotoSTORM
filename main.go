package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template    // stores our parsed home.gohtml template to use inside home function
var contactTemplate *template.Template // stores our parsed contact.gohtml template to use inside contact function
var faqTemplate *template.Template     // stores our parsed faq.gohtml template to use inside faq function

//'home' handler function processes incoming web requests to access home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}

//'contact' handler function processes incoming web requests to access contact page
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}

//'faq' handler function processes incoming web requests to access faq page
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTemplate.Execute(w, nil); err != nil { //write the results to w to ensure they are returned to the user who is making a web request.
		panic(err)
	}
}
func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml") //Parse-Files will open up the template file and attempt to validate it.
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	faqTemplate, err = template.ParseFiles("views/faq.gohtml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()              //New gorilla/mux router
	r.HandleFunc("/", home)           // tells the router to call the home function when the user wants to visit home page - indicated by path '/'
	r.HandleFunc("/contact", contact) // tells the router to call the home function when the user wants to visit contact page - indicated by path '/contact'
	r.HandleFunc("/faq", faq)         // tells the router to call the home function when the user wants to visit faq page - indicated by path '/faq'

	http.ListenAndServe(":8080", r) // starts up a web server listening on port 8080 using our gorilla/mux router as the default handler for web requests.

}
