package evaluation

import (
	"math/bits"
	"nicarao/board"
)

//https://www.chessprogramming.org/King_Safety
func virtualMobility(b *board.Board) (int, int) {
	var whiteKing board.Square = board.Square(bits.TrailingZeros64(b.Bitboards[board.White][board.King]))
	var blackKing board.Square = board.Square(bits.TrailingZeros64(b.Bitboards[board.Black][board.King]))
	var mg int = int(b.GetBishopAttacks(whiteKing, board.White)|b.GetRookAttacks(whiteKing, board.White)-
		b.GetBishopAttacks(blackKing, board.Black)|b.GetRookAttacks(blackKing, board.Black)) *
		kingSafetyWeight[middleGame]
	var eg int = int(b.GetBishopAttacks(whiteKing, board.White)|b.GetRookAttacks(whiteKing, board.White)-
		b.GetBishopAttacks(blackKing, board.Black)|b.GetRookAttacks(blackKing, board.Black)) *
		kingSafetyWeight[endGame]
	return mg, eg
}
