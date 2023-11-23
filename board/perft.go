package board

import "math/bits"

func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}

	var totalNodes int64
	var moves []Move
	b.GeneratePseudoMoves(&moves)
	var color bool = b.WhiteToMove
	var unMakeInfo UnMakeInfo
	var kingSquare Square
	for _, move := range moves {
		unMakeInfo = b.MakeMove(move)
		kingSquare = Square(bits.TrailingZeros64(b.Bitboards[color][King]))
		if !b.IsUnderAttack(kingSquare, color) {
			totalNodes += Perft(b, depth-1)
		}
		b.UnMakeMove(move, &unMakeInfo)
	}
	return totalNodes
}
