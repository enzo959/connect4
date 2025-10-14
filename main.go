package main

import (
	"Power4/internal"
	"log"
	"net/http"
	"text/template"
)

type StartPageData struct {
	Message      string
	Error        string
	Player1Value string
	Player2Value string
	Color1Value  string
	Color2Value  string
}

func FullGrid() {

}

func CheckWin() {

}

func TokenMove() {

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/start.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		data := StartPageData{}
		tmpl.Execute(w, data)
		return
	}

	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	color1 := r.FormValue("color1")
	color2 := r.FormValue("color2")

	data := StartPageData{
		Message:      "Configuration reçue (fonctionnalité à venir)",
		Player1Value: player1,
		Player2Value: player2,
		Color1Value:  color1,
		Color2Value:  color2,
	}
	tmpl.Execute(w, data)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {

}

func resetHandler(w http.ResponseWriter, r *http.Request) {

}

var game *internal.Game

func main() {
	game = internal.NewGame()

	http.HandleFunc("/", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/reset", resetHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
