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
