package evaluation

var MaterialOpeningWeights [7]int = [7]int{0, 90, 320, 320, 500, 1000, 10000}
var MaterialEndingWeights [7]int

//Nota: La movilidad de los caballos y alfiles es mas valiosa en la apertura y mediojuego que las torres, damas y rey.
var MobilityOpeningWeights [7]int = [7]int{0, 0, 15, 15, 10, 10, 5}
