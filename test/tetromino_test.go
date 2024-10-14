package test

import (
	"testing"

	"github.com/arnoldadero/tetris-optimizer/pkg/tetris"
)

func TestTetrominoValidation(t *testing.T) {
	validShapes := [][]tetris.Block{
		tetris.StandardTetrominoShapes[0], // I
		tetris.StandardTetrominoShapes[1], // O
		tetris.StandardTetrominoShapes[2], // T
		tetris.StandardTetrominoShapes[3], // S
		tetris.StandardTetrominoShapes[4], // Z
		tetris.StandardTetrominoShapes[5], // J
		tetris.StandardTetrominoShapes[6], // L
	}

	for i, shape := range validShapes {
		tetromino := tetris.Tetromino{
			Blocks: shape,
		}
		if err := tetromino.Validate(); err != nil {
			t.Errorf("Valid shape %d failed validation: %v", i, err)
		}
	}

	invalidShapes := [][]tetris.Block{
		// Five blocks
		{
			{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0},
		},
		// Disconnected blocks
		{
			{X: 0, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 1}, {X: 4, Y: 2},
		},
		// Non-standard shape
		{
			{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2},
		},
	}

	for i, shape := range invalidShapes {
		tetromino := tetris.Tetromino{
			Blocks: shape,
		}
		if err := tetromino.Validate(); err == nil {
			t.Errorf("Invalid shape %d passed validation", i)
		}
	}
}

func TestTetrominoRotation(t *testing.T) {
	// Test rotating the I tetromino
	tetromino := tetris.Tetromino{
		Blocks: []tetris.Block{
			{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0},
		},
		Rotation: 0,
		Letter:   'A',
	}

	expectedAfterRotation := []tetris.Block{
		{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3},
	}

	tetromino.Rotate()

	if tetromino.Rotation != 90 {
		t.Errorf("Expected rotation to be 90, got %d", tetromino.Rotation)
	}

	normalized := tetris.NormalizeBlocks(tetromino.Blocks)
	expectedNormalized := tetris.NormalizeBlocks(expectedAfterRotation)

	if !tetris.CompareBlockSlices(normalized, expectedNormalized) {
		t.Errorf("Rotation failed. Expected %v, got %v", expectedNormalized, normalized)
	}
}