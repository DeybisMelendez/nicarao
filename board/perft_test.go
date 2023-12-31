package board

import (
	"testing"
)

func TestStartingPos(t *testing.T) {
	var board Board = Board{}
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
		{5, 4865609}, // Nodos en profundidad 5
		//{6, 119060324}, // Nodos en profundidad 6
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestPosition2(t *testing.T) {
	var board Board = Board{}
	board.ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 48},      // Nodos en profundidad 1
		{2, 2039},    // Nodos en profundidad 2
		{3, 97862},   // Nodos en profundidad 3
		{4, 4085603}, // Nodos en profundidad 4
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

func TestPosition3(t *testing.T) {
	//Init()
	var board Board = Board{}
	board.ParseFEN("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 14},       // Nodos en profundidad 1
		{2, 191},      // Nodos en profundidad 2
		{3, 2812},     // Nodos en profundidad 3
		{4, 43238},    // Nodos en profundidad 4
		{5, 674624},   // Nodos en profundidad 5
		{6, 11030083}, // Nodos en profundidad 6
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestPosition4(t *testing.T) {
	var board Board = Board{}
	board.ParseFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 6},        // Nodos en profundidad 1
		{2, 264},      // Nodos en profundidad 2
		{3, 9467},     // Nodos en profundidad 3
		{4, 422333},   // Nodos en profundidad 4
		{5, 15833292}, // Nodos en profundidad 5
		//{6, 706045033}, // Nodos en profundidad 6
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestPosition5(t *testing.T) {
	var board Board = Board{}
	board.ParseFEN("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 44},      // Nodos en profundidad 1
		{2, 1486},    // Nodos en profundidad 2
		{3, 62379},   // Nodos en profundidad 3
		{4, 2103487}, // Nodos en profundidad 4
		//{5, 89941194}, // Nodos en profundidad 5
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}
func TestPosition6(t *testing.T) {
	var board Board = Board{}
	board.ParseFEN("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10 ")

	type testCase struct {
		depth      int
		nodesCount int64
	}

	testCases := []testCase{
		{1, 46},      // Nodos en profundidad 1
		{2, 2079},    // Nodos en profundidad 2
		{3, 89890},   // Nodos en profundidad 3
		{4, 3894594}, // Nodos en profundidad 4
		//{5, 164075551}, // Nodos en profundidad 5
	}

	for _, tc := range testCases {
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

//https://www.chessprogramming.net/perfect-perft/
func TestTalkChess(t *testing.T) {

	type testCase struct {
		depth      int
		nodesCount int64
		fen        string
	}

	testCases := []testCase{
		{6, 1134888, "3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"},
		{6, 1015133, "8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1"},
		{6, 1440467, "8/8/1k6/2b5/2pP4/8/5K2/8 b - d3 0 1"},
		{6, 661072, "5k2/8/8/8/8/8/8/4K2R w K - 0 1"},
		{6, 803711, "3k4/8/8/8/8/8/8/R3K3 w Q - 0 1"},
		{4, 1274206, "r3k2r/1b4bq/8/8/8/8/7B/R3K2R w KQkq - 0 1"},
		{4, 1720476, "r3k2r/8/3Q4/8/8/5q2/8/R3K2R b KQkq - 0 1"},
		{6, 3821001, "2K2r2/4P3/8/8/8/8/8/3k4 w - - 0 1"},
		{5, 1004658, "8/8/1P2K3/8/2n5/1q6/8/5k2 b - - 0 1"},
		{6, 217342, "4k3/1P6/8/8/8/8/K7/8 w - - 0 1"},
		{6, 92683, "8/P1k5/K7/8/8/8/8/8 w - - 0 1"},
		{6, 2217, "K1k5/8/P7/8/8/8/8/8 w - - 0 1"},
		{7, 567584, "8/k1P5/8/1K6/8/8/8/8 w - - 0 1"},
		{4, 23527, "8/8/2k5/5q2/5n2/8/5K2/8 b - - 0 1"},
	}

	for _, tc := range testCases {
		var board Board = Board{}
		board.ParseFEN(tc.fen)
		count := Perft(&board, tc.depth)
		if count != tc.nodesCount {
			t.Errorf("Profundidad: %d, Solución: %d nodos, Resultado: %d nodos", tc.depth, tc.nodesCount, count)
		}
	}
}

func TestMoves(t *testing.T) {
	type testCase struct {
		nodesCount int64
		solution   int64
		fen        string
		test       string
	}

	testCases := []testCase{
		{1, 32, "8/PPPPPPPP/8/8/8/8/1rr5/K1k5 w - - 0 1", "Coronación de 8 peones Blancas"},
		{1, 32, "5K1k/5RR1/8/8/8/8/pppppppp/8 b - - 0 1", "Coronación de 8 peones Negras"},
		{1, 0, "NB1BBRRR/BpBPPPPP/1P1PRBP1/4P1P1/6p1/1pp2pP1/5P2/K1k4N w - - 0 1", "Ahogado muchas piezas"},
		{1, 0, "n3nBbB/P1b1P1PN/r1P1PBp1/2PKP1P1/2RPPk2/1q3n1p/7P/8 w - - 0 1", "Ahogado muchas piezas 2"},
		{1, 21, "k6K/rpbpnprp/P1P1P1P1/nppprpbn/1P1P1P1P/rqppqpbp/P1P1P1P1/8 w - - 0 1", "Muchas capturas Peones"},
		{1, 8, "k7/8/2p1p3/1p3p2/3N4/1p3p2/2p1p1r1/7K w - - 0 1", "Capturas de Caballo"},
		{1, 2, "k5r1/8/8/2ppp3/3P4/8/6r1/7K w - - 0 1", "Capturas de Peón bloqueado"},
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
		var board Board = Board{}
		board.ParseFEN(tc.fen)
		var count int64 = Perft(&board, int(tc.nodesCount))
		if tc.solution != count {
			t.Errorf("Prueba de Movimientos: %s, solución: %d nodos, resultado: %d nodos", tc.test, tc.solution, count)
		}
	}
}
