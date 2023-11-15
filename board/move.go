package board

type Move struct {
	Piece     Piece
	From      Square
	To        Square
	Capture   Piece
	Promotion Piece
	IsWhite   bool
}
