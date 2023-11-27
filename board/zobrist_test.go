package board

import "testing"

func TestZobristHash(t *testing.T) {

	type testCase struct {
		depth int
		fen   string
	}

	testCases := []testCase{
		{6, "3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"},
		{6, "8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1"},
		{6, "8/8/1k6/2b5/2pP4/8/5K2/8 b - d3 0 1"},
		{6, "5k2/8/8/8/8/8/8/4K2R w K - 0 1"},
		{6, "3k4/8/8/8/8/8/8/R3K3 w Q - 0 1"},
		{4, "r3k2r/1b4bq/8/8/8/8/7B/R3K2R w KQkq - 0 1"},
		{4, "r3k2r/8/3Q4/8/8/5q2/8/R3K2R b KQkq - 0 1"},
		{6, "2K2r2/4P3/8/8/8/8/8/3k4 w - - 0 1"},
		{5, "8/8/1P2K3/8/2n5/1q6/8/5k2 b - - 0 1"},
		{6, "4k3/1P6/8/8/8/8/K7/8 w - - 0 1"},
		{6, "8/P1k5/K7/8/8/8/8/8 w - - 0 1"},
		{6, "K1k5/8/P7/8/8/8/8/8 w - - 0 1"},
		{7, "8/k1P5/8/1K6/8/8/8/8 w - - 0 1"},
		{4, "8/8/2k5/5q2/5n2/8/5K2/8 b - - 0 1"},
	}

	for _, tc := range testCases {
		var board Board = Board{}
		board.ParseFEN(tc.fen)
		var originalHash = board.Hash
		Perft(&board, tc.depth)
		if originalHash != board.Hash {
			t.Errorf("Hash no válido: Hash Original: %d, Hash del tablero: %d", originalHash, board.Hash)
		}
	}
}
