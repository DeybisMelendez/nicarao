package search

import (
	"math/bits"
	"nicarao/board"
	"nicarao/evaluation"
)

/*
int Quiesce( int alpha, int beta ) {
    int stand_pat = Evaluate();
    if( stand_pat >= beta )
        return beta;
    if( alpha < stand_pat )
        alpha = stand_pat;

    until( every_capture_has_been_examined )  {
        MakeCapture();
        score = -Quiesce( -beta, -alpha );
        TakeBackMove();

        if( score >= beta )
            return beta;
        if( score > alpha )
           alpha = score;
    }
    return alpha;
}
*/

func Quiesce(b *board.Board, alpha int16, beta int16) int16 {
	var standPat = evaluation.Evaluate(b)
	if standPat >= beta {
		return beta
	}
	if alpha < standPat {
		alpha = standPat
	}
	var captures board.MoveList
	var color uint8 = b.WhiteToMove
	var kingSquare board.Square
	var score int16
	b.GeneratePseudoCaptures(&captures)

	for i := 0; i < int(captures.Index); i++ {
		b.MakeMove(captures.List[i])
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if !b.IsUnderAttack(kingSquare, color) { // Si la captura es legal
			score = -Quiesce(b, -beta, -alpha)
			b.UnMakeMove(captures.List[i])
			if score >= beta {
				return beta
			}
			if deltaPruningisActive {
				if score < alpha-deltaPruning {
					return alpha
				}
			}
			if score > alpha {
				alpha = score
			}
		} else {
			b.UnMakeMove(captures.List[i]) //Si es ilegal se deshace y pasa al siguiente
		}

	}
	return alpha
}
