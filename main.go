package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
