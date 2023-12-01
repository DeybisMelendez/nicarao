package evaluation

//FIXME: Arreglar los weights de la evaluacion
var materialWeights [2][7]int = [2][7]int{
	middleGame: {0, 100, 320, 320, 500, 1000, 0},
	endGame:    {0, 110, 320, 340, 500, 900, 0},
}
var mobilityWeights [2][7]int = [2][7]int{
	middleGame: {0, 0, 1, 1, 1, 0, 0},
	endGame:    {0, 0, 1, 1, 1, 0, 0},
}
var kingSafetyWeight [2]int = [2]int{5, 0}

var tempoWeight int = 28
