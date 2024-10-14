package main

import (
	"fmt"
	"os"

	"github.com/arnoldadero/tetris-optimizer/pkg/tetris"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: tetris-optimizer <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	tetrominoes, err := tetris.ReadTetrominoFile(inputFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d tetrominoes:\n", len(tetrominoes))
	for _, tetromino := range tetrominoes {
		fmt.Printf("Tetromino %c:\n%s\n", tetromino.Letter, tetris.TetrominoToString(tetromino))
	}
}