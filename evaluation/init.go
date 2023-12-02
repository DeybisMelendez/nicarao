package evaluation

import (
	"math"
	"nicarao/board"
)

func init() {
	setDist()
}

func setDist() {
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			distanceBonus[i][j] = 14 - int(math.Abs(float64(col[i])-float64(col[j]))+
				math.Abs(float64(row[i])-float64(row[j])))
			kingTropism[board.Knight][i][j] = int16(distanceBonus[i][j])
			kingTropism[board.Bishop][i][j] += int16(bonusDiagDistance[int(math.Abs(float64(diagNE[i])-float64(diagNE[j])))])
			kingTropism[board.Bishop][i][j] += int16(bonusDiagDistance[int(math.Abs(float64(diagNW[i])-float64(diagNW[j])))])
			kingTropism[board.Rook][i][j] = int16(distanceBonus[i][j] / 2)
			kingTropism[board.Queen][i][j] = int16((distanceBonus[i][j] * 5) / 2)
		}
	}
}
