package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func Evaluate(b *board.Board) int16 {
	// score inicia con el valor del turno a jugar 1 para blancas, -1 para negras
	var score int
	var turn int16 = -1
	if b.WhiteToMove == board.White {
		turn = 1
	}
	materialMg, materialEg := materialEval(b)
	mobilityMg, mobilityEg := mobilityEval(b)
	pstMg, pstEg := getPST(b)
	//FIXME: VirtualMobility reduce la precisión del mate
	kingSafetyMg, kingSafetyEg := virtualMobility(b)

	var opening = materialMg + pstMg + mobilityMg - kingSafetyMg
	var endgame = materialEg + pstEg + mobilityEg - kingSafetyEg
	phase := (bits.OnesCount64(b.Bitboards[board.White][board.Knight]|b.Bitboards[board.Black][board.Knight]) * knightPhase) +
		(bits.OnesCount64(b.Bitboards[board.White][board.Bishop]|b.Bitboards[board.Black][board.Bishop]) * bishopPhase) +
		(bits.OnesCount64(b.Bitboards[board.White][board.Rook]|b.Bitboards[board.Black][board.Rook]) * rookPhase) +
		(bits.OnesCount64(b.Bitboards[board.White][board.Queen]|b.Bitboards[board.Black][board.Queen]) * queenPhase)
	phase = (phase*256 + (totalPhase / 2)) / totalPhase

	score += (int(opening)*(256-phase) + int(endgame)*phase) / 256
	score += tempo(b.GetEnemyColor())
	score *= (100 - int(b.HalfmoveClock)) / 100
	return int16(score) * turn
}
