package main

import (
	"fmt"
	"os"
	"time"
	"tetris/lib"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("error: please provide the path to the tetromino file only")
		return
	}
	fileName := os.Args[1]
	tet, err := lib.InputFileReader(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cleanTetrominos, err := lib.CleanTetromino(tet)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	solvedTetrominos, err := lib.SolveTetris(cleanTetrominos)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range solvedTetrominos.Tet {
		fmt.Println(t)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
