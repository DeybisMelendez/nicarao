package evaluation

//FIXME: Arreglar los weights de la evaluacion
//https://www.chessprogramming.org/Point_Value
var materialWeights [2][7]int = [2][7]int{
	openingPhase: {0, 100, 400, 400, 600, 1200, 0},
	endPhase:     {0, 100, 350, 350, 525, 1000, 0},
}

var mobilityWeights [2][7]int = [2][7]int{
	openingPhase: {0, 0, 10, 10, 10, 0, 0},
	endPhase:     {0, 0, 10, 10, 10, 0, 0},
}
var pawnDoubledWeight int = 50
var pawnIsolatedWeight int = 50
var kingSafetyWeight int = 10

var tempoWeight int = 28
