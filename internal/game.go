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

func NewGame(rows, cols int) *Game {
	return &Game{
		Player1Name:   "",
		Player2Name:   "",
		Player1Color:  "",
		Player2Color:  "",
		CurrentPlayer: "",
		Grid:          NewGrid(rows, cols),
		Winner:        "",
	}
}
