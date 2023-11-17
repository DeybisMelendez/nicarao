package board

import (
	"math/bits"
)

type Move struct {
	Piece       Piece
	From        Square
	To          Square
	Capture     Piece
	Promotion   Piece
	CastleRight uint8 // El único propósito de esta variable es poder recuperar el dato en UnMakeMove
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
		piece = PopBit(piece, move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
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
	s.WhiteToMove = !color
	s.friends = s.GetAll(!color)
	s.enemies = s.GetAll(color)
	s.occupied = s.friends | s.enemies
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
		s.Bitboards[color][move.Promotion] = PopBit(s.Bitboards[color][move.Promotion], move.To)
		s.Bitboards[color][Pawn] = SetBit(s.Bitboards[!color][Pawn], move.From)
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
	s.friends = s.GetAll(color)
	s.enemies = s.GetAll(!color)
	s.occupied = s.friends | s.enemies
}

func (s *Board) GeneratePseudoMoves() []Move {
	//TODO: Implementar captura al paso
	var moves []Move
	color := s.WhiteToMove
	//s.friends = s.GetAll(color)
	//s.enemies = s.GetAll(!color)
	//s.occupied = s.friends | s.enemies
	for _, piece := range pieceTypes {
		pieceBoard := s.Bitboards[color][piece]

		for pieceBoard != 0 {
			from := Square(bits.TrailingZeros64(pieceBoard))
			attacks := s.GenerateAttacksForPiece(piece, from)
			/*if piece == Pawn {
				// No se añade en GenerateAttacksForPiece porque no es un ataque realmente
				attacks |= s.GetPawnPushes(from)
			} else */
			if piece == King {
				shortMask := CastlingBlackShortMask
				longMask := CastlingBlackLongMask
				shortRights := s.CanCastleShort(color)
				longRights := s.CanCastleLong(color)
				shortSquares := squaresBlackShortCastling
				longSquares := squaresBlackLongCastling
				if color {
					shortMask = CastlingWhiteShortMask
					longMask = CastlingWhiteLongMask
					shortSquares = squaresWhiteShortCastling
					longSquares = squaresWhiteLongCastling
				}
				kingsideOK := s.occupied&shortMask == 0 && shortRights && s.AnyUnderAttack(shortSquares...)
				queensideOK := s.occupied&longMask == 0 && longRights && s.AnyUnderAttack(longSquares...)
				if kingsideOK {
					attacks = SetBit(attacks, shortSquares[1])
				}
				if queensideOK {
					attacks = SetBit(attacks, longSquares[0])
				}
			}
			for attacks != 0 {
				to := Square(bits.TrailingZeros64(attacks))
				capture := s.GetPiece(to, !s.WhiteToMove)

				if piece == Pawn && (to < A2 || to > H7) {
					for _, promotion := range piecePromotions {
						move := Move{From: from, To: to, Piece: piece, Capture: capture, Promotion: promotion}
						moves = append(moves, move)
					}
				} else {
					move := Move{From: from, To: to, Piece: piece, Capture: capture}
					moves = append(moves, move)
				}
				attacks &= attacks - 1
			}
			pieceBoard &= pieceBoard - 1
		}
	}
	return moves
}

/*func (s *Board) IsMoveLegal(move *Move) bool {
	color := s.WhiteToMove
	s.MakeMove(move)
	kingBB := s.Bitboards[color][King]
	s.WhiteToMove = color
	s.friends = s.GetAll(!color)
	s.enemies = s.GetAll(color)
	var result bool
	if !s.IsUnderAttack(kingBB) {
		result = true
	}
	s.WhiteToMove = !color
	s.UnMakeMove(move)
	return result
}*/
func (s *Board) IsMoveLegal(move *Move) bool {

	color := s.WhiteToMove
	var temp Board = *s.CloneBoard()
	temp.MakeMove(move)
	kingBB := temp.Bitboards[color][King]
	temp.WhiteToMove = !temp.WhiteToMove
	temp.friends = temp.GetAll(!color)
	temp.enemies = temp.GetAll(color)

	// Verificar si la casilla a la que se mueve el rey está bajo ataque
	return !temp.IsUnderAttack(kingBB)
}
