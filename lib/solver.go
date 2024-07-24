package lib

func backtrackTetris(tetrominos [][]string, tetSolution [][]string, index int) bool {
	if index == len(tetrominos) {
		return true // all tetrominos are full
	}
	tetromino := tetrominos[index]
	for k := 0; k < len(tetSolution); k++ {
		for m := 0; m < len(tetSolution); m++ {
			if fits(tetromino, tetSolution, m, k) {
				for i := 0; i < len(tetromino); i++ {
					for j := 0; j < len(tetromino[i]); j++ {
						if tetromino[i][j] != '.' {
							tetSolution[k+i][m+j] = string(tetromino[i][j])
							// fmt.Println(k+i, m+j, tetSolution[k+i][m+j])
						}
					}
				}
				if backtrackTetris(tetrominos, tetSolution, index+1) {
					return true
				}
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

func SolveTetris(tet *Tetrominos) (*Tetrominos, error) {
	var maxWidth, maxHeight int = calculateInitialGridSize(tet)
	// gridIncrement := 1
	fails := 0
	// length := int(math.Sqrt(float64(maxWidth * maxHeight)))
	for {
		tetSolution := createGrid(maxWidth, maxHeight)

		if backtrackTetris(tet.Tet, tetSolution, 0) {
			return &Tetrominos{Tet: tetSolution}, nil
		}
		fails++

		if fails > 500 {
			maxHeight++
			maxWidth++
			fails = 0
		}
		// if fails > 10 {
		// 	maxWidth += gridIncrement
		// 	// fails =0
		// }
	}
}

func fits(tetromino []string, tetSolution [][]string, x, y int) bool {
	// check if the tetromino can fit in the space left in the tet solution
	if y+len(tetromino) > len(tetSolution) || x+len(tetromino[0]) > len(tetSolution[0]) {
		return false
	}

	// chcek for overlaps
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] != '.' && tetSolution[y+i][x+j] != "." {
				return false
			}
		}
	}
	return true
}
