package tetris

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadTetrominoFile reads a file containing tetromino shapes and returns a slice of strings,
// where each string represents a tetromino shape.
func ReadTetrominoFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tetrominoes []string
	var currentTetromino strings.Builder
	lineCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		if line == "" {
			if currentTetromino.Len() > 0 {
				tetrominoes = append(tetrominoes, currentTetromino.String())
				currentTetromino.Reset()
			}
		} else {
			if len(line) != 4 {
				return nil, fmt.Errorf("invalid tetromino format at line %d: line length must be 4", lineCount)
			}
			currentTetromino.WriteString(line)
			currentTetromino.WriteString("\n")
		}

		if lineCount%5 == 0 && currentTetromino.Len() > 0 {
			tetrominoes = append(tetrominoes, currentTetromino.String())
			currentTetromino.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if currentTetromino.Len() > 0 {
		tetrominoes = append(tetrominoes, currentTetromino.String())
	}

	return tetrominoes, nil
}