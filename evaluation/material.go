package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func getMaterialEval(b *board.Board, piece board.Piece, phase uint8) int {
	return (bits.OnesCount64(b.Bitboards[board.White][piece]) -
		bits.OnesCount64(b.Bitboards[board.Black][piece])) *
		materialWeights[phase][piece]
}
