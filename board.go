package main

type Board struct {
	size  int
	cells [][]byte
}

func newBoard(size int) *Board {
	cells := make([][]byte, size)
	for i := range cells {
		cells[i] = make([]byte, size)
		for j := range cells[i] {
			cells[i][j] = '.'
		}
	}
	return &Board{size: size, cells: cells}
}

func (b *Board) canPlace(t *Tetromino, x, y int) bool {
	for _, block := range t.blocks {
		newX, newY := x+block.x, y+block.y
		if newX < 0 || newX >= b.size || newY < 0 || newY >= b.size {
			return false
		}
		if b.cells[newY][newX] != '.' {
			return false
		}
	}
	return true
}

func (b *Board) place(t *Tetromino, x, y int) {
	for _, block := range t.blocks {
		b.cells[y+block.y][x+block.x] = t.label
	}
}

func (b *Board) remove(t *Tetromino, x, y int) {
	for _, block := range t.blocks {
		b.cells[y+block.y][x+block.x] = '.'
	}
}
