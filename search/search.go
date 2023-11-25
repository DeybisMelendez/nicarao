package search

import (
	"math"
	"math/bits"
	"nicarao/board"
	"nicarao/evaluation"
)

//https://www.chessprogramming.org/Principal_Variation_Search

/*
int pvSearch( int alpha, int beta, int depth ) {
   if( depth == 0 ) return quiesce(alpha, beta);
   bool bSearchPv = true;
   for ( all moves)  {
      make
      if ( bSearchPv ) {
         score = -pvSearch(-beta, -alpha, depth - 1);
      } else {
         score = -zwSearch(-alpha, depth - 1);
         if ( score > alpha ) // in fail-soft ... && score < beta ) is common
            score = -pvSearch(-beta, -alpha, depth - 1); // re-search
      }
      unmake
      if( score >= beta )
         return beta;   // fail-hard beta-cutoff
      if( score > alpha ) {
         alpha = score; // alpha acts like max in MiniMax
         bSearchPv = false;   // *1)
      }
   }
   return alpha;
}

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
1) it is recommend to set bSearchPv outside the score > alpha condition.
*/
//Principal Variation Search with Zero Window Search
func PVSearch(b *board.Board, alpha int16, beta int16, depth uint8) int16 {

	//Transposition Table
	var bestMove board.Move
	var hashFlag uint8 = TTUpperBound
	if ttValue := probeHash(b.Hash, alpha, beta, depth, &bestMove, b.HalfmoveClock); ttValue != NoHashEntry {
		return ttValue
	}
	// Si nodo es terminal
	if depth == 0 {
		var eval int16 = evaluation.Evaluate(b)
		//recordHash(b.Hash,eval,depth,TTExact,¿¿¿¿bestmove????)
		return eval //Quiesce
	}
	var color uint8 = b.WhiteToMove
	var bSearchPv bool = true
	var moves board.MoveList
	var kingSquare board.Square
	var hasLegalMove bool
	var score int16
	var bestScore int16 = math.MinInt16
	b.GeneratePseudoMoves(&moves)
	//Aplicar ordenamiento de movimientos aqui

	for i := 0; i < int(moves.Index); i++ {
		var move board.Move = moves.List[i]
		b.MakeMove(move)
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if !b.IsUnderAttack(kingSquare, color) { //Si el movimiento es legal!
			hasLegalMove = true
			if bSearchPv {
				score = -PVSearch(b, -beta, -alpha, depth-1)
			} else {
				score = -zwSearch(b, -alpha, depth-1)
				if score > alpha { // in fail-soft ... && score < beta ) is common
					score = -PVSearch(b, -beta, -alpha, depth-1) // re-search
				}
			}
			b.UnMakeMove(move)
			if score >= beta {
				recordHash(b.Hash, beta, depth, TTLowerBound, move, b.HalfmoveClock)
				return beta // fail-hard beta-cutoff
			}
			if score > alpha {
				hashFlag = TTExact
				alpha = score     // alpha acts like max in MiniMax
				bSearchPv = false // Probar luego sacando esta sentencia de esta condición
			}
			//Se recolecta el mejor movimiento posible de la posición
			if score > bestScore { // UpperBound y Exact
				bestMove = move
				bestScore = score
			}

		} else {
			b.UnMakeMove(move)
		}
	}
	if !hasLegalMove {
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if b.IsUnderAttack(kingSquare, color) {
			return -MateValue + b.Ply //Jaque mate
		} else {
			return 0 //Tablas por ahogado
		}
	}
	recordHash(b.Hash, alpha, depth, hashFlag, bestMove, b.HalfmoveClock)
	return alpha
}

//fail-hard zero window search, returns either beta-1 or beta
func zwSearch(b *board.Board, beta int16, depth uint8) int16 {
	if depth == 0 {
		return evaluation.Evaluate(b) //Quiesce
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
	return beta - 1 // fail-hard, return alpha
}
