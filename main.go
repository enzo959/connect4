package main

import (
	"Power4/internal"
	"log"
	"net/http"
	"strconv"
)

func FullGrid() {

}

func CheckWin() {

}

func TokenMove() {

}

func gameHandler(w http.ResponseWriter, r *http.Request) {

}

func resetHandler(w http.ResponseWriter, r *http.Request) {

}

var game *internal.Game

func main() {
	game = internal.NewGame()

	http.HandleFunc("/", internal.StartHandler)
	//http.HandleFunc("/", gameHandler)
	//http.HandleFunc("/game", gameHandler)
	//http.HandleFunc("/reset", resetHandler)

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		colParam := r.URL.Query().Get("col")
		if colParam == "" {
			w.Write([]byte("Ajoute un coup avec ?col=numero\n"))
			return
		}

		col, err := strconv.Atoi(colParam)
		if err != nil || col < 0 || col >= len(game.Grid[0]) {
			w.Write([]byte("Colonne invalide\n"))
			return
		}

		// Déterminer le symbole du joueur
		symbol := "X"
		if game.CurrentPlayer == "X" {
			symbol = "O"
		}

		// Jouer dans la colonne
		for i := len(game.Grid) - 1; i >= 0; i-- {
			if game.Grid[i][col] == "" {
				game.Grid[i][col] = symbol
				game.CurrentPlayer = symbol
				break
			}
		}

		// Réafficher la grille
		for _, row := range game.Grid {
			for _, cell := range row {
				if cell == "" {
					w.Write([]byte(". "))
				} else {
					w.Write([]byte(cell + " "))
				}
			}
			w.Write([]byte("\n"))
		}
	})

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
