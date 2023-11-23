package board

type UnMakeInfo struct {
	Enpassant Square
	Castling  uint8
}

func (s *Board) MakeMove(move Move) UnMakeInfo {
	var color bool = s.WhiteToMove
	var unMakeInfo UnMakeInfo = UnMakeInfo{
		Enpassant: s.Enpassant,
		Castling:  s.Castling,
	}
	var piece Piece = move.Piece()
	var capture Piece = move.Capture()
	var promo = move.Promotion()
	var to = move.To()
	var from = move.From()
	s.Enpassant = 0
	switch move.Flag() {
	case QuietMoves:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.friends = SetPopBit(s.friends, to, from)
	case DoublePawnPush:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.friends = SetPopBit(s.friends, to, from)
		if color {
			s.Enpassant = to - 8
		} else {
			s.Enpassant = to + 8
		}
	case Capture:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.Bitboards[!color][capture] = PopBit(s.Bitboards[!color][capture], to)
		s.friends = SetPopBit(s.friends, to, from)
		s.enemies = PopBit(s.enemies, to)
	case Promotion:
		s.Bitboards[color][piece] = PopBit(s.Bitboards[color][piece], from)
		s.Bitboards[color][promo] = SetBit(s.Bitboards[color][promo], to)
		s.friends = PopBit(s.friends, from)
		s.friends = SetBit(s.friends, to)
	case CapturePromotion:
		s.Bitboards[color][piece] = PopBit(s.Bitboards[color][piece], from)
		s.Bitboards[color][promo] = SetBit(s.Bitboards[color][promo], to)
		s.Bitboards[!color][capture] = PopBit(s.Bitboards[!color][capture], to)
		s.friends = SetPopBit(s.friends, to, from)
		s.enemies = PopBit(s.enemies, to)
	case KingCastle:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.friends = SetPopBit(s.friends, to, from)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H1), F1)
			s.friends = SetBit(PopBit(s.friends, H1), F1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H8), F8)
			s.friends = SetBit(PopBit(s.friends, H8), F8)
		}
	case QueenCastle:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.friends = SetPopBit(s.friends, to, from)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A1), D1)
			s.friends = SetBit(PopBit(s.friends, A1), D1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A8), D8)
			s.friends = SetBit(PopBit(s.friends, A8), D8)
		}
	case EnpassantCapture:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], to, from)
		s.friends = SetPopBit(s.friends, to, from)
		if color {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], to-8)
			s.enemies = PopBit(s.enemies, to-8)
		} else {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], to+8)
			s.enemies = PopBit(s.enemies, to+8)
		}
	}
	// Casos muy especiales sobre derechos de enroque
	if capture == Rook { // En caso de que una torre sea capturada en su casilla inicial
		if (to == A1 && !s.WhiteToMove) || (to == A8 && s.WhiteToMove) {
			s.HandleCastle(!color, CastleLong, true)
		} else if (to == H1 && !s.WhiteToMove) || (to == H8 && s.WhiteToMove) {
			s.HandleCastle(!color, CastleShort, true)
		}
	}
	if piece == Rook { // En caso de que tenga derecho a enrocar y este en su casilla inicial
		if (from == A1 && s.WhiteToMove) || (from == A8 && !s.WhiteToMove) {
			s.HandleCastle(color, CastleLong, true)
		}
		if (from == H1 && s.WhiteToMove) || (from == H8 && !s.WhiteToMove) {
			s.HandleCastle(color, CastleShort, true)
		}
	} else if piece == King {
		s.HandleCastle(color, CastleShort, true)
		s.HandleCastle(color, CastleLong, true)
	}
	s.WhiteToMove = !color
	var copyFriends uint64 = s.friends
	s.friends = s.enemies
	s.enemies = copyFriends
	s.occupied = s.friends | s.enemies
	return unMakeInfo
}

func (s *Board) UnMakeMove(move Move, info *UnMakeInfo) {
	var color bool = !s.WhiteToMove
	var copyFriends uint64 = s.friends
	var piece = move.Piece()
	var capture = move.Capture()
	var promo = move.Promotion()
	var to = move.To()
	var from = move.From()
	s.friends = s.enemies
	s.enemies = copyFriends
	s.Enpassant = info.Enpassant
	s.Castling = info.Castling

	switch move.Flag() {
	case QuietMoves, DoublePawnPush:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], from, to)
		s.friends = SetPopBit(s.friends, from, to)
	case Capture:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], from, to)
		s.Bitboards[!color][capture] = SetBit(s.Bitboards[!color][capture], to)
		s.friends = SetPopBit(s.friends, from, to)
		s.enemies = SetBit(s.enemies, to)
	case Promotion:
		s.Bitboards[color][piece] = SetBit(s.Bitboards[color][piece], from)
		s.Bitboards[color][promo] = PopBit(s.Bitboards[color][promo], to)
		s.friends = SetBit(s.friends, from)
		s.friends = PopBit(s.friends, to)
	case CapturePromotion:
		s.Bitboards[color][piece] = SetBit(s.Bitboards[color][piece], from)
		s.Bitboards[color][promo] = PopBit(s.Bitboards[color][promo], to)
		s.Bitboards[!color][capture] = SetBit(s.Bitboards[!color][capture], to)
		s.friends = SetPopBit(s.friends, from, to)
		s.enemies = SetBit(s.enemies, to)
	case KingCastle:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], from, to)
		s.friends = SetPopBit(s.friends, from, to)
		if color {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H1), F1)
			s.friends = PopBit(SetBit(s.friends, H1), F1)
		} else {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H8), F8)
			s.friends = PopBit(SetBit(s.friends, H8), F8)
		}
	case QueenCastle:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], from, to)
		s.friends = SetPopBit(s.friends, from, to)
		if color {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A1), D1)
			s.friends = PopBit(SetBit(s.friends, A1), D1)
		} else {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A8), D8)
			s.friends = PopBit(SetBit(s.friends, A8), D8)
		}
	case EnpassantCapture:
		s.Bitboards[color][piece] = SetPopBit(s.Bitboards[color][piece], from, to)
		s.friends = SetPopBit(s.friends, from, to)
		if color {
			s.Bitboards[!color][Pawn] = SetBit(s.Bitboards[!color][Pawn], to-8)
			s.enemies = SetBit(s.enemies, to-8)
		} else {
			s.Bitboards[!color][Pawn] = SetBit(s.Bitboards[!color][Pawn], to+8)
			s.enemies = SetBit(s.enemies, to+8)
		}
	}
	s.WhiteToMove = color
	s.occupied = s.friends | s.enemies
}
