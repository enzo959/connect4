package internal

import (
	"html/template"
	"net/http"
)

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
		mode := r.FormValue("mode")

		var message string
		if player1 == player2 {
			message = "Les deux joueurs ne peuvent pas avoir le même nom."
		} else if color1 == color2 {
			message = "Les deux joueurs ne peuvent pas choisir la même couleur."
		}

		if message != "" {
			// Affiche le formulaire avec le message d'erreur et valeurs déjà saisies
			tmpl.Execute(w, StartPageData{
				Message:      message,
				PlayerValue1: player1,
				PlayerValue2: player2,
				ColorValue1:  color1,
				ColorValue2:  color2,
			})
			return
		}

		GameInstance = NewGame()
		GameInstance.PlayerName1 = player1
		GameInstance.PlayerName2 = player2
		GameInstance.PlayerColor1 = color1
		GameInstance.PlayerColor2 = color2
		GameInstance.CurrentPlayer = player1
		GameInstance.Mode = mode
		GameInstance.Grid = NewGrid(6, 7)

		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, StartPageData{})
}
