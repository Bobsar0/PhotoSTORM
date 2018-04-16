package controllers

import (
	"fmt"
	"net/http"

	"github.com/bobsar0/PhotoSTORM/models"
	"github.com/bobsar0/PhotoSTORM/views"
)

//Users controller
type Users struct {
	NewView *views.View //stores the new user view
	LoginView *views.View //stores the login view
	us *models.UserService //for easy access by handler methods
}

//NewUsers func sets up all the views our users controller will need.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:	us,
	}
}

// NewUserForm is used to render the form where a user can create a new user account. GET /signup
func (u *Users) NewUserForm(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//SignupForm contains fields required to be filled by the user in the form
type SignupForm struct { //the struct tags are to let the schema package know about the input fields
	Name	 string `schema:"name"`
	Email    string `schema:"email"` 
	Password string `schema:"password"`
}

type LoginForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
} 

//Create is used to process the signup form when a user tries to create a new user account. POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User is", user)
}

// Login is used to process the login form when a user tries to log in as an existing user (via email & pw).
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	// We will eventually do something to see if the
	// information provided is correct.
}