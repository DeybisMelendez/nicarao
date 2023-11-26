package search

import "nicarao/board"

var historyMoves [2][7][64]uint8

func saveHistoryMove(color uint8, move board.Move, depth uint8) {
	// FIXME: Pensar que hacer si alcanza la puntuación maxima 255
	historyMoves[color][move.Piece()][move.To()] += depth * depth
}

func getHistoryMove(color uint8, move board.Move) uint8 {
	return historyMoves[color][move.Piece()][move.To()]
}
