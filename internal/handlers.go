package internal

import (
	"html/template"
	"net/http"
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

		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, StartPageData{})
}
