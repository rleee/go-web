package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   string = "views/layouts/"
	templateDir string = "views/"
	templateExt string = ".gohtml"
)

// View will return pointer to parsed template
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render template attached to View type
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// ServeHTTP will serving view
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// NewView will get parsed template
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
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

// addTemplatePath, append template path
// infront of template name
//
// Eg: {home} become {views/home}
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templateDir + f
	}
}

// addTemplateExt, append template extention
// to behind of template name
//
// Eg: {home} become {home.gohtml}
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + templateExt
	}
}
