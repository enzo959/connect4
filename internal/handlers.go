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
	WinsPlayer1   int
	WinsPlayer2   int
	Draws         int
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
		WinsPlayer1:   GameStats.WinsPlayer1,
		WinsPlayer2:   GameStats.WinsPlayer2,
		Draws:         GameStats.Draws,
	}

	tmpl.Execute(w, data)
}
