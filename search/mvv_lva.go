package search

import "nicarao/board"

var MVV_LVA = [7][7]uint8{
	{ //Nothing
		0, 5, 4, 3, 2, 1, 0,
	},
	{ //Pawn
		0, 15, 14, 13, 12, 11, 10,
	},
	{ //Knight
		0, 25, 24, 23, 22, 21, 20,
	},
	{ //Bishop
		0, 35, 34, 33, 32, 31, 30,
	},
	{ //Rook
		0, 45, 44, 43, 42, 41, 40,
	},
	{ //Queen
		0, 55, 54, 53, 52, 51, 50,
	},
	{ //King
		0, 65, 64, 63, 62, 61, 60,
	},
}

func getMVV_LVA(move board.Move) uint8 {
	return MVV_LVA[move.Capture()][move.Piece()] //[Victims][Agressors]
}
