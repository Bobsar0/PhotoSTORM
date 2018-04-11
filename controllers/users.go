package controllers

import (
	"fmt"
	"net/http"

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

//Create is used to process the signup form when a user tries to create a new user account.

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, r.PostForm["email"])    //prints value stored in our PostForm map with key "email"
	fmt.Fprintln(w, r.PostForm["password"]) //prints value stored in our PostForm map with key "password"
}
