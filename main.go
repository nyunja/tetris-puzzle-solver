package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(tet)
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
		if len(line) != 4 {
			continue
		}
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
		tetromino = append(tetromino, newLine)
		if linesRead == 4 {
			tetrominos = append(tetrominos, tetromino)
			tetromino = nil
			linesRead = 0
			tetrominoLabel++
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
	return maxWidth, maxHeight
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
					for j := 0; j < len(tetromino[0]); j++ {
						if k+len(tetromino) > len(tetSolution) || m+len(tetromino[0]) > len(tetromino[0]) {
							break
						}
						// if tetromino[i][j] != '.' && tetSolution[k+i][m+j] != "." {
						// 	continue
						// }
						if tetromino[i][j] != '.' && tetSolution[k+i][m+j] == "." {
							tetSolution[k+i][m+j] = string(tetromino[i][j])
						}
						// fmt.Printf("%s %s\n", string(tetromino[i][j]), tetSolution[k+i][m+j])
					}
				}
			}
		}
	}
	return true
}

func solveTetris(tet *Tetrominos) (*Tetrominos, error) {
	var maxWidth, maxHeight int = calculateInitialGridSize(tet)
	gridIncrement := 2
	// fails := 0
	// length := int(math.Sqrt(float64(maxWidth * maxHeight)))
	for {
		tetSolution := createGrid(maxWidth, maxHeight)

		if backtrackTetris(tet.tet, tetSolution, 0) {
			return &Tetrominos{tet: tetSolution}, nil
		}
		// fails++
		// if fails > 5 {
		// 	gridIncrement *= 2
		// 	fails =0
		// }
		gridIncrement *= 2
		maxWidth += gridIncrement
		maxHeight += gridIncrement
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
	fmt.Println(tetromino)
	for y := 0; y < len(tetromino); y++ {
		if strings.Count(tetromino[y], ".") == len(tetromino[y]) {
			tetromino = append(tetromino[:y], tetromino[y+1:]...)
			y--
		}
	}
	fmt.Println("horizontal: ", tetromino)

	for x := 0; x < len(tetromino[0]); x++ {
		fmt.Println(tetromino)
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
	connection := 0
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			if tetromino[y][x] != '.' {
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

	if connection >= 6 && connection <= 8 {
		return true
	}
	return false
}

func hasSuffix(s, suffix string) bool {
	return s[len(s)-len(suffix):] == suffix
}

func isValidFile(fileName string) bool {
	return hasSuffix(fileName, ".txt")
}
