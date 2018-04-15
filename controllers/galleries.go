//Not to be used yet in production 
package controllers

import(
	"net/http"

	"github.com/bobsar0/PhotoSTORM/views"
)

//Galleries controller
type Galleries struct{
	Gallery *views.View
}

func NewGalleries () *Galleries{
	return &Galleries{
		Gallery: views.NewView("bootstrap", "gallery"),
	}
}

func (g *Galleries) NewUserForm(w http.ResponseWriter, r *http.Request) {
	if err := g.Gallery.Render(w, nil); err!=nil{
		panic(err)
	}
}

