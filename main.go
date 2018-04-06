package main

import (
	"fmt"      //package used to format strings and output them to different places.
	"net/http" //package used to both create an app capable of responding to web requests and making web requests to other servers.
)

//handler function processes incoming web requests and determines what to return to the client
func handlerFunc(w http.ResponseWriter, r *http.Request) { // 'w' allows us to modify response to client. 'r' is used to access data from the web request
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" { // if home page
		fmt.Fprint(w, "<h1>Welcome to PhotoSTORM.com!</h1>")
	} else if r.URL.Path == "/contact" { //if contact page
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@PhotoSTORM.com\">"+
			"support@PhotoSTORM.com</a>.")
	} else { 
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1>"+
			"<p>Please email us if you keep being sent to an "+
			"invalid page.</p>")
	}
}
func main() {
	http.HandleFunc("/", handlerFunc) // handlerFunc is set as the function to handle web requests going to our server with the path '/'.
	http.ListenAndServe(":8080", nil) // starts up a web server listening on port 8080 using the default http handlers identified by 'nil'.
}
