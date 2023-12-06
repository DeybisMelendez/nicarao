package uci

import (
	"nicarao/board"
	"nicarao/search"
	"strconv"
	"time"
)

// HandleGo maneja el comando 'go' de UCI
func HandleGo(tokens []string) {
	//var depth int
	var wtime int
	var btime int
	var winc int
	var binc int

	//var movetime int
	//var movestogo int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		/*case "depth":
		if i+1 < len(tokens) {
			// Parsear el valor de profundidad
			depth = parseDepth(tokens[i+1])
		}*/
		case "wtime":
			if i+1 < len(tokens) {
				// Parsear el tiempo restante para las blancas
				wtime = parseTime(tokens[i+1])
			}
		case "btime":
			if i+1 < len(tokens) {
				// Parsear el tiempo restante para las negras
				btime = parseTime(tokens[i+1])
			}
		case "winc":
			if i+1 < len(tokens) {
				// Parsear el tiempo adicional para las blancas
				winc = parseTime(tokens[i+1])
			}
		case "binc":
			if i+1 < len(tokens) {
				// Parsear el tiempo adicional para las negras
				binc = parseTime(tokens[i+1])
			}
			/*case "movetime":
				if i+1 < len(tokens) {
					// Parsear el tiempo por movimiento
					movetime = parseTime(tokens[i+1])
				}
			case "movestogo":
				if i+1 < len(tokens) {
					// Parsear la cantidad de movimientos restantes para el tiempo de control
					movestogo = parseMovesToGo(tokens[i+1])
				}*/
		}
	}
	if wtime != 0 || btime != 0 { // Time to Play
		var timeLeft int
		if boardUCI.WhiteToMove == board.White {
			timeLeft = (wtime / 60) + winc
		} else {
			timeLeft = (btime / 60) + binc
		}
		var start int64 = time.Now().UnixMilli()
		go search.SearchWithStopTime(&boardUCI, start+int64(timeLeft))
	}

	// Iniciar la búsqueda con la profundidad proporcionada
	// Llama a tu función de búsqueda aquí
	// searchResult := YourSearchFunction(board, depth)

	// Envía el mejor movimiento encontrado al UCI
	// Por ejemplo, si tienes una función que retorna el mejor movimiento, podrías enviarlo así:
	// fmt.Println("bestmove", searchResult.BestMove)

	// ... (resto del código)
}

func parseTime(timeStr string) int {
	timeValue, err := strconv.Atoi(timeStr)
	if err != nil {
		// Manejo de error, retorna un valor por defecto
		return 0
	}
	// En UCI el tiempo se especifica en milisegundos
	return timeValue
}

/*
func parseDepth(depthStr string) int {
	// Parsea el valor de profundidad y devuelve un entero
	// Maneja cualquier error de conversión aquí
	// Por ejemplo, strconv.Atoi() o similar
	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		// Manejar error de conversión
		return 0 // Retorna un valor por defecto en caso de error
	}
	return depth
}

func parseMovesToGo(movesStr string) int {
	movesValue, err := strconv.Atoi(movesStr)
	if err != nil {
		// Manejo de error, retorna un valor por defecto
		return 0
	}
	return movesValue
}
*/
