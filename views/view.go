package views

import "html/template"

// View will return pointer to parsed template
type View struct {
	Template *template.Template
}

// NewView will get parsed template
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
