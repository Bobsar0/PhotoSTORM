package middleware

import (
	"fmt"
	"net/http"

	"github.com/bobsar0/PhotoSTORM/context"
	"github.com/bobsar0/PhotoSTORM/models"
)

type RequireUser struct {
	models.UserService
}

// ApplyFn will return an http.HandlerFunc that will check to see if a user is logged in
// and then either call next(w, r) if they are, or redirect them to the
// login page if they are not.
func (ru *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	// We want to return a dynamically created func(http.ResponseWriter, *http.Request)
	// but we also need to convert it into an http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if a user is logged in.
		// If so, call next(w, r)
		// If not, http.Redirect to "/login"
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		//use the UserService to look up a user via the remember
		// token stored on the cookie we just retrieved. As long as the user has a valid
		// session this should return a user, otherwise we will get an error indicating that
		// the user needs redirected to the login page.
		user, err := ru.UserService.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		fmt.Println("User found: ", user)
		// Get the context from our request
		ctx := r.Context()
		// Create a new context from the existing one that has
		// our user stored in it with the private user key
		ctx = context.WithUser(ctx, user)
		// Create a new request from the existing one with our
		// context attached to it and assign it back to `r`.
		r = r.WithContext(ctx)
		// Call next(w, r) with our updated context.
		next(w, r)
	})
}

func (mw *RequireUser) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}
