package controllers

import (
	"net/http"

	"github.com/gorilla/schema" //This package makes it easier to convert form values into a Go struct.
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err //function wont ever panic and will instead return an error when one occurs
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
