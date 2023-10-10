package life

import "fmt"

type Rule int

const (
	none Rule = iota
	lonely
	overpopulation
	birth
)

type Cell struct {
	Current         bool
	Next            bool
	DebugNeighbours int
	Rule            Rule
}

type World struct {
	Grid [][]Cell
}

func CreateWorld(startGrid string) *World {
	var grid [][]Cell             // this will have length 0
	grid = append(grid, []Cell{}) // initialise first row

	maxColumn := 120

	row := 0
	column := 0

	for _, c := range startGrid {
		if c == '\n' {
			// fill in missing cells at the end of the row
			missingCells := maxColumn - column
			for missingCells > 0 {
				grid[row] = append(grid[row], Cell{false, false, 0, none})
				column++
				missingCells--
			}

			// now create a new row
			row += 1
			grid = append(grid, []Cell{})
			column = 0
		} else {
			cell := Cell{c != ' ', false, 0, none}
			grid[row] = append(grid[row], cell)
			column += 1
		}
	}

	// fill in missing cells at the end of the final row
	missingCells := maxColumn - column
	for missingCells > 0 {
		grid[row] = append(grid[row], Cell{false, false, 0, 0})
		column++
		missingCells--
	}

	return &World{Grid: grid}
}

func (w *World) NextGeneration() {

	// Loop through every cell
	for y, row := range w.Grid {
		for x, _ := range row {

			// finding number of neighbours that are alive
			aliveNeighbours := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {

					// make sure the cell we're checking exists
					if y+j >= 0 && y+j < len(w.Grid) {
						if x+i >= 0 && x+i < len(w.Grid[y+j]) {
							if w.Grid[y+j][x+i].Current {
								aliveNeighbours += 1
							}
						}
					}
				}
			}
			if w.Grid[y][x].Current {
				aliveNeighbours -= 1
			}

			// Implement the Rules of Life
			w.Grid[y][x].DebugNeighbours = aliveNeighbours

			if w.Grid[y][x].Current && aliveNeighbours < 2 {
				// Cell is lonely and dies
				w.Grid[y][x].Next = false
				w.Grid[y][x].Rule = lonely
			} else if w.Grid[y][x].Current && aliveNeighbours > 3 {
				// Cell dies due to over population
				w.Grid[y][x].Next = false
				w.Grid[y][x].Rule = overpopulation
			} else if aliveNeighbours == 3 {
				// A new cell is born
				w.Grid[y][x].Next = true
				w.Grid[y][x].Rule = birth
			} else {
				// Remains the same
				w.Grid[y][x].Next = w.Grid[y][x].Current
				w.Grid[y][x].Rule = none
			}
		}

		// finally, after applying all the rules, set the current state to the next state
	}
	for y, row := range w.Grid {
		for x, _ := range row {
			w.Grid[y][x].Current = w.Grid[y][x].Next
		}
	}
}

func (w *World) Print() string {
	alive := 0
	output := ""
	//output := "LIFE\n\n" + fmt.Sprintf("%v rows\n\n", len(w.Grid))

	for _, row := range w.Grid {
		for _, cell := range row {
			switch cell.Rule {
			case none:
				if cell.Current {
					output += "+"
				} else {
					output += " "
				}
			case lonely:
				output += "/"
			case overpopulation:
				output += "·"
			case birth:
				output += "#"
			}
			alive++
		}
		output += "\n"
	}
	//output += fmt.Sprintf("%v alive", alive)
	return output
}

func (w *World) PrintSVG() string {
	output := ""
	for y, row := range w.Grid {
		for x, cell := range row {
			if cell.Current {
				output += fmt.Sprintf(`<rect fill="#ffffff" x="%v" y="%v" height="1" width="1" />`, x, y)
			}
		}
	}
	return output
}

func main() {
}
