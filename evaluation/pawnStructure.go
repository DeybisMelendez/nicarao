package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func getStructurePawnEval(b *board.Board) int {
	var whiteVal int
	var blackVal int
	var doubled int
	for file := 0; file < 8; file++ {
		//Restando el peón de la fila, el resultado son los peones doblados
		doubled = bits.OnesCount64(b.Bitboards[board.White][board.Pawn]&onlyFile[file]) - 1
		if doubled > 0 {
			whiteVal -= doubled * pawnDoubledWeight
		}
		doubled = bits.OnesCount64(b.Bitboards[board.Black][board.Pawn]&onlyFile[file]) - 1
		if doubled > 0 {
			blackVal -= doubled * pawnDoubledWeight
		}
	}
	return whiteVal - blackVal
}
