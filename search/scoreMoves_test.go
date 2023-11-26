package search

import (
	"nicarao/board"
	"testing"
)

func TestRecapturedValue(t *testing.T) {

	type testCase struct {
		value  uint8
		square board.Square
		fen    string
	}

	testCases := []testCase{
		{0, board.D8, "3q3k/4P3/8/8/8/8/8/K7 w - - 0 1"},
		{9, board.E3, "7k/8/8/8/3p4/4pQ2/8/K7 w - - 0 1"},
		{5, board.C4, "7k/8/8/2r5/2r1R3/8/8/K7 w - - 0 1"},
		{1, board.C4, "6bk/8/8/8/2r5/3P4/8/K7 w - - 0 1"},
		{1, board.F6, "7k/6p1/8/4Pp2/8/8/8/K7 w - f6 0 1"},
	}

	for _, tc := range testCases {
		var moves board.MoveList
		var b board.Board
		b.ParseFEN(tc.fen)
		b.GeneratePseudoCaptureSquare(&moves, tc.square)
		var value uint8 = recapturedValue(&b, moves.List[0])
		if value != tc.value {
			t.Errorf("FEN: %s, Valor: %d, Resultado: %d", tc.fen, tc.value, value)
		}
	}
}
