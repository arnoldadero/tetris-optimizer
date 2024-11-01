package main

import "math"

func solve(pieces []*Tetromino) (*Board, bool) {
	// handle empty input 
	if len(pieces) == 0 {
		return nil, false
    }
	// Calculate minimum board size
	minSize := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))

	for size := minSize; size <= len(pieces)*4; size++ {
		board := newBoard(size)
		if backtrack(board, pieces, 0) {
			return board, true
		}
	}

	return nil, false
}

func backtrack(board *Board, pieces []*Tetromino, pieceIndex int) bool {
	if pieceIndex >= len(pieces) {
		return true
	}

	piece := pieces[pieceIndex]
	for y := 0; y < board.size; y++ {
		for x := 0; x < board.size; x++ {
			if board.canPlace(piece, x, y) {
				board.place(piece, x, y)
				if backtrack(board, pieces, pieceIndex+1) {
					return true
				}
				board.remove(piece, x, y)
			}
		}
	}

	return false
}
