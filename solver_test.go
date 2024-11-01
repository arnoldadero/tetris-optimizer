package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name      string
		pieces    []*Tetromino
		wantSize  int
		wantSolve bool
	}{
		{
			name: "Single I tetromino",
			pieces: []*Tetromino{
				must(newTetromino([]string{
					"...#",
					"...#",
					"...#",
					"...#",
				}, 'A')),
			},
			wantSize:  2,
			wantSolve: true,
		},
		{
			name: "Two pieces - I and Square",
			pieces: []*Tetromino{
				must(newTetromino([]string{
					"...#",
					"...#",
					"...#",
					"...#",
				}, 'A')),
				must(newTetromino([]string{
					"....",
					"..##",
					"..##",
					"....",
				}, 'B')),
			},
			wantSize:  4,
			wantSolve: true,
		},
		{
			name:      "Empty pieces",
            pieces:    []*Tetromino{},
            wantSize:  0,
            wantSolve: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            board, ok := solve(tt.pieces)
            if ok != tt.wantSolve {
                t.Errorf("solve() success = %v, want %v", ok, tt.wantSolve)
            }
            if tt.wantSolve {
                if board == nil {
                    t.Errorf("solve() returned nil board, but expected a valid solution")
                } else if board.size < tt.wantSize {
                    t.Errorf("solve() board size = %v, want >= %v", board.size, tt.wantSize)
                }
            } else {
                if board != nil {
                    t.Errorf("solve() returned non-nil board, but expected nil for unsolvable case")
                }
            }
        })
    }
}

func TestBacktrack(t *testing.T) {
	board := newBoard(4)
	pieces := []*Tetromino{
		must(newTetromino([]string{
			"...#",
			"...#",
			"...#",
			"...#",
		}, 'A')),
	}

	if !backtrack(board, pieces, 0) {
		t.Error("backtrack() should find solution for single piece")
	}

	// Test with empty board
	emptyBoard := newBoard(0)
	if backtrack(emptyBoard, pieces, 0) {
		t.Error("backtrack() should fail with size 0 board")
	}

	// Test with no pieces
	if !backtrack(board, nil, 0) {
		t.Error("backtrack() should succeed with no pieces")
	}
}

// Helper function for tests
func must(t *Tetromino, err error) *Tetromino {
	if err != nil {
		panic(err)
	}
	return t
}