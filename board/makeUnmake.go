package board

func (s *Board) MakeMove(move Move) {
	s.Ply++
	s.pushUnMakeInfo()
	var color uint8 = s.WhiteToMove
	var enemy uint8 = s.GetEnemyColor()
	var piece Piece = move.Piece()
	var capture Piece = move.Capture()
	var promo = move.Promotion()
	var to = move.To()
	var from = move.From()
	var toBB uint64 = 1 << to
	var fromBB uint64 = 1 << from
	var fromToBB uint64 = fromBB ^ toBB
	s.Enpassant = 0
	switch move.Flag() {
	case QuietMoves:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
	case DoublePawnPush:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Enpassant = to - 8
		} else {
			s.Enpassant = to + 8
		}
	case Capture:
		s.Bitboards[color][piece] ^= fromToBB
		s.Bitboards[enemy][capture] ^= toBB
		s.friends ^= fromToBB
		s.enemies ^= toBB
	case Promotion:
		s.Bitboards[color][piece] &= ^fromBB
		s.Bitboards[color][promo] |= toBB
		s.friends &= ^fromBB
		s.friends |= toBB
	case CapturePromotion:
		s.Bitboards[color][piece] &= ^fromBB
		s.Bitboards[color][promo] |= toBB
		s.Bitboards[enemy][capture] ^= toBB
		s.friends ^= fromToBB
		s.enemies ^= toBB
	case KingCastle:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[color][Rook] ^= (1 << F1) ^ (1 << H1)
			s.friends ^= (1 << F1) ^ (1 << H1)
		} else {
			s.Bitboards[color][Rook] ^= (1 << F8) ^ (1 << H8)
			s.friends ^= (1 << F8) ^ (1 << H8)
		}
	case QueenCastle:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[color][Rook] ^= (1 << D1) ^ (1 << A1)
			s.friends ^= (1 << D1) ^ (1 << A1)
		} else {
			s.Bitboards[color][Rook] ^= (1 << D8) ^ (1 << A8)
			s.friends ^= (1 << D8) ^ (1 << A8)
		}
	case EnpassantCapture:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[enemy][Pawn] &= ^(1 << (to - 8))
			s.enemies &= ^(1 << (to - 8))
		} else {
			s.Bitboards[enemy][Pawn] &= ^(1 << (to + 8))
			s.enemies &= ^(1 << (to + 8))
		}
	}
	// Casos muy especiales sobre derechos de enroque
	if capture == Rook { // En caso de que una torre sea capturada en su casilla inicial
		if (to == A1 && s.WhiteToMove == Black) || (to == A8 && s.WhiteToMove == White) {
			s.HandleCastle(enemy, CastleLong, true)
		} else if (to == H1 && s.WhiteToMove == Black) || (to == H8 && s.WhiteToMove == White) {
			s.HandleCastle(enemy, CastleShort, true)
		}
	}
	if piece == Rook { // En caso de que tenga derecho a enrocar y este en su casilla inicial
		if (from == A1 && s.WhiteToMove == White) || (from == A8 && s.WhiteToMove == Black) {
			s.HandleCastle(color, CastleLong, true)
		}
		if (from == H1 && s.WhiteToMove == White) || (from == H8 && s.WhiteToMove == Black) {
			s.HandleCastle(color, CastleShort, true)
		}
	} else if piece == King {
		s.HandleCastle(color, CastleShort, true)
		s.HandleCastle(color, CastleLong, true)
	}
	s.FlipTurn()
	var copyFriends uint64 = s.friends
	s.friends = s.enemies
	s.enemies = copyFriends
	s.occupied = s.friends | s.enemies
}

func (s *Board) UnMakeMove(move Move) {
	s.Ply--
	s.popUnMakeInfo()
	var color uint8 = s.GetEnemyColor()
	var enemy uint8 = s.WhiteToMove
	var copyFriends uint64 = s.friends
	var piece = move.Piece()
	var capture = move.Capture()
	var promo = move.Promotion()
	var to = move.To()
	var from = move.From()
	var toBB uint64 = 1 << to
	var fromBB uint64 = 1 << from
	var fromToBB uint64 = fromBB ^ toBB
	s.friends = s.enemies
	s.enemies = copyFriends

	switch move.Flag() {
	case QuietMoves, DoublePawnPush:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
	case Capture:
		s.Bitboards[color][piece] ^= fromToBB
		s.Bitboards[enemy][capture] |= toBB
		s.friends ^= fromToBB
		s.enemies |= toBB
	case Promotion:
		s.Bitboards[color][piece] |= fromBB
		s.Bitboards[color][promo] &= ^toBB
		s.friends |= fromBB
		s.friends &= ^toBB
	case CapturePromotion:
		s.Bitboards[color][piece] |= fromBB
		s.Bitboards[color][promo] &= ^toBB
		s.Bitboards[enemy][capture] |= toBB
		s.friends ^= fromToBB
		s.enemies |= toBB
	case KingCastle:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[color][Rook] ^= (1 << H1) ^ (1 << F1)
			s.friends ^= (1 << H1) ^ (1 << F1)
		} else {
			s.Bitboards[color][Rook] ^= (1 << H8) ^ (1 << F8)
			s.friends ^= (1 << H8) ^ (1 << F8)
		}
	case QueenCastle:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[color][Rook] ^= (1 << A1) ^ (1 << D1)
			s.friends ^= (1 << A1) ^ (1 << D1)
		} else {
			s.Bitboards[color][Rook] ^= (1 << A8) ^ (1 << D8)
			s.friends ^= (1 << A8) ^ (1 << D8)
		}
	case EnpassantCapture:
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[enemy][Pawn] |= (1 << (to - 8))
			s.enemies |= (1 << (to - 8))
		} else {
			s.Bitboards[enemy][Pawn] |= (1 << (to + 8))
			s.enemies |= (1 << (to + 8))
		}
	}
	s.FlipTurn()
	s.occupied = s.friends | s.enemies
}
