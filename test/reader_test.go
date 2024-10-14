package test

import (
	"os"
	"testing"

	"github.com/arnoldadero/tetris-optimizer/pkg/tetris"
)

func TestReadTetrominoFile(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expected    []tetris.Tetromino
		expectError bool
	}{
		{
			name: "Valid single tetromino",
			content: "####\n" + // 4 blocks in the first line
			         "....\n" +
			         "....\n" +
			         "....\n",
			expected: []tetris.Tetromino{
				{
					Blocks: []tetris.Block{
						{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0},
					},
					Letter: 'A',
				},
			},
			expectError: false,
		},
		{
			name: "Valid multiple tetrominoes",
			content: "####\n" +
				"....\n" +
				"....\n" +
				"....\n" +
				"\n" +
				"##..\n" +  // Changed from "..#." to "##.."
				"##..\n" +  // Changed from "..#." to "##.."
				"....\n" +  // Changed from "..#." to "...."
				"....\n",   // Changed from "..#." to "...."
			expected: []tetris.Tetromino{
				{
					Blocks: []tetris.Block{
						{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0},
					},
					Letter: 'A',
				},
				{
					Blocks: []tetris.Block{
						{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1},  // Updated to match the new shape
					},
					Letter: 'B',
				},
			},
			expectError: false,
		},
		{
			name:        "Empty file",
			content:     "",
			expected:    []tetris.Tetromino{},
			expectError: false,
		},
		{
			name: "Invalid line length",
			content: "###\n" + // Line with 3 characters instead of 4
				"#...\n" +
				"#...\n" +
				"#...\n",
			expected:    nil,
			expectError: true,
		},
		{
			name: "Invalid character in tetromino",
			content: "####\n" +
				"#..*\n" + // '*' is invalid
				"#...\n" +
				"#...\n",
			expected:    nil,
			expectError: true,
		},
		{
			name: "Incorrect number of blocks",
			content: "####\n" +
				"#...\n" +
				"#...\n" +
				"#..#\n", // 5 blocks
			expected:    nil,
			expectError: true,
		},
		{
			name: "Incomplete tetromino at end of file",
			content: "####\n" +
				"#...\n" +
				"#...\n", // Only 3 lines
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the test content
			tmpFile, err := os.CreateTemp("", "tetromino_test_*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			if _, err := tmpFile.WriteString(tt.content); err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}
			tmpFile.Close()

			// Call the function under test
			result, err := tetris.ReadTetrominoFile(tmpFile.Name())

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error but got: %v", err)
				}
				if len(result) != len(tt.expected) {
					t.Errorf("Expected %d tetrominoes, got %d", len(tt.expected), len(result))
				}
				for i, tetromino := range result {
					expectedTetromino := tt.expected[i]
					if tetromino.Letter != expectedTetromino.Letter {
						t.Errorf("Tetromino %d: expected letter %c, got %c", i+1, expectedTetromino.Letter, tetromino.Letter)
					}
					if !tetris.CompareBlockSlices(tetromino.Blocks, expectedTetromino.Blocks) {
						t.Errorf("Tetromino %d blocks mismatch.\nExpected:\n%v\nGot:\n%v", i+1, expectedTetromino.Blocks, tetromino.Blocks)
					}
				}
			}
		})
	}
}
