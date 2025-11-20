package main

import (
	"Power4/internal"
	"log"
	"net/http"
)

func main() {
	internal.GameInstance = internal.NewGame()

	http.HandleFunc("/", internal.StartHandler)
	http.HandleFunc("/game", internal.GameHandler)
	http.HandleFunc("/play", internal.PlayHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
