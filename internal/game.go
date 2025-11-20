package internal

type Game struct {
	PlayerName1   string
	PlayerName2   string
	CurrentPlayer string
	PlayerColor1  string
	PlayerColor2  string
	Grid          [][]string
	Winner        string
}

var GameInstance *Game

type Stats struct {
	WinsPlayer1 int
	WinsPlayer2 int
	Draws       int
}

var GameStats = &Stats{} // statistiques pour la partie actuelle

// Creat grid
func NewGrid(rows, cols int) [][]string {
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = ""
		}
	}
	return grid
}

func NewGame() *Game {
	return &Game{
		PlayerName1:   "",
		PlayerName2:   "",
		PlayerColor1:  "",
		PlayerColor2:  "",
		CurrentPlayer: "",
		Grid:          NewGrid(6, 7),
		Winner:        "",
	}
}

func (g *Game) IsFull() bool {
	for r := 0; r < len(g.Grid); r++ {
		for c := 0; c < len(g.Grid[0]); c++ {
			if g.Grid[r][c] == "" {
				return false
			}
		}
	}
	return true
}

func (g *Game) CheckWinner() bool {
	if g == nil {
		return false
	}

	rows := len(g.Grid)
	if rows == 0 {
		return false
	}
	cols := len(g.Grid[0])

	dirs := [][2]int{
		{0, 1},  // droite
		{1, 0},  // bas
		{1, 1},  // diag bas-droite
		{1, -1}, // diag bas-gauche
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := g.Grid[r][c]
			if color == "" {
				continue
			}

			for _, d := range dirs {
				count := 1
				nr, nc := r+d[0], c+d[1]

				for nr >= 0 && nr < rows && nc >= 0 && nc < cols && g.Grid[nr][nc] == color {
					count++
					if count == 4 {
						if color == g.PlayerColor1 {
							g.Winner = g.PlayerName1
						} else if color == g.PlayerColor2 {
							g.Winner = g.PlayerName2
						} else {
							g.Winner = "Inconnu"
						}
						return true
					}
					nr += d[0]
					nc += d[1]
				}
			}
		}
	}

	return false
}

func (g *Game) PlayMove(col int) {
	if g == nil || g.Winner != "" {
		return
	}

	if col < 0 || col >= len(g.Grid[0]) {
		return
	}

	// trouver la première case vide en partant du bas
	for r := len(g.Grid) - 1; r >= 0; r-- {
		if g.Grid[r][col] == "" {

			// déterminer la couleur selon CurrentPlayer
			var color string
			if g.CurrentPlayer == g.PlayerName1 {
				color = g.PlayerColor1
			} else {
				color = g.PlayerColor2
			}

			g.Grid[r][col] = color

			// vérifier victoire
			if g.CheckWinner() {
				switch g.Winner {
				case g.PlayerName1:
					GameStats.WinsPlayer1++
				case g.PlayerName2:
					GameStats.WinsPlayer2++
				}
				return
			}

			// grille pleine => match nul
			if g.IsFull() {
				g.Winner = "draw"
				GameStats.Draws++
				return
			}

			// changer le joueur
			if g.CurrentPlayer == g.PlayerName1 {
				g.CurrentPlayer = g.PlayerName2
			} else {
				g.CurrentPlayer = g.PlayerName1
			}

			return
		}
	}

	// colonne pleine -> on ne fait rien
}
