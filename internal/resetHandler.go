package internal

import "net/http"

// ResetHandler réinitialise la partie et redirige vers le formulaire d'accueil
func ResetHandler(w http.ResponseWriter, r *http.Request) {

	GameInstance = nil
	GameStats = &Stats{}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
