package stateless_front

import (
	"net/http"
)

func run() {
	port := "4500"

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/post", post)
	mux.HandleFunc("/read", read)
}
