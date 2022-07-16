package stateless_front

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Main page type in login or read in url</h1>"))
}
