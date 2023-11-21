package board

import (
	"math/bits"
)

type MoveFlag uint8
type Move struct {
	Piece     Piece  //Piece es la pieza que se moverá
	From      Square //From es la ubicación de la pieza actualmente
	To        Square //To es la ubicación a la que se dirige la pieza
	Capture   Piece  //Capture es la pieza que se debería capturar ubicada en To
	Promotion Piece
	Flag      MoveFlag
}
type UnMakeInfo struct {
	Enpassant Square
	Castling  uint8
}

func (s *Board) MakeMove(move *Move) UnMakeInfo {
	var color bool = s.WhiteToMove
	var unMakeInfo UnMakeInfo = UnMakeInfo{
		Enpassant: s.Enpassant,
		Castling:  s.Castling,
	}
	s.Enpassant = 0
	switch move.Flag {
	case QuietMoves:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.friends = SetPopBit(s.friends, move.To, move.From)
	case DoublePawnPush:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.friends = SetPopBit(s.friends, move.To, move.From)
		if color {
			s.Enpassant = move.To - 8
		} else {
			s.Enpassant = move.To + 8
		}
	case Capture:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
		s.friends = SetPopBit(s.friends, move.To, move.From)
		s.enemies = PopBit(s.enemies, move.To)
	case Promotion:
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.From)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
		s.friends = PopBit(s.friends, move.From)
		s.friends = SetBit(s.friends, move.To)
	case CapturePromotion:
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.From)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
		s.friends = PopBit(s.friends, move.From)
		s.friends = SetBit(s.friends, move.To)
		s.enemies = PopBit(s.enemies, move.To)
	case KingCastle:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.friends = SetPopBit(s.friends, move.To, move.From)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H1), F1)
			s.friends = SetBit(PopBit(s.friends, H1), F1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H8), F8)
			s.friends = SetBit(PopBit(s.friends, H8), F8)
		}
	case QueenCastle:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.friends = SetPopBit(s.friends, move.To, move.From)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A1), D1)
			s.friends = SetBit(PopBit(s.friends, A1), D1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A8), D8)
			s.friends = SetBit(PopBit(s.friends, A8), D8)
		}

	case EnpassantCapture:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		s.friends = SetPopBit(s.friends, move.To, move.From)
		if color {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To-8)
			s.enemies = PopBit(s.enemies, move.To-8)
		} else {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To+8)
			s.enemies = PopBit(s.enemies, move.To+8)
		}
	}
	// Casos muy especiales sobre derechos de enroque
	if move.Capture == Rook { // En caso de que una torre sea capturada en su casilla inicial
		if (move.To == A1 && !s.WhiteToMove) || (move.To == A8 && s.WhiteToMove) {
			s.HandleCastle(!color, CastleLong, true)
		} else if (move.To == H1 && !s.WhiteToMove) || (move.To == H8 && s.WhiteToMove) {
			s.HandleCastle(!color, CastleShort, true)
		}
	}
	if move.Piece == Rook { // En caso de que tenga derecho a enrocar y este en su casilla inicial
		if (move.From == A1 && s.WhiteToMove) || (move.From == A8 && !s.WhiteToMove) {
			s.HandleCastle(color, CastleLong, true)
		}
		if (move.From == H1 && s.WhiteToMove) || (move.From == H8 && !s.WhiteToMove) {
			s.HandleCastle(color, CastleShort, true)
		}
	}
	if move.Piece == King {
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

func (s *Board) UnMakeMove(move *Move, info *UnMakeInfo) {
	var color bool = !s.WhiteToMove
	var copyFriends uint64 = s.friends
	s.friends = s.enemies
	s.enemies = copyFriends
	s.Enpassant = info.Enpassant
	s.Castling = info.Castling

	switch move.Flag {
	case QuietMoves, DoublePawnPush:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.From, move.To)
		s.friends = SetPopBit(s.friends, move.From, move.To)
	case Capture:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.From, move.To)
		s.Bitboards[!color][move.Capture] = SetBit(s.Bitboards[!color][move.Capture], move.To)
		s.friends = SetPopBit(s.friends, move.From, move.To)
		s.enemies = SetBit(s.enemies, move.To)
	case Promotion:
		s.Bitboards[color][move.Piece] = SetBit(s.Bitboards[color][move.Piece], move.From)
		s.Bitboards[color][move.Promotion] = PopBit(s.Bitboards[color][move.Promotion], move.To)
		s.friends = SetBit(s.friends, move.From)
		s.friends = PopBit(s.friends, move.To)
	case CapturePromotion:
		s.Bitboards[color][move.Piece] = SetBit(s.Bitboards[color][move.Piece], move.From)
		s.Bitboards[color][move.Promotion] = PopBit(s.Bitboards[color][move.Promotion], move.To)
		s.Bitboards[!color][move.Capture] = SetBit(s.Bitboards[!color][move.Capture], move.To)
		s.friends = SetBit(s.friends, move.From)
		s.friends = PopBit(s.friends, move.To)
		s.enemies = SetBit(s.enemies, move.To)
	case KingCastle:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.From, move.To)
		s.friends = SetPopBit(s.friends, move.From, move.To)
		if color {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H1), F1)
			s.friends = PopBit(SetBit(s.friends, H1), F1)
		} else {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], H8), F8)
			s.friends = PopBit(SetBit(s.friends, H8), F8)
		}
	case QueenCastle:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.From, move.To)
		s.friends = SetPopBit(s.friends, move.From, move.To)
		if color {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A1), D1)
			s.friends = PopBit(SetBit(s.friends, A1), D1)
		} else {
			s.Bitboards[color][Rook] = PopBit(SetBit(s.Bitboards[color][Rook], A8), D8)
			s.friends = PopBit(SetBit(s.friends, A8), D8)
		}
	case EnpassantCapture:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.From, move.To)
		s.friends = SetPopBit(s.friends, move.From, move.To)
		if color {
			s.Bitboards[!color][Pawn] = SetBit(s.Bitboards[!color][Pawn], move.To-8)
			s.enemies = SetBit(s.enemies, move.To-8)
		} else {
			s.Bitboards[!color][Pawn] = SetBit(s.Bitboards[!color][Pawn], move.To+8)
			s.enemies = SetBit(s.enemies, move.To+8)
		}
	}
	s.WhiteToMove = color
	s.occupied = s.friends | s.enemies
}

func (s *Board) GeneratePseudoMoves() []Move {
	var moves []Move
	var color bool = s.WhiteToMove
	var from Square
	for _, piece := range pieceTypes {
		var pieceBoard uint64 = s.Bitboards[color][piece]

		for pieceBoard != 0 {
			from = Square(bits.TrailingZeros64(pieceBoard))
			moves = append(moves, s.GeneratePseudoMovesForPiece(piece, from, color)...)
			pieceBoard &= pieceBoard - 1
		}
	}
	return moves
}
func (s *Board) GeneratePseudoMovesForPiece(piece Piece, from Square, color bool) []Move {
	var attacks uint64
	var to Square
	var capture Piece
	var flag MoveFlag
	var moves []Move
	switch piece {
	case Pawn:
		attacks = s.GetPawnAttacks(from, color) | s.GetPawnPushes(from, color)
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = s.GetPiece(to, !s.WhiteToMove)
			if (to < A2 && !color) || (to > H7 && color) { //Coronación de peón
				if capture == None {
					flag = Promotion
				} else {
					flag = CapturePromotion
				}
				for _, promo := range piecePromotions {
					moves = append(moves, Move{From: from, To: to, Piece: piece,
						Capture: capture, Promotion: promo, Flag: flag})
				}
			} else if (color && to-from == 16) || (!color && from-to == 16) { //Doble peón
				moves = append(moves, Move{From: from, To: to,
					Piece: piece, Flag: DoublePawnPush, Capture: capture})
			} else if s.Enpassant != 0 && s.Enpassant == to { //Enpassant Capture
				moves = append(moves, Move{From: from, To: to,
					Piece: piece, Capture: capture, Flag: EnpassantCapture})
			} else {
				if capture == None {
					flag = QuietMoves
				} else {
					flag = Capture
				}
				moves = append(moves, Move{From: from, To: to, //Pawn move
					Piece: piece, Flag: flag, Capture: capture}) // en teoría no necesita Capture
			}
			attacks &= attacks - 1
		}
	case Knight, Bishop, Rook, Queen:
		switch piece {
		case Knight:
			attacks = s.GetKnightAttacks(from, color)
		case Bishop:
			attacks = s.GetBishopAttacks(from, color)
		case Rook:
			attacks = s.GetRookAttacks(from, color)
		case Queen:
			attacks = s.GetBishopAttacks(from, color) | s.GetRookAttacks(from, color)
		}
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = s.GetPiece(to, !s.WhiteToMove)
			flag = QuietMoves
			if capture != None {
				flag = Capture
			}
			moves = append(moves, Move{From: from, To: to,
				Piece: piece, Flag: flag, Capture: capture})
			attacks &= attacks - 1
		}
	case King:
		attacks = s.GetKingAttacks(from, color)
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = s.GetPiece(to, !s.WhiteToMove)
			flag = QuietMoves
			var dist int = int(to) - int(from)
			if dist == 2 { // Enroque corto
				flag = KingCastle
			} else if dist == -2 { // Enroque largo
				flag = QueenCastle
			} else if capture != None { // Si hay captura no hay enroque
				flag = Capture
			}
			moves = append(moves, Move{From: from, To: to,
				Piece: piece, Flag: flag, Capture: capture})
			attacks &= attacks - 1
		}
	}
	return moves
}
