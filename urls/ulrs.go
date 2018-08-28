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
	Url  string
	View func(w http.ResponseWriter, r *http.Request)
}


var UrlPatterns = []Url{
	Url{"/", views.Home},
	Url{"/about", views.About},
	Url{"/number/{number}", views.Number},
}
