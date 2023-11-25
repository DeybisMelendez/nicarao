package search

import (
	"nicarao/board"
	"testing"
)

func TestTranspositionTableBestMove(t *testing.T) {

	type testCase struct {
		fen string
		id  string
	}

	testCases := []testCase{
		{"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -", "Position 2 (Perft)"},
		{"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -", "Position 3 (Perft)"},
		{"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", "Position 4 (Perft)"},
	}

	for _, tc := range testCases {
		var b board.Board
		b.ParseFEN(tc.fen)
		for i := uint8(1); i < 5; i++ {
			PVSearch(&b, -Inf, Inf, i)
			for _, ttEntry := range transpositionTable {
				if ttEntry.Hash != 0 && ttEntry.BestMove == 0 {
					t.Errorf("Se encontró una entrada válida sin movimiento almacenado, Hash: %d Score: %d Depth: %d", ttEntry.Hash, ttEntry.Value, ttEntry.Depth)
				}
			}
		}
	}
}
