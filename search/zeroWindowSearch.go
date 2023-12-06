package search

import (
	"math/bits"
	"nicarao/board"
	"nicarao/evaluation"
)

/*
// fail-hard zero window search, returns either beta-1 or beta
int zwSearch(int beta, int depth ) {
   // alpha == beta - 1
   // this is either a cut- or all-node
   if( depth == 0 ) return quiesce(beta-1, beta);
   for ( all moves)  {
     make
     score = -zwSearch(1-beta, depth - 1);
     unmake
     if( score >= beta )
        return beta;   // fail-hard beta-cutoff
   }
   return beta-1; // fail-hard, return alpha
}
*/

//fail-hard zero window search, returns either beta-1 or beta
func zwSearch(b *board.Board, beta int16, depth uint8) int16 {
	if isTimeToStop(b.Nodes) {
		return 0
	}
	if depth == 0 {
		return evaluation.Evaluate(b) //Quiesce tarda demasiado?
	}
	var moves board.MoveList
	var kingSquare board.Square
	var color uint8 = b.WhiteToMove
	var score int16
	var move board.Move
	b.GeneratePseudoMoves(&moves)
	//Aplicar ordenamiento de movimientos aqui

	for i := 0; i < int(moves.Index); i++ {
		move = moves.List[i]
		b.MakeMove(move)
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if !b.IsUnderAttack(kingSquare, color) { //Si el movimiento es legal!
			score = -zwSearch(b, 1-beta, depth-1)
			b.UnMakeMove(move)
			if score >= beta {
				return beta // fail-hard beta-cutoff
			}
		} else {
			b.UnMakeMove(move)
		}
	}
	if isTimeToStop(b.Nodes) {
		return 0
	}
	return beta - 1 // fail-hard, return alpha
}
