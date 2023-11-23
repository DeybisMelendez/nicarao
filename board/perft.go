package board

import "math/bits"

func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}

	var totalNodes int64
	var moves MoveList = MoveList{}
	b.GeneratePseudoMoves(&moves)
	var color bool = b.WhiteToMove
	var kingSquare Square
	for i := uint8(0); i < moves.Index; i++ {
		b.MakeMove(moves.List[i])
		kingSquare = Square(bits.TrailingZeros64(b.Bitboards[color][King]))
		if !b.IsUnderAttack(kingSquare, color) {
			totalNodes += Perft(b, depth-1)
		}
		b.UnMakeMove(moves.List[i])
	}
	return totalNodes
}
