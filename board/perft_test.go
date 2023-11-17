package board

import (
	"testing"
)

func TestStartingPos(t *testing.T) {
	var board Board = *NewBoard()
	board.ParseFEN(StartingPos)

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

func TestPromotions(t *testing.T) {
	type testCase struct {
		solution int64
		fen      string
		test     string
	}

	testCases := []testCase{
		{4, "7k/P7/8/8/8/6pp/6np/6nK w - - 0 1", "Promote 1 Pawn White"},
		{32, "8/PPPPPPPP/8/8/8/8/1rr5/K1k5 w - - 0 1", "Promote 8 Pawn White"},
		{32, "5K1k/5RR1/8/8/8/8/pppppppp/8 b - - 0 1", "Promote 8 Pawn Black"},
	}

	for _, tc := range testCases {
		var board Board = *NewBoard()
		board.ParseFEN(tc.fen)
		var count int64 = Perft(&board, 1)
		if tc.solution != count {
			t.Errorf("Prueba de coronación: %s, nodos esperados: %d, %d nodos encontrados", tc.test, tc.solution, count)
		}
	}
}
func TestDraw(t *testing.T) {

	type testCase struct {
		nodesCount int64
		fen        string
		test       string
	}

	testCases := []testCase{
		{1, "8/8/3bb3/2b5/4K3/8/4b3/k7 w - - 0 1", "Only Bishops"},
		{1, "8/8/3r1r2/4K3/3r1r2/8/8/k7 w - - 0 1", "Only Rooks"},
		{1, "8/8/3q4/8/4K3/8/5q2/k7 w - - 0 1", "Only Queens"},
		{1, "8/5n2/8/n2K2n1/8/1n6/8/k7 w - - 0 1", "Only Knights"},
		{1, "k7/8/8/8/8/3ppp2/2p3p1/4K3 w - - 0 1", "Only Pawns and King"},
		{1, "8/8/8/8/8/k1b5/1r6/K7 w - - 0 1", "Rook and Bishop"},
		{1, "8/8/8/8/8/4k3/4p3/4K3 w - - 0 1", "Only 1 Pawn"},
	}

	for _, tc := range testCases {
		var board Board = *NewBoard()
		board.ParseFEN(tc.fen)
		var count int64 = Perft(&board, int(tc.nodesCount))
		if count != 0 {
			t.Errorf("Prueba de tablas: %s, se encontraron %d nodos disponibles", tc.test, count)
		}
	}
}
