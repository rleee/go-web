package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFound)
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/profile/:userId", profile)
	http.ListenAndServe(":8000", router)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello there!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Checkout our latest collections at <a href=\"testlink.com\">fotoku.com</a>")
}

func profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hello there! %s, welcome back.", ps.ByName("userId"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h3>Sorry, the page you are looking for doesn't exists :(</h3> <p>If you encounter this multiple times, please contact us.</p>")
}
