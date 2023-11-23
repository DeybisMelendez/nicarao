package board

import "math/bits"

func (s *Board) GeneratePseudoMoves(moves *MoveList) {
	var color bool = s.WhiteToMove
	var from Square
	var pieceBoard uint64
	for _, piece := range pieceTypes {
		pieceBoard = s.Bitboards[color][piece]

		for pieceBoard != 0 {
			from = Square(bits.TrailingZeros64(pieceBoard))
			s.GeneratePseudoMovesForPiece(piece, from, color, moves)
			pieceBoard &= pieceBoard - 1
		}
	}
}

func (s *Board) GeneratePseudoMovesForPiece(piece Piece, from Square, color bool, moves *MoveList) {
	var attacks uint64
	var to Square
	var capture Piece
	var flag MoveFlag
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
					moves.Push(NewMove(piece, from, to, capture, promo, flag))
				}
			} else if (color && to-from == 16) || (!color && from-to == 16) { //Doble peón
				moves.Push(NewMove(piece, from, to, capture, 0, DoublePawnPush))
			} else if s.Enpassant != 0 && s.Enpassant == to { //Enpassant Capture
				moves.Push(NewMove(piece, from, to, capture, 0, EnpassantCapture))
			} else {
				if capture == None {
					flag = QuietMoves
				} else {
					flag = Capture
				}
				moves.Push(NewMove(piece, from, to, capture, 0, flag))
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
			moves.Push(NewMove(piece, from, to, capture, 0, flag))
			attacks &= attacks - 1
		}
	case King:
		attacks = s.GetKingAttacks(from, color)
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = s.GetPiece(to, !s.WhiteToMove)
			flag = QuietMoves
			var dist int8 = int8(to) - int8(from)
			if dist == 2 { // Enroque corto
				flag = KingCastle
			} else if dist == -2 { // Enroque largo
				flag = QueenCastle
			} else if capture != None { // Si hay captura no hay enroque
				flag = Capture
			}
			moves.Push(NewMove(piece, from, to, capture, 0, flag))
			attacks &= attacks - 1
		}
	}
}
