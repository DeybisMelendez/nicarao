package board

func (s *Board) FlipTurn() {
	if s.WhiteToMove == White {
		s.WhiteToMove = Black
	} else {
		s.WhiteToMove = White
	}
}

func (s *Board) GetEnemyColor() uint8 {
	if s.WhiteToMove == White {
		return Black
	} else {
		return White
	}
}

//GetAll devuelve un bitboard con todas las casillas ocupadas del jugador indicado
func (s *Board) GetAll(color uint8) uint64 {
	return s.Bitboards[color][Pawn] | s.Bitboards[color][Knight] | s.Bitboards[color][Bishop] | s.Bitboards[color][Rook] | s.Bitboards[color][Queen] | s.Bitboards[color][King]
}

//GetPiece devuelve la pieza que está controlando la casilla indicada
func (s *Board) GetPiece(square Square, color uint8) Piece {
	var mask uint64 = 1 << square
	for _, piece := range pieceTypes {
		if s.Bitboards[color][piece]&mask != 0 {
			return piece
		}
	}
	return None
}

func (s *Board) IsCapture(square Square) bool {
	return s.occupied&(1<<square) != 0
}

//CanCastle devuelve true si el jugador indicado tiene derecho a enrocar corto
func (s *Board) CanCastle(color uint8, isShort uint8) bool {
	return castling[color][isShort]&s.Castling != 0
}

//UpdateCastling cambia el derecho a enroque al nuevo estado
func (s *Board) UpdateCastling(Castling uint8) {
	s.Castling |= Castling
}

//RemoveCastling elimina un derecho de enroque según se indique
func (s *Board) RemoveCastling(Castling uint8) {
	s.Castling &= ^Castling
}

//HandleCastle controla la modificación del derecho a enroque
func (s *Board) HandleCastle(color uint8, isShort uint8, remove bool) {
	if remove {
		s.RemoveCastling(castling[color][isShort])
	} else {
		s.UpdateCastling(castling[color][isShort])
	}
}
