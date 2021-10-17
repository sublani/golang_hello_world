package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	http.ListenAndServe("direccion", nil)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
