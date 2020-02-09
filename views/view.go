package views

import "html/template"

// View will return pointer to parsed template
type View struct {
	Template *template.Template
	Layout   string
}

// NewView will get parsed template
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml",
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
