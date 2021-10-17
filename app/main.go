package main

import (
	"fmt"      // Imprimir en consola
	"io"       // Ayuda a escribir en la respuesta
	"log"      //Loguear si algo sale mal
	"net/http" // El paquete HTTP
)

func main() {

	http.HandleFunc("/prueba1", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste prueba1")
	})

	http.HandleFunc("/prueba2", func(w http.Response1Writer, peticion *http.Request) {
		io.WriteString(w, "Solicitaste prueba2")
	})

	http.HandleFunc("/prueba3", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste prueba3")
	})

	http.HandleFunc("/prueba4", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste prueba4")
	})

	http.HandleFunc("/prueba5", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste prueba5")
	})

	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
