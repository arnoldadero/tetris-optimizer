package tetris

import (
	"bufio"
	"fmt"
	"os"
)

// ReadTetrominoFile reads a file containing tetromino shapes and returns a slice of Tetromino structs.
func ReadTetrominoFile(filePath string) ([]Tetromino, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tetrominoes []Tetromino
	var currentTetrominoLines []string
	lineCount := 0
	tetrominoCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		if line == "" {
			if len(currentTetrominoLines) > 0 {
				tetromino, err := parseTetromino(currentTetrominoLines)
				if err != nil {
					return nil, fmt.Errorf("error parsing tetromino at lines %d-%d: %w", lineCount-len(currentTetrominoLines), lineCount, err)
				}
				tetromino.Letter = rune('A' + tetrominoCount)
				tetrominoes = append(tetrominoes, tetromino)
				tetrominoCount++
				currentTetrominoLines = []string{}
			}
		} else {
			if len(line) != 4 {
				return nil, fmt.Errorf("invalid tetromino format at line %d: line length must be 4", lineCount)
			}
			currentTetrominoLines = append(currentTetrominoLines, line)
		}

		// Automatically finalize tetromino after 4 lines
		if len(currentTetrominoLines) == 4 {
			tetromino, err := parseTetromino(currentTetrominoLines)
			if err != nil {
				return nil, fmt.Errorf("error parsing tetromino at lines %d-%d: %w", lineCount-3, lineCount, err)
			}
			tetromino.Letter = rune('A' + tetrominoCount)
			tetrominoes = append(tetrominoes, tetromino)
			tetrominoCount++
			currentTetrominoLines = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Handle any remaining tetromino
	if len(currentTetrominoLines) > 0 {
		if len(currentTetrominoLines) != 4 {
			return nil, fmt.Errorf("incomplete tetromino at end of file")
		}
		tetromino, err := parseTetromino(currentTetrominoLines)
		if err != nil {
			return nil, fmt.Errorf("error parsing tetromino at end of file: %w", err)
		}
		tetromino.Letter = rune('A' + tetrominoCount)
		tetrominoes = append(tetrominoes, tetromino)
	}

	return tetrominoes, nil
}

// parseTetromino converts tetromino lines into a Tetromino struct.
func parseTetromino(lines []string) (Tetromino, error) {
	var blocks []Block
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				blocks = append(blocks, Block{X: x, Y: y})
			} else if char != '.' {
				return Tetromino{}, fmt.Errorf("invalid character '%c' in tetromino", char)
			}
		}
	}
	if len(blocks) != 4 {
		return Tetromino{}, fmt.Errorf("invalid number of blocks: expected 4, got %d", len(blocks))
	}

	tetromino := Tetromino{
		Blocks: blocks,
	}

	// Validate tetromino shape
	if err := tetromino.Validate(); err != nil {
		return Tetromino{}, err
	}

	return tetromino, nil
}
