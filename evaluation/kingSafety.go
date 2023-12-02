package evaluation

import (
	"math/bits"
	"nicarao/board"
)

//https://www.chessprogramming.org/King_Safety
func virtualMobility(b *board.Board, whiteKing board.Square, blackKing board.Square) int {
	return (bits.OnesCount64(b.GetBishopAttacks(whiteKing, board.White)|b.GetRookAttacks(whiteKing, board.White)) -
		bits.OnesCount64(b.GetBishopAttacks(blackKing, board.Black)|b.GetRookAttacks(blackKing, board.Black))) *
		kingSafetyWeight // * kingSafetyWeight[middleGame]
	/*var eg int = (bits.OnesCount64(b.GetBishopAttacks(whiteKing, board.White)|b.GetRookAttacks(whiteKing, board.White)) -
	bits.OnesCount64(b.GetBishopAttacks(blackKing, board.Black)|b.GetRookAttacks(blackKing, board.Black))) *
	kingSafetyWeight[endGame]* kingSafetyWeight[middleGame]*/
	//return mg, 0
}

func getKingTropismPiece(b *board.Board, pieceType board.Piece, whiteKing board.Square, blackKing board.Square) int {
	var pieceBoard = b.Bitboards[board.White][pieceType]
	var value int16
	var square board.Square
	for pieceBoard != 0 {
		square = board.Square(bits.TrailingZeros64(pieceBoard))
		value += kingTropism[pieceType][square][blackKing]
		pieceBoard &= pieceBoard - 1
	}
	pieceBoard = b.Bitboards[board.Black][pieceType]
	for pieceBoard != 0 {
		square = board.Square(bits.TrailingZeros64(pieceBoard))
		value -= kingTropism[pieceType][square][whiteKing]
		pieceBoard &= pieceBoard - 1
	}
	return int(value)
}
