package internal

import (
	"net/http"
	"strconv"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	var colStr string
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erreur lecture formulaire", http.StatusBadRequest)
			return
		}
		colStr = r.FormValue("col")
	} else {
		colStr = r.URL.Query().Get("col")
	}

	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Error(w, "Colonne invalide", http.StatusBadRequest)
		return
	}

	GameInstance.PlayMove(col)
	if GameInstance.Mode == "solo" && GameInstance.CurrentPlayer == GameInstance.PlayerName2 && GameInstance.Winner == "" {
		GameInstance.AIMove()
	}

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
