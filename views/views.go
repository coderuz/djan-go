package views
/** 
	Write Views in this package:
	func Example(w http.ResponseWriter, r *http.Request){}
*/
import (
	"net/http"
	"fmt"
	// "../models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		http.NotFound(w, r)
	}
}

func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is about page")
}
