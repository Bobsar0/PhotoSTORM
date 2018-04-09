package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

//Adding glob helper variables
var (
	LayoutDir   = "views/layouts/" //specifies the layout directory
	TemplateExt = ".gohtml"        //tells us what extension we expect all our template files to match
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt) //returns a slice of filepaths as []string that includes all files ending with .gohtml
	if err != nil {
		panic(err)
	}
	return files
}

//View stores our application templates and ultimately renders them as HTML pages
type View struct {
	Template *template.Template
	Layout   string // stores the layout template we want our View to execute
}

// NewView appends template files to the list of files provided, parses the template files, constructs the *View for us and returns it. Only called when program first starts running
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) //We pass in all of	the files returned by the layoutFiles function call (... unpacks the files and lists them as comma separated values)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//Render contains logic for rendering the view. Offloads the responsibility from our handlers.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
