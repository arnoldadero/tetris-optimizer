package main

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"4x4 board", 4, 4},
		{"5x5 board", 5, 5},
		{"Empty board", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := newBoard(tt.size)
			if board.size != tt.want {
				t.Errorf("newBoard() size = %v, want %v", board.size, tt.want)
			}
			if len(board.cells) != tt.size {
				t.Errorf("newBoard() cells length = %v, want %v", len(board.cells), tt.size)
			}
			for i := 0; i < tt.size; i++ {
				if len(board.cells[i]) != tt.size {
					t.Errorf("newBoard() row %d length = %v, want %v", i, len(board.cells[i]), tt.size)
				}
				for j := 0; j < tt.size; j++ {
					if board.cells[i][j] != '.' {
						t.Errorf("newBoard() cell [%d][%d] = %c, want .", i, j, board.cells[i][j])
					}
				}
			}
		})
	}
}

func TestBoardOperations(t *testing.T) {
	board := newBoard(4)
	tetromino, _ := newTetromino([]string{
		"...#",
		"...#",
		"...#",
		"...#",
	}, 'A')

	t.Run("Place tetromino", func(t *testing.T) {
		if !board.canPlace(tetromino, 0, 0) {
			t.Error("canPlace() should return true for valid placement")
		}
		board.place(tetromino, 0, 0)
		if board.canPlace(tetromino, 0, 0) {
			t.Error("canPlace() should return false after placement")
		}
	})

	t.Run("Remove tetromino", func(t *testing.T) {
		board.remove(tetromino, 0, 0)
		if !board.canPlace(tetromino, 0, 0) {
			t.Error("canPlace() should return true after removal")
		}
	})

	t.Run("Out of bounds placement", func(t *testing.T) {
		if board.canPlace(tetromino, -1, 0) {
			t.Error("canPlace() should return false for negative x")
		}
		if board.canPlace(tetromino, 0, -1) {
			t.Error("canPlace() should return false for negative y")
		}
		if board.canPlace(tetromino, board.size, 0) {
			t.Error("canPlace() should return false for x >= size")
		}
		if board.canPlace(tetromino, 0, board.size) {
			t.Error("canPlace() should return false for y >= size")
		}
	})
}
