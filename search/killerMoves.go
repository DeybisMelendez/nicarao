package search

import "nicarao/board"

var killerMoves [2][killerMovesMaxPly]board.Move

func saveKillerMove(ply int16, move board.Move) {
	move.SetScore(0) //Debemos eliminar la puntuación para comparar correctamente
	if ply < killerMovesMaxPly && killerMoves[1][ply] != move &&
		killerMoves[0][ply] != move {

		killerMoves[1][ply] = killerMoves[0][ply]
		killerMoves[0][ply] = move
	}
}

func isKillerMove(ply int16, move board.Move) uint8 {
	if ply < killerMovesMaxPly {
		if killerMoves[0][ply] == move {
			return 2
		} else if killerMoves[1][ply] == move {
			return 1
		}
	}
	return 0
}
