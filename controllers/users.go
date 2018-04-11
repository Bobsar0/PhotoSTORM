package controllers

import (
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