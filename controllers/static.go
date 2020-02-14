package controllers

import (
	"004-go-web/views"
)

// NewStatic to make new static page of Home or Contact
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static will hold Home and Contact
type Static struct {
	Home    *views.View
	Contact *views.View
}
