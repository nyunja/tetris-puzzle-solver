package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: please provide the path to the tetromino file only")
		return
	}
	path := "/files/static/"
	fileName, err := getFilePath(path)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	fmt.Println(fileName)
}

// func hasSuffix(s, suffix string) bool {
// 	return s[len(s)-len(suffix):] == suffix
// }

func getFilePath(s string) (string, error) {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	if s == "" {
		return "", errors.New("invalid URL")
	}
	return s + ".txt", nil
}
