package internal

type Game struct {
	PlayerName1   string
	PlayerName2   string
	CurrentPlayer string
	PlayerColor1  string
	PlayerColor2  string
	Grid          [][]string
	Winner        string
	Mode          string
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
						switch color {
						case g.PlayerColor1:
							g.Winner = g.PlayerName1
						case g.PlayerColor2:
							g.Winner = g.PlayerName2
						default:
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

			// grille pleine -> match nul
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

}

func (g *Game) AIMove() {
	if g == nil || g.Winner != "" {
		return
	}

	aiColor := g.PlayerColor2
	playerColor := g.PlayerColor1

	// 1. Bloquer l'adversaire
	for c := 0; c < len(g.Grid[0]); c++ {
		tempGrid := copyGrid(g.Grid)
		if dropToken(tempGrid, c, playerColor) {
			if checkWinnerGrid(tempGrid, playerColor) {
				g.PlayMove(c) // bloquer le joueur
				return
			}
		}
	}

	// 2. Tenter de gagner
	for c := 0; c < len(g.Grid[0]); c++ {
		tempGrid := copyGrid(g.Grid)
		if dropToken(tempGrid, c, aiColor) {
			if checkWinnerGrid(tempGrid, aiColor) {
				g.PlayMove(c)
				return
			}
		}
	}

	// 3. Aligner 2 ou 3 jetons : choisir centre si possible
	center := len(g.Grid[0]) / 2
	if g.Grid[0][center] == "" {
		g.PlayMove(center)
		return
	}

	// 4. Choisir première colonne vide
	for c := 0; c < len(g.Grid[0]); c++ {
		if g.Grid[0][c] == "" {
			g.PlayMove(c)
			return
		}
	}
}

func copyGrid(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]string, rows)
	for r := 0; r < rows; r++ {
		newGrid[r] = make([]string, cols)
		for c := 0; c < cols; c++ {
			newGrid[r][c] = grid[r][c]
		}
	}
	return newGrid
}

func dropToken(grid [][]string, col int, color string) bool {
	for r := len(grid) - 1; r >= 0; r-- {
		if grid[r][col] == "" {
			grid[r][col] = color
			return true
		}
	}
	return false
}

func checkWinnerGrid(grid [][]string, color string) bool {
	rows := len(grid)
	cols := len(grid[0])
	dirs := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != color {
				continue
			}
			for _, d := range dirs {
				count := 1
				nr, nc := r+d[0], c+d[1]
				for nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == color {
					count++
					if count == 4 {
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
