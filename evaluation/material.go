package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func materialEval(b *board.Board) (int, int) {
	var totalMG int
	var totalEG int
	for i := board.Pawn; i < board.King; i++ {
		totalMG += (bits.OnesCount64(b.Bitboards[board.White][i]) -
			bits.OnesCount64(b.Bitboards[board.Black][i])) *
			materialWeights[middleGame][i]
		totalEG += (bits.OnesCount64(b.Bitboards[board.White][i]) -
			bits.OnesCount64(b.Bitboards[board.Black][i])) *
			materialWeights[endGame][i]
	}
	return (totalMG), (totalEG)
}
