package views

import "html/template"

// NewView appends template files to the list of files provided, parses the template files, constructs the *View for us and returns it.
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template //stores the parsed template that we want to execute
}
