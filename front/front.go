package front

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Baca file HTML
	http.ServeFile(w, r, "index.html")

}
