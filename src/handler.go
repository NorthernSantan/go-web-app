package main

import (
	"fmt"
	"net/http"
)

func Handler() {

	fmt.Printf("Curretnly in web server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
