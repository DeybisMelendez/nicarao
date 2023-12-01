package search

import "nicarao/board"

func LMRisOk(move board.Move) bool {
	return move.Flag() != board.Capture && move.Flag() != board.CapturePromotion && move.Flag() != board.Promotion
}
