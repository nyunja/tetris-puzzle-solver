package lib

// places the tetrominos in the final solution and removes them if they dont fit
func backtrackTetris(tetrominos [][]string, tetSolution [][]string, index int) bool {
	if index == len(tetrominos) {
		return true // all tetrominos are full
	}
	tetromino := tetrominos[index]
	for k := 0; k < len(tetSolution); k++ {
		for m := 0; m < len(tetSolution); m++ {
			if fits(tetromino, tetSolution, m, k) {
				// add tetromino
				for i := 0; i < len(tetromino); i++ {
					for j := 0; j < len(tetromino[i]); j++ {
						if tetromino[i][j] != '.' {
							tetSolution[k+i][m+j] = string(tetromino[i][j])
						}
					}
				}
				// recursion to add subsequent tetrominos
				if backtrackTetris(tetrominos, tetSolution, index+1) {
					return true
				}
				// removes tetrominos when backtrack is false
				for i := 0; i < len(tetromino); i++ {
					for j := 0; j < len(tetromino[i]); j++ {
						if tetromino[i][j] != '.' {
							tetSolution[k+i][m+j] = "."
						}
					}
				}
			}
		}
	}
	return false
}

// creates initial grid and resizes it if the tetrominos dont fit
func SolveTetris(tet *Tetrominos) (*Tetrominos, error) {
	var maxWidth, maxHeight int = calculateInitialGridSize(tet)
	fails := 0
	for {
		tetSolution := createGrid(maxWidth, maxHeight)

		if backtrackTetris(tet.Tet, tetSolution, 0) {
			return &Tetrominos{Tet: tetSolution}, nil
		}
		fails++
		// increases grid size when failsed attempts are more than 500
		if fails > 500 {
			maxHeight++
			maxWidth++
			fails = 0
		}
	}
}

// checks if a tetromino fits at a particular point
func fits(tetromino []string, tetSolution [][]string, x, y int) bool {
	// check if the tetromino can fit in the space left in the tet solution
	if y+len(tetromino) > len(tetSolution) || x+len(tetromino[0]) > len(tetSolution[0]) {
		return false
	}

	// check for overlaps
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] != '.' && tetSolution[y+i][x+j] != "." {
				return false
			}
		}
	}
	return true
}
