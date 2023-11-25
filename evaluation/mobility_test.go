package evaluation

import (
	"nicarao/board"
	"testing"
)

func TestMobilityValues(t *testing.T) {

	type testCase struct {
		mobility int16
		fen      string
		id       string
	}

	testCases := []testCase{
		{0, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "Posición inicial"},
		{135, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1", "1.e4"},
		{115, "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1", "1.d4"},
		{55, "rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R b KQkq - 1 1", "1.Cf3"},
	}

	for _, tc := range testCases {
		var b board.Board
		b.ParseFEN(tc.fen)
		var result int16 = mobilityEval(&b)
		if result != tc.mobility {
			t.Errorf("%s, Solución: %d cp, Resultado: %d cp", tc.id, tc.mobility, result)
		}
	}
}
