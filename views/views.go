package views

import (
	"github.com/gorilla/mux"
)

/**
Write Views in this package:
func Example(w http.ResponseWriter, r *http.Request){}
*/
import (
	"fmt"
	"net/http"
	// "../models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		http.NotFound(w, r)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

func Number(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	params := mux.Vars(r)
	fmt.Fprintf(w, params["number"])
}
