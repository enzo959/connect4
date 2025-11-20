package internal

import (
	"html/template"
	"net/http"
	"strconv"
)

type StartPageData struct {
	Message      string
	PlayerValue1 string
	PlayerValue2 string
	ColorValue1  string
	ColorValue2  string
}

type GamePageData struct {
	Grid          [][]string
	PlayerName1   string
	PlayerName2   string
	PlayerColor1  string
	PlayerColor2  string
	CurrentPlayer string
	Winner        string
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/game.html")
	if err != nil {
		http.Error(w, "Erreur chargement game.html", http.StatusInternalServerError)
		return
	}

	data := GamePageData{
		Grid:          GameInstance.Grid,
		PlayerName1:   GameInstance.PlayerName1,
		PlayerName2:   GameInstance.PlayerName2,
		PlayerColor1:  GameInstance.PlayerColor1,
		PlayerColor2:  GameInstance.PlayerColor2,
		CurrentPlayer: GameInstance.CurrentPlayer,
		Winner:        GameInstance.Winner,
	}

	tmpl.Execute(w, data)
}

// Affiche le formulaire d'accueil
func StartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/start.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		player1 := r.FormValue("player1")
		player2 := r.FormValue("player2")
		color1 := r.FormValue("color1")
		color2 := r.FormValue("color2")

		GameInstance.PlayerName1 = player1
		GameInstance.PlayerName2 = player2
		GameInstance.PlayerColor1 = color1
		GameInstance.PlayerColor2 = color2
		GameInstance.CurrentPlayer = player1

		GameInstance.Grid = NewGrid(6, 7)

		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	// Réinitialiser la partie
	GameInstance = NewGame()

	tmpl.Execute(w, StartPageData{})
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	colStr := r.URL.Query().Get("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Error(w, "Colonne invalide", http.StatusBadRequest)
		return
	}

	GameInstance.PlayMove(col)

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// ResetHandler réinitialise la partie et redirige vers le formulaire d'accueil
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	GameInstance = NewGame()                      // réinitialise la partie
	http.Redirect(w, r, "/", http.StatusSeeOther) // renvoie vers la page d'accueil
}
