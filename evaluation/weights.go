package evaluation

//FIXME: Arreglar los weights de la evaluacion
var materialWeights [2][7]int = [2][7]int{
	middleGame: {0, 100, 320, 320, 500, 1000, 10000},
	endGame:    {0, 100, 320, 320, 500, 1000, 10000},
}
var mobilityWeights [2][7]int = [2][7]int{
	middleGame: {0, 0, 8, 8, 5, 5, 0},
	endGame:    {0, 0, 8, 8, 5, 5, 0},
}
var kingSafetyWeight [2]int = [2]int{0, 0}

//var MaterialOpeningWeights [7]int = [7]int{0, 90, 320, 320, 500, 1000, 10000}
//var MaterialEndingWeights [7]int = [7]int{0, 120, 320, 350, 500, 900, 10000}

//Nota: La movilidad de los caballos y alfiles es mas valiosa en la apertura y mediojuego que las torres, damas y rey.
//var MobilityOpeningWeights [7]int = [7]int{0, 0, 15, 15, 10, 10, 5}
//var MobilityEndingWeights [7]int = [7]int{0, 5, 10, 10, 10, 10, 10}
