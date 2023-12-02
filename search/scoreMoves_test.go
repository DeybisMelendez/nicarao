package search

import (
	"nicarao/board"
	"testing"
)

/*func TestRecapturedValue(t *testing.T) {

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
		var value int8 = see(&b, moves.List[0])
		if value != tc.value {
			t.Errorf("FEN: %s, Valor: %d, Resultado: %d", tc.fen, tc.value, value)
		}
	}
}*/
func TestSee(t *testing.T) {
	type testCase struct {
		fen      string
		solution int8
		square   board.Square
	}

	testCases := []testCase{
		{"7k/5nb1/6n1/4p3/2N3N1/2Bn1N2/1P5B/K7 w - - 0 1", 1, board.E5},
		{"7k/8/5p2/4q3/3P4/8/8/K7 w - - 0 1", 8, board.E5},
		{"7k/8/3p4/4p3/3P1Q2/8/8/K7 w - - 0 1", 1, board.E5},
		{"7k/8/3p1p2/4p3/3P1Q2/8/8/K7 w - - 0 1", 0, board.E5},
		{"7k/8/3b4/4p3/3P1Q2/8/8/K7 w - - 0 1", 1, board.E5},
		{"7K/3B4/4r3/4n3/3B4/7k/8/8 w - - 0 1", 3, board.E5},
	}

	for _, tc := range testCases {
		b := &board.Board{}
		b.ParseFEN(tc.fen)
		square := board.Square(tc.square)
		result := see(b, square)
		if result != tc.solution {
			t.Errorf("see() en la posición %d: resultado incorrecto, se esperaba %d pero se obtuvo %d", square, tc.solution, result)
		}
	}
}
