package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_file.txt>")
		fmt.Println("ERROR: Invalid number of arguments")
		return
	}

	pieces, err := readTetrominoes(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	board, ok := solve(pieces)
	if !ok {
		fmt.Println("ERROR")
		return
	}

	printBoard(board)
}
