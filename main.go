package main

import (
	"Power4/internal"
	"log"
	"net/http"
)

func FullGrid() {

}

func CheckWin() {

}

func TokenMove() {

}

var game *internal.Game

func main() {
	internal.GameInstance = internal.NewGame()

	http.HandleFunc("/", internal.StartHandler)
	http.HandleFunc("/game", internal.GameHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
