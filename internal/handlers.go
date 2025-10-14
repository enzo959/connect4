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

type GamePageData struct {
	Grid [][]string
}

func resetHandler(w http.ResponseWriter, r *http.Request) {

}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/game.html")
	if err != nil {
		http.Error(w, "Erreur chargement game.html", http.StatusInternalServerError)
		return
	}

	game := NewGame() // pour l’instant une grille vide 6x7
	data := GamePageData{Grid: game.Grid}
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
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, StartPageData{})
}
