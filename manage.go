package main

/** Never Mind this Package */
/** Write Your views in views package */
/** Write Your urlpatterns in urls package */
/** Make Models Methods in models package */

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"os"

	"./urls"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "runserver":
			runserver()
		default:
			fmt.Println("Command Not Found:", os.Args[1])
		}
	} else {
		fmt.Println("Expecting Command")
	}
}

func runserver() {
	router := mux.NewRouter()
	for _, url := range urls.UrlPatterns {
		router.HandleFunc(url.Url, url.View)
	}
	fmt.Println("Starting Server http://localhost:8080/")
	http.ListenAndServe(":8080", router)
}
