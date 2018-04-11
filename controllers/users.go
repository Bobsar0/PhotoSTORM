package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema" //This package makes it easier to convert form values into a Go struct.

	"github.com/bobsar0/PhotoSTORM/views"
)

//Users controller
type Users struct {
	NewView *views.View //stores the new user view
}

//NewUsers func sets up all the views our users controller will need.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// NewUserForm is used to render the form where a user can create a new user account.
// GET /signup
func (u *Users) NewUserForm(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//SignupForm contains fields required tobe filled by the user in the form
type SignupForm struct {
	Email    string `schema:"email"` //the struct tags are to let the schema package know about the input fields
	Password string `schema:"password"`
}

//Create is used to process the signup form when a user tries to create a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	form := SignupForm{}
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
