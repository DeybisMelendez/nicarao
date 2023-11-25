package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func mobilityEval(b *board.Board) int16 {
	var total int
	for i := 2; i < 7; i++ {
		total += (getMobility(b, board.Piece(i), board.White) -
			getMobility(b, board.Piece(i), board.Black)) *
			MobilityOpeningWeights[i]
	}
	return int16(total)
}

func getMobility(b *board.Board, piece board.Piece, color uint8) int {
	var pieceBoard = b.Bitboards[color][piece]
	var from board.Square
	var attacks uint64
	for pieceBoard != 0 {
		from = board.Square(bits.TrailingZeros64(pieceBoard))
		switch piece {
		//Nota: Me parece a mi que el principio de actividad (movilidad) no aplica en los peones
		//En los peones se debe evaluar la estructura y el espacio.
		/*case board.Pawn:
		attacks |= b.GetPawnAttacks(from, color)*/
		case board.Knight:
			attacks |= b.GetKnightAttacks(from, color)
		case board.Bishop:
			attacks |= b.GetBishopAttacks(from, color)
		case board.Rook:
			attacks |= b.GetRookAttacks(from, color)
		case board.Queen:
			attacks |= b.GetBishopAttacks(from, color) | b.GetRookAttacks(from, color)
		case board.King:
			attacks |= b.GetKingAttacks(from, color)
		}
		pieceBoard &= pieceBoard - 1
	}
	return bits.OnesCount64(attacks)
}
