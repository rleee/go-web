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
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// New will be used to render signup form
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// SignUpForm will be use to store form data
type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create will creating new user
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignUpForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
