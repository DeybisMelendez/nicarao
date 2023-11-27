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
	var minIndexMoves uint8
	var scoreMove uint8
	var bestI uint8
	b.GeneratePseudoCaptures(&captures)
	scoreCaptures(&captures)
	for minIndexMoves < captures.Index {
		if orderingMoveIsActive {
			bestI = minIndexMoves
			for i := minIndexMoves; i < captures.Index; i++ { // Se itera a través de los pseudo moves
				var move board.Move = captures.List[i]
				var newScore uint8 = move.GetScore()
				if newScore > scoreMove {
					scoreMove = newScore
					bestI = i
				}
			}
			//Intercambiamos el mejor pseudo move y lo colocamos al inicio para no volverlo a revisar
			captures.List[bestI], captures.List[minIndexMoves] = captures.List[minIndexMoves], captures.List[bestI]
		}
		//Recuperamos el pseudo move mejor evaluado actual
		var move board.Move = captures.List[minIndexMoves]

		//Sumamos el indice para no evaluar movimientos ya evaluados
		minIndexMoves++

		b.MakeMove(move)
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if !b.IsUnderAttack(kingSquare, color) { // Si la captura es legal
			score = -Quiesce(b, -beta, -alpha)
			b.UnMakeMove(move)
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
			b.UnMakeMove(move) //Si es ilegal se deshace y pasa al siguiente
		}
	}
	for i := 0; i < int(captures.Index); i++ {

	}
	return alpha
}
