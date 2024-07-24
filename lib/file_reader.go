package lib

import (
	"bufio"
	"errors"
	"os"
)

var (
	ErrInvalidTetSize = errors.New("ERROR")
	ErrInvalidTetFile = errors.New("ERROR")
	ErrInvalidTetType = errors.New("ERROR")
)

func InputFileReader(fileName string) (*Tetrominos, error) {
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
	return &Tetrominos{Tet: tetrominos}, nil
}
