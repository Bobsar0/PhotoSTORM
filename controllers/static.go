package controllers

import "github.com/bobsar0/PhotoSTORM/views"

//Static controller. Holds all of our pages that are pretty close to being static pages.
type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}

//NewStatic function initializes our static controller and our views.
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq:     views.NewView("bootstrap", "static/faq"),
	}
}
