package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"bytes"
	"io"
)

//Adding glob helper variables
var (
	LayoutDir   = "views/layouts/" //specifies the layout directory
	TemplateDir = "views/"
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
	//Using the view file path helpers
	addTemplatePath(files) //prepends "views/" to filename
	addTemplateExt(files) //appends ".gohtml" to filename

	files = append(files, layoutFiles()...) //We pass in all of	the files returned by the layoutFiles function call (... unpacks the files and lists them as comma separated values)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalln(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//Render contains logic for rendering the view. Offloads the responsibility from our handlers.
func (v *View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	
	switch data.(type) {
	case Data:
		// do nothing
	default:
		data = Data{
			Yield: data,
		}
	}
	var buf bytes.Buffer
	err := v.Template.ExecuteTemplate(&buf, v.Layout, data)
	if err != nil {
		http.Error(w, "Something went wrong. If the problem persists, please email support@lenslocked.com",
			http.StatusInternalServerError)
		return
	}
	// If we get here that means our template executed correctly
	// and we can copy the buffer to the ResponseWriter
		io.Copy(w, &buf)
}

//ServeHTTP() method is written so our View type can implement http.Handler interface for use in the r.Handle() method
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

// addTemplatePath takes in a slice of strings representing file paths for templates, 
//and it prepends the TemplateDir directory to each string in the slice.
// Eg the input {"home"} would result in the output: {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, j := range files {
		files[i] = TemplateDir + j
	}
}
// addTemplateExt takes in a slice of strings representing file paths for templates 
//and it appends the TemplateExt extension to each string in the slice
//Eg the input {"home"} would result in the output {"home.gohtml"} if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, j := range files {
		files[i] = j + TemplateExt
	}
}