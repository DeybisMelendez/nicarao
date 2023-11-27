package search

import (
	"nicarao/board"
	"testing"
)

func TestMatedIn(t *testing.T) {

	type testCase struct {
		mateIn int16
		fen    string
		id     string
	}

	testCases := []testCase{
		{1, "kn6/pp6/8/3N4/8/8/8/K7 w - - 0 1", "1.-Mate en 1"},
		{1, "nn6/kpP5/pp6/8/8/8/8/7K w - - 0 1", "2.-Mate en 1"},
		{1, "3N4/4BPP1/5Kpp/1BPp1Q1b/1p1kq2R/2b1r2n/1p1ppr2/8 w - - 0 1", "3.-Mate en 1"},

		{2, "5Kbk/6pp/6P1/8/8/8/8/7R w - - 0 1", "1.- Mate en 2"},
		{2, "b1B3Q1/5K2/5NP1/n7/2p2k1P/3pN2R/1B1P4/4qn2 w - - 0 1", "2.- Mate en 2"},
		{2, "6Q1/1Nn5/2p1rp2/2p5/2r1k2P/2PN3K/4PP2/1B2R3 w - - 0 1", "3.- Mate en 2"},
		{2, "1n3NR1/2B1R3/2pK1p2/2N2p2/5kbp/1r1p4/3ppr2/4b1QB w - - 0 1", "4.- Mate en 2"},
		{2, "8/8/5b2/8/8/5pk1/8/2BQK2R w K - 0 1", "5.- Mate en 2"},
		{2, "7k/1n6/7K/2p3Q1/1pr2p2/1qb5/8/1b6 w - - 0 1", "6.- Mate en 2"},
		{2, "3K4/B3p3/7p/2N2b1Q/p1rkPP1n/r2p3R/N4P2/3n4 w - - 0 1", "7.- Mate en 2"},
		{2, "5br1/pKn1P3/2Q5/3n4/2P1kP1N/1Npp1R2/r5B1/qb6 w - - 0 1", "8.- Mate en 2"},
		{2, "8/7B/2K1p1Q1/3pN3/3k1p2/1P6/3P1q2/3n4 w - - 0 1", "9.- Mate en 2"},
		{2, "8/4Kn2/1p1P2p1/Q1B2kp1/1P1r4/3P1N1P/4N1n1/1B2b3 w - - 0 1", "9.- Mate en 2"},

		{3, "1k6/1P5Q/8/7B/8/5K2/8/8 w - - 0 1", "1.- Mate en 3"},
		{3, "8/3Q4/3p4/2b2n2/3N1p2/nR2B2b/p7/k1K5 w - - 0 1", "2.- Mate en 3"},
		{3, "4K3/4P3/P1N2pP1/1BPk1P2/4N3/pP1p1Q2/1br2P2/q1r5 w - - 0 1", "3.- Mate en 3"},
		{3, "8/4p3/2pNP3/Kp6/8/k1p5/p1P5/B7 w - - 0 1", "4.- Mate en 3"},
		//{3, "1n2R3/2N5/3bBp2/1Kp1p3/1p1Pkpp1/3N4/3PPQ1n/3r4 w - - 0 1", "5.- Mate en 3"}, //FIXME: No encuentra todos mates

		{4, "8/1P5B/8/2P5/8/6K1/NkP3p1/RN6 w - - 0 1", "1.- Mate en 4"},
		{4, "2NK3k/2Np4/1p6/1b1P4/p1p2pRp/2r2p2/5P2/4n1n1 w - - 0 1", "2.- Mate en 4"}, // Tarda 1 segundo!

		//{4, "rnbq2k1/ppp1b1p1/6P1/3p4/3P1P2/6r1/PP4P1/R2QKB1R w KQ - 0 16", "2.- Mate en 4"},

		//{6, "2r1b3/3P4/3n1prp/2p2Nbk/q7/4N3/2p5/3n3K w - - 0 1", "1.- Mate en 7"},

	}

	for _, tc := range testCases {
		var b board.Board
		var score int16
		var mate int16
		b.ParseFEN(tc.fen)
		for i := uint8(1); i < 18; i++ {
			score = PVSearch(&b, -Inf, Inf, i)
			if 1+(MateValue-score)/2 == tc.mateIn {
				mate = tc.mateIn
				break
			}
		}
		if mate != tc.mateIn {
			t.Errorf("%s, Solución: Mate en %d, Resultado: Mate en %d", tc.id, tc.mateIn, score)
		}
	}
}
