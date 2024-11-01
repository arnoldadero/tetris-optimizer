package main

import (
	"reflect"
	"testing"
)

func TestNewTetromino(t *testing.T) {
	tests := []struct {
		name    string
		pattern []string
		label   byte
		want    *Tetromino
		wantErr bool
	}{
		{
			name: "Valid I tetromino",
			pattern: []string{
				"...#",
				"...#",
				"...#",
				"...#",
			},
			label: 'A',
			want: &Tetromino{
				blocks: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
				label:  'A',
			},
			wantErr: false,
		},
		{
			name: "Valid square tetromino",
			pattern: []string{
				"....",
				"..##",
				"..##",
				"....",
			},
			label: 'B',
			want: &Tetromino{
				blocks: []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
				label:  'B',
			},
			wantErr: false,
		},
		{
			name: "Invalid size",
			pattern: []string{
				"...",
				"...",
				"...",
			},
			label:   'C',
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid block count",
			pattern: []string{
				"....",
				"..#.",
				"..#.",
				"....",
			},
			label:   'D',
			want:    nil,
			wantErr: true,
		},
		{
			name: "Disconnected blocks",
			pattern: []string{
				"#..#",
				"....",
				"#..#",
				"....",
			},
			label:   'E',
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTetromino(tt.pattern, tt.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTetromino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidTetromino(t *testing.T) {
	tests := []struct {
		name   string
		blocks []Point
		want   bool
	}{
		{
			name:   "Valid I shape",
			blocks: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			want:   true,
		},
		{
			name:   "Valid square shape",
			blocks: []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			want:   true,
		},
		{
			name:   "Invalid - disconnected",
			blocks: []Point{{0, 0}, {0, 1}, {2, 0}, {2, 1}},
			want:   false,
		},
		{
			name:   "Invalid - wrong count",
			blocks: []Point{{0, 0}, {0, 1}, {0, 2}},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidTetromino(tt.blocks); got != tt.want {
				t.Errorf("isValidTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}
