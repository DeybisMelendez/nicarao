package board

import (
	"testing"
)

func TestStartingPos(t *testing.T) {
	//Init()
	var board Board = *NewBoard()
	board.ParseFEN(StartingPos)

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 20},      // Nodos en profundidad 1
		{2, 400},     // Nodos en profundidad 2
		{3, 8902},    // Nodos en profundidad 3
		{4, 197281},  // Nodos en profundidad 4
		{5, 4865609}, // Nodos en profundidad 5 14 segundos
		//{6, 119060324}, // Nodos en profundidad 6 Nota: Demasiado lento para llegar al nodo 6
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestKiwipete(t *testing.T) {
	//Init()
	var board Board = *NewBoard()
	board.ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - ")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 48},   // Nodos en profundidad 1
		{2, 2039}, // Nodos en profundidad 2
		//{3, 97862},   // Nodos en profundidad 3
		//{4, 4085603}, // Nodos en profundidad 4
		//{5, 193690690}, // Nodos en profundidad 5
		//{6, 8031647685}, // Nodos en profundidad 6
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestMoves(t *testing.T) {
	//Init()
	type testCase struct {
		nodesCount int64
		solution   int64
		fen        string
		test       string
	}

	testCases := []testCase{
		{1, 4, "7k/P7/8/8/8/6pp/6np/6nK w - - 0 1", "Promoción 1 Peón"},
		{1, 32, "8/PPPPPPPP/8/8/8/8/1rr5/K1k5 w - - 0 1", "Coronación de 8 peones Blancas"},
		{1, 32, "5K1k/5RR1/8/8/8/8/pppppppp/8 b - - 0 1", "Coronación de 8 peones Negras"},
		{1, 0, "8/8/3bb3/2b5/4K3/8/4b3/k7 w - - 0 1", "Ahogado solo Alfiles"},
		{1, 0, "8/8/3r1r2/4K3/3r1r2/8/8/k7 w - - 0 1", "Ahogado solo Torres"},
		{1, 0, "8/8/3q4/8/4K3/8/5q2/k7 w - - 0 1", "Ahogado solo Damas"},
		{1, 0, "8/5n2/8/n2K2n1/8/1n6/8/k7 w - - 0 1", "Ahogado solo caballos"},
		{1, 0, "k7/8/8/8/8/3ppp2/2p3p1/4K3 w - - 0 1", "Ahogado solo peones"},
		{1, 0, "8/8/8/8/8/k1b5/1r6/K7 w - - 0 1", "Ahogado Alfil y Torre"},
		{1, 0, "8/8/8/8/8/4k3/4p3/4K3 w - - 0 1", "Ahogado Peón y Rey"},
		{1, 4, "k5r1/8/8/3p4/2pRp3/3p4/6r1/7K w - - 0 1", "Capturas de Torre"},
		{1, 4, "k5r1/8/8/2p1p3/3B4/2p1p3/6r1/7K w - - 0 1", "Capturas de Alfil"},
		{1, 8, "k7/8/2p1p3/1p3p2/3N4/1p3p2/2p1p1r1/7K w - - 0 1", "Capturas de Caballo"},
		{1, 2, "k5r1/8/8/2ppp3/3P4/8/6r1/7K w - - 0 1", "Capturas de Peón bloqueado"},
		{1, 21, "k6K/rpbpnprp/P1P1P1P1/nppprpbn/1P1P1P1P/rqppqpbp/P1P1P1P1/8 w - - 0 1", "Muchas capturas Peones"},
		{1, 7, "k3r3/5Q2/8/8/8/8/8/4K3 w - - 0 1", "Jaque"},
		{1, 3, "k3r3/8/8/7Q/7b/8/8/4K3 w - - 0 1", "Jaque doble 1"},
		{1, 26, "4k3/8/8/8/4b3/8/8/R3K2R w KQ - 0 1", "Enroques"},
		{1, 24, "4k3/6r1/8/8/8/8/1b6/R3K2R w KQ - 0 1", "Sin poder Enrocar"},
		{1, 22, "4k3/8/8/8/8/8/8/RN2K1nR w KQ - 0 1", "Sin poder Enrocar 2"},
		{1, 4, "4k3/8/8/8/8/2b5/3N4/4K3 w - - 0 1", "Pieza clavada"},
		{1, 6, "4k3/8/8/b7/1P6/8/8/4K3 w - - 0 1", "Peón clavado con captura"},
		{1, 15, "k3r3/1b5b/2B1P1P1/8/rqR1K1Br/8/2R1P1R1/1q2r2q w - - 0 1", "Muchas clavadas"},
		{1, 2, "k7/8/8/4pP2/8/8/1rr5/K7 w - e6 0 2", "Captura al paso Blancas"},
		{1, 2, "k7/1R6/1R6/8/4Pp2/8/8/K7 b - e3 0 1", "Captura al paso Negras"},
		{1, 4, "k7/8/8/4PpP1/8/1r6/1r6/K7 w - f6 0 2", "2 Capturas al paso"},
	}

	for _, tc := range testCases {
		var board Board = *NewBoard()
		board.ParseFEN(tc.fen)
		var count int64 = Perft(&board, int(tc.nodesCount))
		if tc.solution != count {
			t.Errorf("Prueba de Movimientos: %s, solución: %d nodos, resultado: %d nodos", tc.test, tc.solution, count)
		}
	}
}
