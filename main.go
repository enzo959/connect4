package main

import (
	"log"
	"net/http"
	"text/template"
)

type Game struct {
	PlayerName1   string
	PlayerName2   string
	CurrentPlayer string
	PlayerColor1  string
	PlayerColor2  string
	Grid          [][]string
	Winner        string
}

type StartPageData struct {
	Message      string
	Error        string
	Player1Value string
	Player2Value string
	Color1Value  string
	Color2Value  string
}

// Creat grid
func NewGrid(rows, cols int) [][]string {
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = ""
		}
	}
	return grid
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

func main() {
	http.HandleFunc("/", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/reset", resetHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
