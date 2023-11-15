package board

import (
	"testing"
)

func TestPerft(t *testing.T) {
	// Aquí deberías inicializar un tablero con una posición específica para probar
	// y luego verificar el número esperado de nodos generados a ciertas profundidades.
	// Por ejemplo:
	var board Board = *NewBoard()
	board.ParseFEN(StartingPos)
	// Asegúrate de reemplazar "InitializeBoardWithPosition" con tu función de inicialización real.

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 20},     // Nodos en profundidad 1
		{2, 400},    // Nodos en profundidad 2
		{3, 8902},   // Nodos en profundidad 3
		{4, 197281}, // Nodos en profundidad 4
		// Agrega más casos de prueba según sea necesario...
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Para profundidad %d, se esperaban %d nodos, pero se encontraron %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}
