package controllers

import (
	"004-go-web/views"
)

// NewStatic to make new static page of Home or Contact
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

// Static will hold Home and Contact
type Static struct {
	Home    *views.View
	Contact *views.View
}
