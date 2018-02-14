package main

import (
	"fmt"
	"net/http"
)

var static = http.FileServer(http.Dir("files"))

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/health":
		fmt.Fprintln(w, ":)")
	default:
		w.Header().Set("Cache-Control", "private")
		static.ServeHTTP(w, r)
	}
}

func main() {
	err := http.ListenAndServe(":8989", http.HandlerFunc(handle))

	if err != nil {
		panic(err)
	}
}
