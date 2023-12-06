package evaluation

import (
	"math/bits"
	"nicarao/board"
)

func getStructurePawnEval(b *board.Board) int {
	var whiteVal int
	var blackVal int
	var value int
	var whitePawns uint64 = b.Bitboards[board.White][board.Pawn]
	var blackPawns uint64 = b.Bitboards[board.Black][board.Pawn]
	var whitePawnFile int
	var blackPawnFile int
	for file := 0; file < 8; file++ {
		whitePawnFile = bits.OnesCount64(whitePawns & onlyFile[file])
		blackPawnFile = bits.OnesCount64(blackPawns & onlyFile[file])
		//Evaluando peones aislados
		if file > 0 && file < 7 {
			if whitePawnFile > 0 {
				if bits.OnesCount64(whitePawns&onlyFile[file-1]) == 0 &&
					bits.OnesCount64(whitePawns&onlyFile[file+1]) == 0 {
					whiteVal -= pawnIsolatedWeight
				}
			}
			if blackPawnFile > 0 {
				if bits.OnesCount64(blackPawns&onlyFile[file-1]) == 0 &&
					bits.OnesCount64(blackPawns&onlyFile[file+1]) == 0 {
					blackVal -= pawnIsolatedWeight
				}
			}
		} else if file == 0 {
			if whitePawnFile > 0 {
				if bits.OnesCount64(whitePawns&onlyFile[file+1]) == 0 {
					whiteVal -= pawnIsolatedWeight
				}
			}
			if blackPawnFile > 0 {
				if bits.OnesCount64(blackPawns&onlyFile[file+1]) == 0 {
					blackVal -= pawnIsolatedWeight
				}
			}
		} else if file == 7 {
			if whitePawnFile > 0 {
				if bits.OnesCount64(whitePawns&onlyFile[file-1]) == 0 {
					whiteVal -= pawnIsolatedWeight
				}
			}
			if blackPawnFile > 0 {
				if bits.OnesCount64(blackPawns&onlyFile[file-1]) == 0 {
					blackVal -= pawnIsolatedWeight
				}
			}
		}
		//Evaluando peones doblados
		//Restando el peón de la fila, el resultado son los peones doblados
		value = whitePawnFile - 1
		if value > 0 {
			whiteVal -= value * pawnDoubledWeight
		}
		value = blackPawnFile - 1
		if value > 0 {
			blackVal -= value * pawnDoubledWeight
		}
	}
	return whiteVal - blackVal
}
