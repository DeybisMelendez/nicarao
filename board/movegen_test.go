package board

import "testing"

func TestGeneratePseudoCaptures(t *testing.T) {
	var b Board = *NewBoard()
	b.ParseFEN("7k/8/8/1B6/pRp5/3p4/pNP5/K2b1Q2 w - - 0 1")

	type testCase struct {
		square   Square
		captures uint8
	}

	testCases := []testCase{
		{A4, 3},
		{C4, 3},
		{D3, 3},
		{D1, 2},
		{A2, 1},
		{H8, 0},
	}
	for _, tc := range testCases {
		var moves MoveList
		b.GeneratePseudoCaptures(&moves, tc.square)
		if moves.Index != tc.captures {
			t.Errorf("Casilla: %d, Capturas: %d, Resultado: %d", tc.square, tc.captures, moves.Index)
		}
	}
}
