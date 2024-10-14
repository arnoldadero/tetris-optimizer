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
		expected    []string
		expectError bool
	}{
		{
			name: "Valid single tetromino",
			content: "####\n" +
				"#...\n" +
				"#...\n" +
				"#...\n",
			expected: []string{
				"####\n#...\n#...\n#...\n",
			},
			expectError: false,
		},
		{
			name: "Valid multiple tetrominoes",
			content: "####\n" +
				"#...\n" +
				"#...\n" +
				"#...\n" +
				"\n" +
				"..#.\n" +
				"..#.\n" +
				"..#.\n" +
				"..#.\n",
			expected: []string{
				"####\n#...\n#...\n#...\n",
				"..#.\n..#.\n..#.\n..#.\n",
			},
			expectError: false,
		},
		{
			name:        "Empty file",
			content:     "",
			expected:    []string{},
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
					if tetromino != tt.expected[i] {
						t.Errorf("Tetromino %d mismatch.\nExpected:\n%s\nGot:\n%s", i+1, tt.expected[i], tetromino)
					}
				}
			}
		})
	}
}
