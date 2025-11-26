package internal

func (g *Game) AIMove() {
	if g == nil || g.Winner != "" {
		return
	}

	tryMove := func(color string) bool {
		for c := 0; c < len(g.Grid[0]); c++ {
			temp := copyGrid(g.Grid)
			if dropToken(temp, c, color) && checkWinnerGrid(temp, color) {
				g.PlayMove(c)
				return true
			}
		}
		return false
	}

	aiColor := g.PlayerColor2
	playerColor := g.PlayerColor1

	// 1. Bloquer ou 2. Gagner
	if tryMove(playerColor) || tryMove(aiColor) {
		return
	}

	// 3. Jouer au centre si libre
	center := len(g.Grid[0]) / 2
	if g.Grid[0][center] == "" {
		g.PlayMove(center)
		return
	}

	// 4. Première colonne vide
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
