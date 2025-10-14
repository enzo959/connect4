package internal

import (
	"html/template"
	"net/http"
)

type StartPageData struct {
	Message      string
	Player1Value string
	Player2Value string
	Color1Value  string
	Color2Value  string
}

func resetHandler(w http.ResponseWriter, r *http.Request) {

}

func gameHandler(w http.ResponseWriter, r *http.Request) {

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

		data := StartPageData{
			Message:      "Configuration reçue (fonctionnalité à venir)",
			Player1Value: player1,
			Player2Value: player2,
			Color1Value:  color1,
			Color2Value:  color2,
		}

		tmpl.Execute(w, data)
		return
	}

	tmpl.Execute(w, StartPageData{})
}
