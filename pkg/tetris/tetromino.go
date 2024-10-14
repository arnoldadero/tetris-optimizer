package tetris

import (
	"errors"
	"fmt"
	"strings"
)

// Block represents a single block in the tetromino.
type Block struct {
	X int
	Y int
}

// Tetromino represents a tetromino shape with its blocks and rotation state.
type Tetromino struct {
	Blocks   []Block
	Rotation int  // Rotation state (0, 90, 180, 270)
	Letter   rune // Assigned letter identifier (A, B, C, etc.)
}

// Predefined standard tetromino shapes
var StandardTetrominoShapes = [][]Block{
	// I
	{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	},
	// O
	{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	},
	// T
	{
		{1, 0},
		{0, 1},
		{1, 1},
		{2, 1},
	},
	// S
	{
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
	},
	// Z
	{
		{0, 0},
		{1, 0},
		{1, 1},
		{2, 1},
	},
	// J
	{
		{0, 0},
		{0, 1},
		{1, 1},
		{2, 1},
	},
	// L
	{
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
	},
}

// Validate checks if the tetromino matches one of the standard shapes.
func (t *Tetromino) Validate() error {
	normalized := NormalizeBlocks(t.Blocks)

	for _, shape := range StandardTetrominoShapes {
		normalizedShape := NormalizeBlocks(shape)
		if CompareBlockSlices(normalized, normalizedShape) {
			return nil
		}
	}
	return errors.New("invalid tetromino shape")
}

// NormalizeBlocks shifts the blocks so that the top-left block is at (0,0).
func NormalizeBlocks(blocks []Block) []Block {
	minX, minY := blocks[0].X, blocks[0].Y
	for _, b := range blocks[1:] {
		if b.X < minX {
			minX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
	}
	normalized := make([]Block, len(blocks))
	for i, b := range blocks {
		normalized[i] = Block{X: b.X - minX, Y: b.Y - minY}
	}
	return normalized
}

// CompareBlockSlices checks if two slices of blocks are identical.
func CompareBlockSlices(a, b []Block) bool {
	if len(a) != len(b) {
		return false
	}
	blockMap := make(map[string]bool)
	for _, block := range a {
		key := fmt.Sprintf("%d,%d", block.X, block.Y)
		blockMap[key] = true
	}
	for _, block := range b {
		key := fmt.Sprintf("%d,%d", block.X, block.Y)
		if !blockMap[key] {
			return false
		}
	}
	return true
}

// Rotate rotates the tetromino 90 degrees clockwise.
func (t *Tetromino) Rotate() {
	for i, block := range t.Blocks {
		t.Blocks[i] = Block{Y: block.X, X: -block.Y}
	}
	t.Rotation = (t.Rotation + 90) % 360
	t.Blocks = NormalizeBlocks(t.Blocks)
}

// TetrominoToString returns a string representation of the tetromino.
func TetrominoToString(t Tetromino) string {
	// Determine the size of the grid
	maxX, maxY := 0, 0
	for _, block := range t.Blocks {
		if block.X > maxX {
			maxX = block.X
		}
		if block.Y > maxY {
			maxY = block.Y
		}
	}

	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, block := range t.Blocks {
		grid[block.Y][block.X] = t.Letter
	}

	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}

	return sb.String()
}
