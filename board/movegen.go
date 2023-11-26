package board

import "math/bits"

func (s *Board) GeneratePseudoMoves(moves *MoveList) {
	s.saveUnMakeInfo()
	var color uint8 = s.WhiteToMove
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

func (s *Board) GeneratePseudoMovesForPiece(piece Piece, from Square, color uint8, moves *MoveList) {
	var attacks uint64
	var to Square
	var flag MoveFlag
	var capture Piece
	var enemy = s.GetEnemyColor()
	switch piece {
	case Pawn:
		attacks = s.GetPawnAttacks(from, color) // | s.GetPawnPushes(from, color)
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = s.GetPiece(to, enemy)
			if ((1<<to)&Rank1 != 0) || ((1<<to)&Rank8 != 0) { //Coronación de peón
				for _, promo := range piecePromotions {
					moves.Push(NewMove(piece, from, to, capture, promo, CapturePromotion))
				}
			} else if s.Enpassant == to { //Captura al paso
				moves.Push(NewMove(piece, from, to, 0, 0, EnpassantCapture))
			} else {
				moves.Push(NewMove(piece, from, to, capture, 0, Capture)) //Capturas de peón
			}
			attacks &= attacks - 1
		}
		attacks = s.GetPawnPushes(from, color)
		// No hay capturas en los movimientos de peones
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			if (color == White && to-from == 16) || (color == Black && from-to == 16) { //Doble peón
				moves.Push(NewMove(piece, from, to, 0, 0, DoublePawnPush))
			} else if ((1<<to)&Rank1 != 0) || ((1<<to)&Rank8 != 0) { //Coronación de peón
				for _, promo := range piecePromotions {
					moves.Push(NewMove(piece, from, to, capture, promo, Promotion))
				}
			} else {
				moves.Push(NewMove(piece, from, to, capture, 0, QuietMoves))
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
			capture = None
			if s.IsCapture(to) {
				flag = Capture
				capture = s.GetPiece(to, enemy)
			} else {
				flag = QuietMoves
			}
			moves.Push(NewMove(piece, from, to, capture, 0, flag))
			attacks &= attacks - 1
		}
	case King:
		attacks = s.GetKingAttacks(from, color)
		for attacks != 0 {
			to = Square(bits.TrailingZeros64(attacks))
			capture = None
			var dist int8 = int8(to) - int8(from)
			if dist == 2 { // Enroque corto
				flag = KingCastle
			} else if dist == -2 { // Enroque largo
				flag = QueenCastle
			} else {
				flag = QuietMoves
				if s.IsCapture(to) {
					flag = Capture
					capture = s.GetPiece(to, enemy)
				}
			}
			moves.Push(NewMove(piece, from, to, capture, 0, flag))
			attacks &= attacks - 1
		}
	}
}

func (s *Board) GeneratePseudoCaptureSquare(moves *MoveList, square Square) {
	s.saveUnMakeInfo()
	var color uint8 = s.WhiteToMove
	var from Square
	var pieceBoard uint64
	var captureMask uint64 = 1 << square
	var attacks uint64
	var to Square
	var capture Piece
	var enemy = s.GetEnemyColor()
	for _, piece := range pieceTypes {
		pieceBoard = s.Bitboards[color][piece]
		for pieceBoard != 0 {
			from = Square(bits.TrailingZeros64(pieceBoard))
			switch piece {
			case Pawn:
				attacks = s.GetPawnAttacks(from, color) & captureMask
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = s.GetPiece(to, enemy)
					if ((1<<to)&Rank1 != 0) || ((1<<to)&Rank8 != 0) { //Coronación de peón
						for _, promo := range piecePromotions {
							moves.Push(NewMove(piece, from, to, capture, promo, CapturePromotion))
						}
					} else if s.Enpassant == to { //Captura al paso
						moves.Push(NewMove(piece, from, to, capture, 0, EnpassantCapture))
					} else {
						moves.Push(NewMove(piece, from, to, capture, 0, Capture)) //Capturas de peón
					}
					attacks &= attacks - 1
				}
			case Knight, Bishop, Rook, Queen:
				switch piece {
				case Knight:
					attacks = s.GetKnightAttacks(from, color) & captureMask
				case Bishop:
					attacks = s.GetBishopAttacks(from, color) & captureMask
				case Rook:
					attacks = s.GetRookAttacks(from, color) & captureMask
				case Queen:
					attacks = (s.GetBishopAttacks(from, color) | s.GetRookAttacks(from, color)) & captureMask
				}
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = None
					if s.IsCapture(to) {
						capture = s.GetPiece(to, enemy)
						moves.Push(NewMove(piece, from, to, capture, 0, Capture))
					}
					attacks &= attacks - 1
				}
			case King:
				attacks = s.GetKingAttacks(from, color) & captureMask
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = None
					if s.IsCapture(to) {
						capture = s.GetPiece(to, enemy)
						moves.Push(NewMove(piece, from, to, capture, 0, Capture))
					}
					attacks &= attacks - 1
				}
			}
			pieceBoard &= pieceBoard - 1
		}
	}
}

func (s *Board) GeneratePseudoCaptures(moves *MoveList) {
	s.saveUnMakeInfo()
	var color uint8 = s.WhiteToMove
	var from Square
	var pieceBoard uint64
	var attacks uint64
	var to Square
	var capture Piece
	var enemy = s.GetEnemyColor()
	for _, piece := range pieceTypes {
		pieceBoard = s.Bitboards[color][piece]
		for pieceBoard != 0 {
			from = Square(bits.TrailingZeros64(pieceBoard))
			switch piece {
			case Pawn:
				attacks = s.GetPawnAttacks(from, color)
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = s.GetPiece(to, enemy)
					if ((1<<to)&Rank1 != 0) || ((1<<to)&Rank8 != 0) { //Coronación de peón
						for _, promo := range piecePromotions {
							moves.Push(NewMove(piece, from, to, capture, promo, CapturePromotion))
						}
					} else if s.Enpassant == to { //Captura al paso
						moves.Push(NewMove(piece, from, to, capture, 0, EnpassantCapture))
					} else {
						moves.Push(NewMove(piece, from, to, capture, 0, Capture)) //Capturas de peón
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
					attacks = (s.GetBishopAttacks(from, color) | s.GetRookAttacks(from, color))
				}
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = None
					if s.IsCapture(to) {
						capture = s.GetPiece(to, enemy)
						moves.Push(NewMove(piece, from, to, capture, 0, Capture))
					}
					attacks &= attacks - 1
				}
			case King:
				attacks = s.GetKingAttacks(from, color)
				for attacks != 0 {
					to = Square(bits.TrailingZeros64(attacks))
					capture = None
					if s.IsCapture(to) {
						capture = s.GetPiece(to, enemy)
						moves.Push(NewMove(piece, from, to, capture, 0, Capture))
					}
					attacks &= attacks - 1
				}
			}
			pieceBoard &= pieceBoard - 1
		}
	}
}
