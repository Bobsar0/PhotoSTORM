package main

import (
	"fmt" 	   //package used to format strings and output them to different places.
	"net/http" //package used to both create an app capable of responding to web requests and making web requests to other servers.
)

//handler function processes incoming web requests and determines what to return to the client
func handlerFunc(w http.ResponseWriter, r *http.Request) { // 'w' allows us to modify response to client. 'r' is used to access data from the web request
	fmt.Fprint(w, "<h1>Welcome to PhotoSTORM.com</h1>") //'fmt.Fprint' writes the HTML string to 'w'
}
func main() {
	http.HandleFunc("/", handlerFunc) // handlerFunc is set as the function to handle web requests going to our server with the path '/'.
	http.ListenAndServe(":8080", nil) // starts up a web server listening on port 8080 using the default http handlers identified by 'nil'.
}
