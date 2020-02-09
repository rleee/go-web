package controllers

import (
	"004-go-web/views"
	"fmt"
	"net/http"
)

// Users struct
type Users struct {
	NewView *views.View
}

// NewUsers will parse the signup form template
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New will be used to render signup form
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create will creating new user
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User created")
}
