package board

type Board struct {
	WhiteToMove    bool
	Bitboards      map[bool]map[Piece]uint64 // Mapa para almacenar bitboards por color y tipo de pieza
	Enpassant      Square
	CastlingRights uint8
	Halfmoveclock  uint8
	friends        uint64
	enemies        uint64
	occupied       uint64
}

func NewBoard() *Board {
	return &Board{
		Bitboards: map[bool]map[Piece]uint64{
			White: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
			Black: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
		},
	}
}
func (s *Board) CloneBoard() *Board {
	clonedBoard := &Board{
		WhiteToMove:    s.WhiteToMove,
		Enpassant:      s.Enpassant,
		CastlingRights: s.CastlingRights,
		Halfmoveclock:  s.Halfmoveclock,
	}

	// Copiar bitboards
	clonedBoard.Bitboards = make(map[bool]map[Piece]uint64)
	for color, pieces := range s.Bitboards {
		clonedBoard.Bitboards[color] = make(map[Piece]uint64)
		for piece, bb := range pieces {
			clonedBoard.Bitboards[color][piece] = bb
		}
	}

	return clonedBoard
}
func (s *Board) GetOccupied() uint64 {
	return s.Bitboards[White][Pawn] | s.Bitboards[White][Bishop] | s.Bitboards[White][Knight] |
		s.Bitboards[White][Rook] | s.Bitboards[White][Queen] | s.Bitboards[White][King] |
		s.Bitboards[Black][Pawn] | s.Bitboards[Black][Bishop] | s.Bitboards[Black][Knight] |
		s.Bitboards[Black][Rook] | s.Bitboards[Black][Queen] | s.Bitboards[Black][King]
}
func (s *Board) GetAll(color bool) uint64 {
	//var pieceTypes = []Piece{Pawn, Bishop, Knight, Rook, Queen, King}
	var total uint64
	for _, piece := range pieceTypes {
		total |= s.Bitboards[color][piece]
	}
	return total
}

func (s *Board) GetPiece(square Square, color bool) Piece {
	var mask uint64 = SetBit(0, square)

	for _, piece := range pieceTypes {
		if s.Bitboards[color][piece]&mask != 0 {
			return piece
		}
	}
	return None
}

func (s *Board) CanCastleShort(isWhite bool) bool {
	castleRight := CastlingWhiteShort
	if !isWhite {
		castleRight = CastlingBlackShort
	}
	return s.CastlingRights&castleRight != 0
}

func (s *Board) CanCastleLong(isWhite bool) bool {
	castleRight := CastlingWhiteLong
	if !isWhite {
		castleRight = CastlingBlackLong
	}
	return s.CastlingRights&castleRight != 0
}

func (s *Board) UpdateCastlingRights(castleRight uint8) {
	s.CastlingRights |= castleRight
}

func (s *Board) RemoveCastlingRights(castleRight uint8) {
	s.CastlingRights &= ^castleRight
}

func (s *Board) HandleCastle(isWhite bool, isShort bool, remove bool) {
	var castle uint8
	if isWhite {
		if isShort {
			castle = CastlingWhiteShort
		} else {
			castle = CastlingWhiteLong
		}
	} else {
		if isShort {
			castle = CastlingBlackShort
		} else {
			castle = CastlingBlackLong
		}
	}
	if remove {
		s.RemoveCastlingRights(castle)
	} else {
		s.UpdateCastlingRights(castle)
	}
}
