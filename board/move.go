package board

import "math/bits"

type Move struct {
	Piece       Piece
	From        Square
	To          Square
	Capture     Piece
	Promotion   Piece
	CastleRight uint8
}

func (s *Board) MakeMove(move *Move) {
	color := s.WhiteToMove
	piece := SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
	move.CastleRight = s.CastlingRights
	if move.Capture != None {
		capture := s.Bitboards[!color][move.Capture]
		s.Bitboards[!color][move.Capture] = PopBit(capture, move.To)
	}
	if move.Promotion != None {
		promotion := s.Bitboards[color][move.Promotion]
		piece = PopBit(piece, move.To)
		s.Bitboards[color][move.Promotion] = SetBit(promotion, move.To)
	}
	if move.Piece == King {
		s.HandleCastle(color, true, true)
		s.HandleCastle(color, false, true)
		dist := int(move.To) - int(move.From)
		if dist == 2 { // Enroque corto
			if color {
				s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H1), F1)
			} else {
				s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H8), F8)
			}
		} else if dist == -2 { // Enroque largo
			if color {
				s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A1), D1)
			} else {
				s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A8), D8)
			}
		}
	}

	if move.Piece == Rook {
		if move.From == A1 || move.From == A8 {
			if s.CanCastleLong(color) {
				s.HandleCastle(color, false, true)
			}
		}
		if move.From == H1 || move.From == H8 {
			if s.CanCastleShort(color) {
				s.HandleCastle(color, true, true)
			}
		}
	}

	s.Bitboards[color][move.Piece] = piece
	s.WhiteToMove = !s.WhiteToMove
}

func (s *Board) UnMakeMove(move *Move) {
	color := !s.WhiteToMove
	s.CastlingRights = move.CastleRight
	piece := PopBit(SetBit(s.Bitboards[color][move.Piece], move.From), move.To)

	if move.Capture != None {
		capture := s.Bitboards[!color][move.Capture]
		s.Bitboards[!color][move.Capture] = SetBit(capture, move.To)
	}
	if move.Promotion != None {
		s.Bitboards[color][move.Promotion] |= PopBit(0, move.To)
	}

	if move.Piece == King {
		dist := int(move.To) - int(move.From)
		if dist == 2 { // Enroque corto
			if color {
				s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H1), F1)
			} else {
				s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H8), F8)
			}
		} else if dist == -2 { // Enroque largo
			if color {
				s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A1), D1)
			} else {
				s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A8), D8)
			}
		}
	}

	s.Bitboards[color][move.Piece] = piece
	s.WhiteToMove = color
}

func (board *Board) GenerateMoves() []Move {
	//TODO: Implementar captura al paso
	var moves []Move
	color := board.WhiteToMove
	friends := board.GetAll(color)
	enemies := board.GetAll(!color)
	occupied := board.GetOccupied()
	kingBB := board.Bitboards[color][King]

	for _, piece := range pieceTypes {
		pieceBoard := board.Bitboards[color][piece]

		for pieceBoard != 0 {
			from := Square(bits.TrailingZeros64(pieceBoard))
			attacks := board.GenerateAttacksForPiece(piece, from, occupied, friends, enemies)
			if piece == Pawn {
				attacks |= GetPawnPushes(color, from, occupied)
			} else if piece == King {
				shortMask := CastlingBlackShortMask
				longMask := CastlingBlackLongMask
				shortRights := board.CanCastleShort(color)
				longRights := board.CanCastleLong(color)
				shortSquares := squaresBlackShortCastling
				longSquares := squaresBlackLongCastling
				if color {
					shortMask = CastlingWhiteShortMask
					longMask = CastlingWhiteLongMask
					shortSquares = squaresWhiteShortCastling
					longSquares = squaresWhiteLongCastling
				}
				kingsideOK := occupied&shortMask == 0 && shortRights && board.AnyUnderAttack(enemies, occupied, friends, shortSquares...)
				queensideOK := occupied&longMask == 0 && longRights && board.AnyUnderAttack(enemies, occupied, friends, longSquares...)
				if kingsideOK {
					attacks |= SetBit(0, shortSquares[1])
					/*move := Move{From: from, To: shortSquares[1], Piece: piece}
					board.MakeMove(move)
					if !board.IsUnderAttack(kingBB, enemies, occupied, friends) {
						// TODO: Agregar una evaluación del movimiento
						moves = append(moves, move)
					}
					board.UnMakeMove(move)*/
				}
				if queensideOK {
					attacks |= SetBit(0, longSquares[0])
					/*move := Move{From: from, To: longSquares[0], Piece: piece}
					board.MakeMove(move)
					if !board.IsUnderAttack(kingBB, enemies, occupied, friends) {
						// TODO: Agregar una evaluación del movimiento
						moves = append(moves, move)
					}
					board.UnMakeMove(move)*/
				}
			}
			for attacks != 0 {
				to := Square(bits.TrailingZeros64(attacks))
				capture := board.GetPiece(to, !board.WhiteToMove)

				if piece == Pawn && (to < A2 || to > H7) {
					for _, promotion := range piecePromotions {
						move := Move{From: from, To: to, Piece: piece, Capture: capture, Promotion: promotion}
						board.MakeMove(&move)
						board.WhiteToMove = color
						if !board.IsUnderAttack(kingBB, enemies, occupied, friends) {
							// TODO: Agregar una evaluación del movimiento
							moves = append(moves, move)
						}
						board.WhiteToMove = !color
						board.UnMakeMove(&move)

					}
				} else {
					move := Move{From: from, To: to, Piece: piece, Capture: capture}
					board.MakeMove(&move)
					board.WhiteToMove = color
					if !board.IsUnderAttack(kingBB, enemies, occupied, friends) {
						// TODO: Agregar una evaluación del movimiento
						moves = append(moves, move)
					}
					board.WhiteToMove = !color
					board.UnMakeMove(&move)
				}
				attacks &= attacks - 1
			}
			pieceBoard &= pieceBoard - 1
		}
	}
	return moves
}
