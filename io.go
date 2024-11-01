package main

import (
	"bufio"
	"fmt"
	"os"
)

func readTetrominoes(filename string) ([]*Tetromino, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pieces []*Tetromino
	scanner := bufio.NewScanner(file)
	currentPiece := make([]string, 0, 4)
	label := byte('A')

	for scanner.Scan() {
		line := scanner.Text()
		
		// Handle empty lines
		if line == "" {
			if len(currentPiece) == 0 {
				// Multiple empty lines or empty line at start
				if len(pieces) == 0 {
					return nil, fmt.Errorf("invalid format: unexpected empty line")
				}
				continue
			}
			
			if len(currentPiece) != 4 {
				return nil, fmt.Errorf("invalid tetromino size")
			}

			piece, err := newTetromino(currentPiece, label)
			if err != nil {
				return nil, err
			}
			pieces = append(pieces, piece)
			label++
			currentPiece = make([]string, 0, 4)
			continue
		}

		// Check if we're expecting a tetromino line
		if len(currentPiece) == 4 {
			return nil, fmt.Errorf("invalid format: missing empty line between tetrominoes")
		}

		// Validate line format
		if len(line) != 4 {
			return nil, fmt.Errorf("invalid line length")
		}
		for _, ch := range line {
			if ch != '.' && ch != '#' {
				return nil, fmt.Errorf("invalid character in tetromino")
			}
		}

		currentPiece = append(currentPiece, line)
	}

	// Handle last piece
	if len(currentPiece) > 0 {
		if len(currentPiece) != 4 {
			return nil, fmt.Errorf("incomplete tetromino")
		}

		piece, err := newTetromino(currentPiece, label)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, piece)
	}

	if len(pieces) == 0 {
		return nil, fmt.Errorf("no valid tetrominoes found")
	}

	return pieces, nil
}

func printBoard(board *Board) {
	for _, row := range board.cells {
		fmt.Println(string(row))
	}
}