package evaluation

const (
	middleGame  uint8 = 0
	endGame     uint8 = 1
	pawnPhase         = 0
	knightPhase       = 1
	bishopPhase       = 1
	rookPhase         = 2
	queenPhase        = 4
	totalPhase        = knightPhase*4 + bishopPhase*4 + rookPhase*4 + queenPhase*2 // + pawnPhase*16
)
