package search

import "nicarao/board"

func LMRisOk(move board.Move) bool {
	return move.Capture() == board.None && move.Promotion() == board.None
}
