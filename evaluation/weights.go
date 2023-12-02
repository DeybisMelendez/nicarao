package evaluation

//FIXME: Arreglar los weights de la evaluacion
var materialWeights [2][7]int = [2][7]int{
	openingPhase: {0, 100, 320, 320, 500, 1000, 0},
	endPhase:     {0, 110, 320, 340, 500, 900, 0},
}

var mobilityWeights [2][7]int = [2][7]int{
	openingPhase: {0, 0, 15, 10, 8, 0, 0},
	endPhase:     {0, 0, 10, 10, 15, 0, 0},
}
var pawnDoubledWeight int = 10

var kingSafetyWeight int = 10

var tempoWeight int = 28
