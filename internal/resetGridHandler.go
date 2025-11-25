package internal

import "net/http"

func ResetGridHandler(w http.ResponseWriter, r *http.Request) {
	// sécurité : si le jeu n'existe plus
	if GameInstance == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	GameInstance.Grid = NewGrid(6, 7)
	GameInstance.Winner = ""
	GameInstance.CurrentPlayer = GameInstance.PlayerName1

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
