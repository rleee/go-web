package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   string = "views/layouts/"
	templateExt string = ".gohtml"
)

// View will return pointer to parsed template
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render template attached to View type
func (v View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// NewView will get parsed template
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) // using ... is to unpack the slice

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// returns all files location string based on
// variable arguments provided in Glob
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}
