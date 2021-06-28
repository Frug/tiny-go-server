package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Printf("Serving on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
