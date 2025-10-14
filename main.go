package main

import (
	"log"
	"net/http"
)

type Game struct {
	PlayerName1   string
	PlayerName2   string
	CurrentPlayer string
	PlayerColor1  string
	PlayerColor2  string
	Winner        string
}

func startHandler(w http.ResponseWriter, r *http.Request) {

}

func gameHandler(w http.ResponseWriter, r *http.Request) {

}

func resetHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/reset", resetHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
