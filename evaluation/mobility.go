package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func getMobilityEval(b *board.Board, piece board.Piece, phase uint8) int {
	var mobility = (getMobilityPiece(b, piece, board.White) -
		getMobilityPiece(b, piece, board.Black)) *
		mobilityWeights[phase][piece]
	return mobility
}

func getMobilityPiece(b *board.Board, piece board.Piece, color uint8) int {
	var pieceBoard = b.Bitboards[color][piece]
	var from board.Square
	var attacks uint64
	for pieceBoard != 0 {
		from = board.Square(bits.TrailingZeros64(pieceBoard))
		switch piece {
		case board.Knight:
			attacks |= b.GetKnightAttacks(from, color)
		case board.Bishop:
			attacks |= b.GetBishopAttacks(from, color)
		case board.Rook:
			attacks |= b.GetRookAttacks(from, color)
			/*case board.Queen:
			attacks |= b.GetBishopAttacks(from, color) | b.GetRookAttacks(from, color)*/
		}
		pieceBoard &= pieceBoard - 1
	}
	return bits.OnesCount64(attacks)
}
