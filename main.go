package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Your current path on out site: ", r.URL.Path, "<br>")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello there!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Checkout our latest collections at <a href=\"testlink.com\">fotoku.com</a>")
	}

}
