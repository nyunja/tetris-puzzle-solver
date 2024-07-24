package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Tetrominos struct {
	tet [][]string
}

// type TetSolution struct {
// 	tetSolution [][]string
// }

var (
	ErrInvalidTetSize = errors.New("tetromino should have 4 lines of 4 characters each")
	ErrInvalidTetFile = errors.New("invalid tetromino file")
	ErrInvalidTetType = errors.New("invalid tetromino type")
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("error: please provide the path to the tetromino file only")
		return
	}
	fileName := os.Args[1]
	tet, err := inputFileReader(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cleanTetrominos, err := cleanTetromino(tet)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	solvedTetrominos, err := solveTetris(cleanTetrominos)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range solvedTetrominos.tet {
		fmt.Println(t)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func inputFileReader(fileName string) (*Tetrominos, error) {
	if !isValidFile(fileName) {
		return nil, ErrInvalidTetFile
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var (
		tetrominos     [][]string
		tetromino      []string
		linesRead      int
		tetrominoLabel = 'A'
	)
	for scanner.Scan() {
		line := scanner.Text()
		// if len(line) != 4 {
		// 	continue
		// }
		if tetrominoLabel > 'Z' {
			return nil, ErrInvalidTetFile
		}
		for _, ch := range line {
			if ch != '.' && ch != '#' {
				return nil, ErrInvalidTetType
			}
		}
		newLine := ""
		for _, ch := range line {
			if ch == '#' {
				newLine += string(tetrominoLabel)
			} else {
				newLine += string(ch)
			}
		}
		linesRead++
		if linesRead == 5 {
			tetrominos = append(tetrominos, tetromino)

			tetromino = nil
			linesRead = 0
			tetrominoLabel++
		} else {
			tetromino = append(tetromino, newLine)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if linesRead != 0 || len(tetrominos) == 0 {
		return nil, ErrInvalidTetFile
	}
	return &Tetrominos{tet: tetrominos}, nil
}

func calculateInitialGridSize(tet *Tetrominos) (int, int) {
	maxWidth, maxHeight := 0, 0
	for _, tetromino := range tet.tet {
		maxHeight = max(maxHeight, len(tetromino))
		maxWidth = max(maxWidth, len(tetromino[0]))
	}
	maxLength := int(math.Sqrt(float64(maxWidth)*float64(maxHeight)))
	return maxLength, maxLength
}

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

func solveTetris(tet *Tetrominos) (*Tetrominos, error) {
	var maxWidth, maxHeight int = calculateInitialGridSize(tet)
	// gridIncrement := 1
	fails := 0
	// length := int(math.Sqrt(float64(maxWidth * maxHeight)))
	for {
		tetSolution := createGrid(maxWidth, maxHeight)

		if backtrackTetris(tet.tet, tetSolution, 0) {
			return &Tetrominos{tet: tetSolution}, nil
		}
		fails++
		
		if fails > 500 {
			maxHeight++
			maxWidth++
			fails =0
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

func cleanTetromino(tet *Tetrominos) (*Tetrominos, error) {
	var tetrominos [][]string
	for _, tetromino := range tet.tet {
		if !isValidTetromino(tetromino) {
			return nil, ErrInvalidTetType
		}
		tetrominos = append(tetrominos, removeDotLines(tetromino))
	}
	return &Tetrominos{tet: tetrominos}, nil
}

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

func hasSuffix(s, suffix string) bool {
	return s[len(s)-len(suffix):] == suffix
}

func isValidFile(fileName string) bool {
	return hasSuffix(fileName, ".txt")
}
