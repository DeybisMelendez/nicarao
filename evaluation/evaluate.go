package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func Evaluate(b *board.Board) int16 {
	// score inicia con el valor del turno a jugar 1 para blancas, -1 para negras
	var score int16 = -1 // el score es multiplicado por el total que se calcula al final
	if b.WhiteToMove == board.White {
		score = 1
	}
	score *= materialEval(b) + mobilityEval(b)
	return score
}

func materialEval(b *board.Board) int16 {
	var total int
	for i := 1; i < 7; i++ {
		total += (bits.OnesCount64(b.Bitboards[board.White][i]) -
			bits.OnesCount64(b.Bitboards[board.Black][i])) *
			MaterialOpeningWeights[i]
	}
	return int16(total)
}
