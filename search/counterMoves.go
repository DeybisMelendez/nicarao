package search

import "nicarao/board"

//https://www.chessprogramming.org/Countermove_Heuristic
var counterMoves [2][7][64]board.Move

func saveCounterMove(color uint8, move board.Move) {
	move.SetScore(0) //Debemos eliminar la puntuación para comparar correctamente
	counterMoves[color][move.Piece()][move.To()] = move
}

func isCounterMove(color uint8, move board.Move) bool {
	return counterMoves[color][move.Piece()][move.To()] == move
}
