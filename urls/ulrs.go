package urls
/** 
	Write UrlPatterns in this package:
	var UrlPatterns = []Url{
		Url{"/example", views.Example},	
	}
*/
import "../views"
import "net/http"

type Url struct {
	Url string
	View func(w http.ResponseWriter, r *http.Request)
}

/** Make handler functions in views package 
	Add More Urls here */
var UrlPatterns = []Url{
	Url{"/", views.Home},
	Url{"/about/", views.About},
}
