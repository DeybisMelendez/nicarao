package evaluation

import "nicarao/board"

func tempo(whiteToMove uint8) int {
	if whiteToMove == board.White {
		return tempoWeight
	}
	return -tempoWeight
}
