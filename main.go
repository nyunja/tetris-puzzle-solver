package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Tetrominos struct {
	tet [][]string
}

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
		tetrominos [][]string
		tetromino []string
		linesRead int
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
			if ch =='#' {
				newLine += string(tetrominoLabel)
			} else {
				newLine += string(ch)
			}
		}
		fmt.Println(newLine)
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
	if linesRead != 0 {
		return nil, ErrInvalidTetFile
	}
	return &Tetrominos{tet: tetrominos}, nil
}



func hasSuffix(s, suffix string) bool {
	return s[len(s)-len(suffix):] == suffix
}

func isValidFile(fileName string) bool {
	return hasSuffix(fileName, ".txt")
}
