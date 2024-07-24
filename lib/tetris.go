package lib

import (
	"math"
	"strings"
)

type Tetrominos struct {
	Tet [][]string
}

// estimates initila grid size
func calculateInitialGridSize(tet *Tetrominos) (int, int) {
	maxWidth, maxHeight := 0, 0
	for _, tetromino := range tet.Tet {
		maxHeight = max(maxHeight, len(tetromino))
		maxWidth = max(maxWidth, len(tetromino[0]))
	}
	maxLength := int(math.Sqrt(float64(maxWidth) * float64(maxHeight)))
	return maxLength, maxLength
}

// creates a gris to be used to solve the tetrominos puzzle
func createGrid(maxWidth, maxHeight int) [][]string {
	tetSolution := make([][]string, maxHeight)
	for i := 0; i < maxHeight; i++ {
		tetSolution[i] = make([]string, maxWidth)
		for j := 0; j < maxWidth; j++ {
			tetSolution[i][j] = "."
		}
	}
	return tetSolution
}

// returns valid tetrominos and an error
func CleanTetromino(tet *Tetrominos) (*Tetrominos, error) {
	var tetrominos [][]string
	for _, tetromino := range tet.Tet {
		if !isValidTetromino(tetromino) {
			return nil, ErrInvalidTetType
		}
		tetrominos = append(tetrominos, removeDotLines(tetromino))
	}
	return &Tetrominos{Tet: tetrominos}, nil
}

// removes all empty or '.' lines vertically or horizontally
func removeDotLines(tetromino []string) []string {
	for y := 0; y < len(tetromino); y++ {
		if strings.Count(tetromino[y], ".") == len(tetromino[y]) {
			tetromino = append(tetromino[:y], tetromino[y+1:]...)
			y--
		}
	}

	for x := 0; x < len(tetromino[0]); x++ {
		isDotColumn := true
		for y := 0; y < len(tetromino); y++ {
			if tetromino[y][x] != '.' {
				isDotColumn = false
				break
			}
		}

		if isDotColumn {
			for y := 0; y < len(tetromino); y++ {
				tetromino[y] = tetromino[y][:x] + tetromino[y][x+1:]
			}
			x--
		}
	}
	return tetromino
}

// checks that the tetromino is 4x4, the non-dots are 4 in total and are connected to each
func isValidTetromino(tetromino []string) bool {
	for y := 0; y < len(tetromino); y++ {
		if len(tetromino) != 4 || len(tetromino[y]) != 4 {
			return false
		}
	}
	nonDots, connection := 0, 0
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			if tetromino[y][x] != '.' {
				nonDots++
				if y > 0 && tetromino[y-1][x] != '.' {
					connection++
				}
				if y < len(tetromino)-1 && tetromino[y+1][x] != '.' {
					connection++
				}
				if x > 0 && tetromino[y][x-1] != '.' {
					connection++
				}
				if x < len(tetromino[y])-1 && tetromino[y][x+1] != '.' {
					connection++
				}
			}
		}
	}
	if (connection < 6 || connection > 8) || nonDots != 4 {
		return false
	}

	return true
}
