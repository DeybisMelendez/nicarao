package search

import (
	"math"
	"math/bits"
	"nicarao/board"
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
1) it is recommend to set bSearchPv outside the score > alpha condition.
*/
//Principal Variation Search with Zero Window Search
func PVSearch(b *board.Board, alpha int16, beta int16, depth uint8) int16 {
	if isTimeToStop(b.Nodes) {
		return 0
	}
	//Transposition Table
	var bestMove board.Move
	var hashFlag uint8 = TTUpperBound
	if isTranspositionTableActive {
		if ttValue := probeHash(b.Hash, alpha, beta, depth, &bestMove, b.HalfmoveClock); ttValue != NoHashEntry {
			return ttValue
		}
	}
	// Si nodo es terminal
	if depth == 0 {
		var eval int16 = Quiesce(b, alpha, beta) //evaluation.Evaluate(b)
		//recordHash(b.Hash,eval,depth,TTExact,¿¿¿¿bestmove????)
		return eval //Quiesce
	}
	if internalIDIsActive {
		if bestMove == 0 && depth > internalIDminDepth { // Si no hay hash move previo
			// Internal Iterative Deepening
			for i := uint8(1); i < depth/2; i++ {
				PVSearch(b, alpha, beta, i)
			}
			bestMove = GetBestMove(b.Hash)
		}
	}

	var color uint8 = b.WhiteToMove
	var bSearchPv bool = true
	var moves board.MoveList
	var kingSquare board.Square
	var hasLegalMove bool
	var score int16
	var bestScore int16 = math.MinInt16
	var minIndexMoves uint8
	var scoreMove uint8
	var newDepth uint8
	var movesNoFailsHigh = 0
	var isLegal bool
	var ourKing board.Square   //= board.Square(bits.TrailingZeros64(b.Bitboards[b.WhiteToMove][board.King]))
	var enemyKing board.Square //= board.Square(bits.TrailingZeros64(b.Bitboards[b.GetEnemyColor()][board.King]))
	var isInCheck bool         // = b.IsUnderAttack(ourKing, b.WhiteToMove)
	var givingCheck bool       //= b.IsUnderAttack(enemyKing, b.GetEnemyColor())
	var isExtended bool
	if (depth >= LMRFullDepthMoves && LMRisActive) || (depth > nullMovePruningR && isNullMovePruningActive) {
		ourKing = board.Square(bits.TrailingZeros64(b.Bitboards[b.WhiteToMove][board.King]))
		enemyKing = board.Square(bits.TrailingZeros64(b.Bitboards[b.GetEnemyColor()][board.King]))
		isInCheck = b.IsUnderAttack(ourKing, b.WhiteToMove)
		givingCheck = b.IsUnderAttack(enemyKing, b.GetEnemyColor())
	}
	b.GeneratePseudoMoves(&moves)
	if orderingMoveIsActive {
		scoreMoves(b, &moves, bestMove) // ¿¡Posiblemente sea mejor intentar puntuar solo jugadas legales!?
	}
	// Selection Sort de mayor a menor
	for minIndexMoves < moves.Index {
		if orderingMoveIsActive {
			var bestI uint8 = minIndexMoves
			for i := minIndexMoves; i < moves.Index; i++ { // Se itera a través de los pseudo moves
				var move board.Move = moves.List[i]
				var newScore uint8 = move.GetScore()
				if newScore > scoreMove {
					scoreMove = newScore
					bestI = i
				}
			}
			//Intercambiamos el mejor pseudo move y lo colocamos al inicio para no volverlo a revisar
			moves.List[bestI], moves.List[minIndexMoves] = moves.List[minIndexMoves], moves.List[bestI]
		}

		//Recuperamos el pseudo move mejor evaluado actual
		var move board.Move = moves.List[minIndexMoves]

		//Sumamos el indice para no evaluar movimientos ya evaluados
		minIndexMoves++
		newDepth = depth - 1
		isExtended = false
		if isInCheck || givingCheck {
			newDepth++ //Check Extension
			isExtended = true
		}
		if move.Capture() != board.None || move.Promotion() != board.None {
			// No reducir capturas
			isExtended = true
		}

		//Mull Nove Pruning
		if isNullMovePruningActive && depth > nullMovePruningR &&
			depth > nullMovePruningLimit && !isInCheck {
			var nullBoard *board.Board = b.MakeNull()
			score = -zwSearch(nullBoard, -beta+1, depth-1-nullMovePruningR)
			if score >= beta {
				return beta
			}
		}

		//Late Move Reduction
		if LMRisActive {
			if !bSearchPv && movesNoFailsHigh >= LMRFullDepthMoves && depth >= LMReductionLimit {
				if !isInCheck && !givingCheck && !isExtended { // Si no está en jaque
					newDepth = depth / 3
				}
			}
		}

		//PVSearch
		b.MakeMove(move)
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		isLegal = !b.IsUnderAttack(kingSquare, color)

		if isLegal { //Si el movimiento es legal!
			hasLegalMove = true
			if bSearchPv {
				score = -PVSearch(b, -beta, -alpha, newDepth)
			} else {
				score = -zwSearch(b, -alpha, newDepth)
				if score > alpha { // in fail-soft ... && score < beta ) is common
					score = -PVSearch(b, -beta, -alpha, newDepth) // re-search
				}
			}
			b.UnMakeMove(move)
			if score >= beta {
				if isTimeToStop(b.Nodes) {
					return 0
				}
				if move.Capture() == board.None {
					saveKillerMove(b.Ply, move)
					saveCounterMove(color, move)
					saveHistoryMove(color, move, newDepth)
				}
				if isTranspositionTableActive {
					recordHash(b.Hash, beta, depth, TTLowerBound, move, b.HalfmoveClock)
				}
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
			movesNoFailsHigh++
		} else { // Si el movimiento es ilegal deshacemos y continuamos con el siguiente
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
	if isTimeToStop(b.Nodes) {
		return 0
	}
	if isTranspositionTableActive {
		recordHash(b.Hash, alpha, depth, hashFlag, bestMove, b.HalfmoveClock)
	}
	return alpha
}
