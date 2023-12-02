package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func Evaluate(b *board.Board) int16 {
	// score inicia con el valor del turno a jugar 1 para blancas, -1 para negras
	var score int
	var turn int16 = -1
	var openingEval int
	var endGameEval int
	var whiteKing board.Square = board.Square(bits.TrailingZeros64(b.Bitboards[board.White][board.King]))
	var blackKing board.Square = board.Square(bits.TrailingZeros64(b.Bitboards[board.Black][board.King]))
	if b.WhiteToMove == board.White {
		turn = 1
	}
	for piece := board.Pawn; piece <= board.King; piece++ {
		//Material Evaluation
		if isMaterialEvalActive {
			openingEval += getMaterialEval(b, piece, openingPhase)
			endGameEval += getMaterialEval(b, piece, endPhase)
		}
		//Mobility Evaluation
		if isMobilityEvalActive && piece != board.Pawn && piece != board.King && piece != board.Queen {
			openingEval += getMobilityEval(b, piece, openingPhase)
			endGameEval += getMobilityEval(b, piece, endPhase)
		}
		//Piece Square Table Evaluation
		if isPSTEvalActive {
			openingEval += getPST(b, piece, openingPhase)
			endGameEval += getPST(b, piece, endPhase)
		}
		// King Tropism
		if isKingTropismActive {
			openingEval += getKingTropismPiece(b, piece, whiteKing, blackKing)
		}
	}
	//Structure Pawn Evaluation
	if isStructurePawnActive {
		var pawns = getStructurePawnEval(b)
		openingEval += pawns
		endGameEval += pawns
	}
	//King Evaluation
	if isVirtualMobilityActive {
		openingEval += virtualMobility(b, whiteKing, blackKing)
	}
	//Tapered Evaluation
	phase := totalPhase - (bits.OnesCount64(b.Bitboards[board.White][board.Knight]|b.Bitboards[board.Black][board.Knight]) * knightPhase) -
		(bits.OnesCount64(b.Bitboards[board.White][board.Bishop]|b.Bitboards[board.Black][board.Bishop]) * bishopPhase) -
		(bits.OnesCount64(b.Bitboards[board.White][board.Rook]|b.Bitboards[board.Black][board.Rook]) * rookPhase) -
		(bits.OnesCount64(b.Bitboards[board.White][board.Queen]|b.Bitboards[board.Black][board.Queen]) * queenPhase)
	if phase < 0 {
		phase = 0
	}
	phase = (phase*256 + (totalPhase / 2)) / totalPhase

	score = (openingEval*(256-phase) + endGameEval*phase) / 256
	// Tempo bonus
	score += tempo(b.WhiteToMove)
	// Rule 50
	score = score * (100 - int(b.HalfmoveClock)) / 100
	return int16(score) * turn
}
