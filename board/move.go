package board

import (
	"math/bits"
)

type Move struct {
	Piece     Piece  //Piece es la pieza que se moverá
	From      Square //From es la ubicación de la pieza actualmente
	To        Square //To es la ubicación a la que se dirige la pieza
	Capture   Piece  //Capture es la pieza que se debería capturar ubicada en To
	Promotion Piece  //Promotion es la pieza que se está coronando, en caso de ser un peón
	Castling  uint8  // El único propósito de esta variable es poder recuperar el dato en UnMakeMove
}

func (s *Board) MakeMove(move *Move) {
	var color bool = s.WhiteToMove
	var piece uint64 = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
	move.Castling = s.Castling
	if move.Capture != None {
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
	}
	if move.Promotion != None {
		piece = PopBit(piece, move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
	}
	if move.Piece == King {
		s.HandleCastle(color, CastleShort, true)
		s.HandleCastle(color, CastleLong, true)
		var dist int = int(move.To) - int(move.From)
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
		if s.CanCastle(color, CastleLong) {
			if move.From == A1 || move.From == A8 {
				s.HandleCastle(color, CastleLong, true)
			}
		}
		if s.CanCastle(color, CastleShort) {
			if move.From == H1 || move.From == H8 {
				s.HandleCastle(color, CastleShort, true)
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
	var color bool = !s.WhiteToMove
	s.Castling = move.Castling
	var piece uint64 = PopBit(SetBit(s.Bitboards[color][move.Piece], move.From), move.To)

	if move.Capture != None {
		s.Bitboards[!color][move.Capture] = SetBit(s.Bitboards[!color][move.Capture], move.To)
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
	var color bool = s.WhiteToMove
	for _, piece := range pieceTypes {
		var pieceBoard uint64 = s.Bitboards[color][piece]

		for pieceBoard != 0 {
			var from Square = Square(bits.TrailingZeros64(pieceBoard))
			var attacks uint64 = s.GenerateAttacksForPiece(piece, from, s.WhiteToMove)
			if piece == King {
				var kingsideOK bool = s.occupied&castlingMask[color][CastleShort] == 0 &&
					s.CanCastle(color, CastleShort) &&
					s.AnyUnderAttack(s.WhiteToMove, castlingSquares[color][CastleShort]...)
				var queensideOK bool = s.occupied&castlingMask[color][CastleLong] == 0 &&
					s.CanCastle(color, CastleLong) &&
					s.AnyUnderAttack(s.WhiteToMove, castlingSquares[color][CastleLong]...)
				if kingsideOK {
					attacks = SetBit(attacks, castlingSquares[color][CastleShort][1])
				}
				if queensideOK {
					attacks = SetBit(attacks, castlingSquares[color][CastleLong][0])
				}
			}
			for attacks != 0 {
				var to Square = Square(bits.TrailingZeros64(attacks))
				var capture Piece = s.GetPiece(to, !s.WhiteToMove)

				if piece == Pawn && (to < A2 || to > H7) {
					for _, promotion := range piecePromotions {
						moves = append(moves, Move{From: from, To: to, Piece: piece,
							Capture: capture, Promotion: promotion})
					}
				} else {
					moves = append(moves, Move{From: from, To: to,
						Piece: piece, Capture: capture})
				}
				attacks &= attacks - 1
			}
			pieceBoard &= pieceBoard - 1
		}
	}
	return moves
}

func (s *Board) IsMoveLegal(move *Move) bool {
	var color bool = s.WhiteToMove
	s.MakeMove(move)
	var kingBB uint64 = s.Bitboards[color][King]
	var result bool
	if !s.IsUnderAttack(kingBB, !color) {
		result = true
	}
	s.UnMakeMove(move)
	return result
}
