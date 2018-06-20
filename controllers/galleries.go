package controllers

import (
	"fmt"
	"net/http"

	"github.com/bobsar0/PhotoSTORM/context"
	"github.com/bobsar0/PhotoSTORM/models"
	"github.com/bobsar0/PhotoSTORM/views"
)

type Galleries struct {
	New *views.View
	ShowView *views.View
	gs  models.GalleryService
}

func NewGalleries(gs models.GalleryService) *Galleries {
	//Sets up the views
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		ShowView: views.NewView("bootstrap", "galleries/show"),
		gs:  gs,
	}
}

type GalleryForm struct {
	Title string `schema:"title"`
}

// POST /galleries
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form GalleryForm
	// Parse the data submitted in the web request into a GalleryForm instance.
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	// Retrieve user from context
	user := context.User(r.Context())
	// Then we update how we build the Gallery model
	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}
	// Call the GalleryService’s Create method passing in the gallery and checking for errors.
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)
}
