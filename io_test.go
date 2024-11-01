package main

import (
	"os"
	"testing"
)

func TestReadTetrominoes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			name: "Valid single tetromino",
			input: "...#\n" +
				"...#\n" +
				"...#\n" +
				"...#\n",
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid multiple tetrominoes",
			input: "...#\n" +
				"...#\n" +
				"...#\n" +
				"...#\n" +
				"\n" +
				"..##\n" +
				"..##\n" +
				"....\n" +
				"....\n",
			want:    2,
			wantErr: false,
		},
		{
			name: "Invalid line length",
			input: "...#\n" +
				"...#\n" +
				"...#\n" +
				"...##\n",
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid block count",
			input: "....\n" +
				"....\n" +
				"....\n" +
				"....\n",
			want:    0,
			wantErr: true,
		},
		{
			name: "Incomplete tetromino",
			input: "...#\n" +
				"...#\n" +
				"...#\n",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpfile, err := os.CreateTemp("", "test.*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Write test input
			if _, err := tmpfile.Write([]byte(tt.input)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Test readTetrominoes
			got, err := readTetrominoes(tmpfile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("readTetrominoes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.want {
				t.Errorf("readTetrominoes() got %v pieces, want %v", len(got), tt.want)
			}
		})
	}

	// Test non-existent file
	t.Run("Non-existent file", func(t *testing.T) {
		_, err := readTetrominoes("nonexistent.txt")
		if err == nil {
			t.Error("readTetrominoes() should fail with non-existent file")
		}
	})
}