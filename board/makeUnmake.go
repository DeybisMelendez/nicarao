package board

func (s *Board) MakeMove(move Move) {
	var color uint8 = s.WhiteToMove
	var enemy uint8 = s.GetEnemyColor()
	var piece Piece = move.Piece()
	var capture Piece = move.Capture()
	var promo Piece = move.Promotion()
	var to Square = move.To()
	var from Square = move.From()
	var toBB uint64 = 1 << to
	var fromBB uint64 = 1 << from
	var fromToBB uint64 = fromBB ^ toBB

	s.Nodes++
	s.Ply++
	s.HalfmoveClock++
	s.pushUnMakeInfo()

	s.Hash ^= whiteToMoveZobrist //Cambiamos el turno de la posición
	s.Hash ^= pieceSquareZobrist[color][piece][from]
	s.Hash ^= pieceSquareZobrist[color][piece][to]
	s.Hash ^= uint64(s.Enpassant) // Eliminamos el enpassant actual

	s.Enpassant = 0

	if piece == Pawn {
		s.HalfmoveClock = 0
	}

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
		s.Hash ^= pieceSquareZobrist[color][capture][to]
		s.HalfmoveClock = 0
	case Promotion:
		s.Bitboards[color][piece] &= ^fromBB
		s.Bitboards[color][promo] |= toBB
		s.friends &= ^fromBB
		s.friends |= toBB
		s.Hash ^= pieceSquareZobrist[color][promo][to]
	case CapturePromotion:
		s.Bitboards[color][piece] &= ^fromBB
		s.Bitboards[color][promo] |= toBB
		s.Bitboards[enemy][capture] ^= toBB
		s.friends ^= fromToBB
		s.enemies ^= toBB
		s.Hash ^= pieceSquareZobrist[color][piece][to] // Eliminamos el peón en octava
		s.Hash ^= pieceSquareZobrist[color][promo][to]
	case KingCastle: // El derecho al enroque se maneja mas adelante
		s.Bitboards[color][piece] ^= fromToBB
		s.friends ^= fromToBB
		if color == White {
			s.Bitboards[color][Rook] ^= (1 << F1) ^ (1 << H1)
			s.friends ^= (1 << F1) ^ (1 << H1)
		} else {
			s.Bitboards[color][Rook] ^= (1 << F8) ^ (1 << H8)
			s.friends ^= (1 << F8) ^ (1 << H8)
		}
	case QueenCastle: // El derecho al enroque se maneja mas adelante
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
		s.HalfmoveClock = 0
	}
	// En caso de que aún existan derechos a enrocar
	if s.Castling != 0 {
		// Casos muy especiales sobre derechos de enroque
		if capture == Rook { // En caso de que una torre sea capturada en su casilla inicial
			if (to == A1 && color == Black) || (to == A8 && color == White) {
				if s.CanCastle(enemy, CastleLong) {
					s.HandleCastle(enemy, CastleLong, true)
					s.Hash ^= castleRightsZobrist[enemy][CastleLong]
				}
			} else if (to == H1 && color == Black) || (to == H8 && color == White) {
				if s.CanCastle(enemy, CastleShort) {
					s.HandleCastle(enemy, CastleShort, true)
					s.Hash ^= castleRightsZobrist[enemy][CastleShort]
				}
			}
		}
		if piece == Rook { // En caso de que tenga derecho a enrocar y se mueva de su casilla inicial
			if (from == A1 && color == White) || (from == A8 && color == Black) {
				if s.CanCastle(color, CastleLong) {
					s.HandleCastle(color, CastleLong, true)
					s.Hash ^= castleRightsZobrist[color][CastleLong]
				}
			}
			if (from == H1 && color == White) || (from == H8 && color == Black) {
				if s.CanCastle(color, CastleShort) {
					s.HandleCastle(color, CastleShort, true)
					s.Hash ^= castleRightsZobrist[color][CastleLong]
				}
			}
		} else if piece == King { //Si el rey se mueve pierde el enroque, incluso al enrocar
			if s.CanCastle(color, CastleShort) {
				s.HandleCastle(color, CastleShort, true)
				s.Hash ^= castleRightsZobrist[color][CastleShort]
			}
			if s.CanCastle(color, CastleLong) {
				s.HandleCastle(color, CastleLong, true)
				s.Hash ^= castleRightsZobrist[color][CastleLong]
			}
		}
	}
	s.Hash ^= uint64(s.Enpassant) //Colocamos el nuevo enpassant
	s.FlipTurn()
	s.friends, s.enemies = s.enemies, s.friends
	s.occupied = s.friends | s.enemies
}

func (s *Board) UnMakeMove(move Move) {
	s.Ply--
	s.popUnMakeInfo()
	var color uint8 = s.GetEnemyColor()
	var enemy uint8 = s.WhiteToMove
	var piece = move.Piece()
	var capture = move.Capture()
	var promo = move.Promotion()
	var to = move.To()
	var from = move.From()
	var toBB uint64 = 1 << to
	var fromBB uint64 = 1 << from
	var fromToBB uint64 = fromBB ^ toBB
	s.friends, s.enemies = s.enemies, s.friends

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
