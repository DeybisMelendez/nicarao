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

type UnMove struct {
	//TODO: Intentar cambiar a Copy Make
	Piece      Piece
	PieceBB    uint64
	Special1   Piece
	Special1BB uint64
	Special2   Piece
	Special2BB uint64
	Flag       MoveFlag
	Enpassant  Square
	Castling   uint8
	Friends    uint64
	Enemies    uint64
	Occupied   uint64
}

func (s *Board) MakeMove(move *Move) *UnMove {
	var color bool = s.WhiteToMove
	var unMove UnMove = UnMove{
		Piece:     move.Piece,
		PieceBB:   s.Bitboards[color][move.Piece],
		Friends:   s.friends,
		Enemies:   s.enemies,
		Occupied:  s.occupied,
		Enpassant: s.Enpassant,
		Castling:  s.Castling,
	}
	s.Enpassant = 0
	unMove.Flag = move.Flag
	switch move.Flag {
	case QuietMoves:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
	case DoublePawnPush:
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		if color {
			s.Enpassant = move.To - 8
		} else {
			s.Enpassant = move.To + 8
		}
	case Capture:
		unMove.Special1 = move.Capture
		unMove.Special1BB = s.Bitboards[!color][move.Capture]
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
	case Promotion:
		unMove.Special2 = move.Promotion
		unMove.Special2BB = s.Bitboards[color][move.Promotion]
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
	case CapturePromotion:
		unMove.Special1 = move.Capture
		unMove.Special1BB = s.Bitboards[!color][move.Capture]
		unMove.Special2 = move.Promotion
		unMove.Special2BB = s.Bitboards[color][move.Promotion]
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
	case KingCastle:
		unMove.Special2 = Rook
		unMove.Special2BB = s.Bitboards[color][Rook]
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.HandleCastle(color, CastleShort, true)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H1), F1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H8), F8)
		}
	case QueenCastle:
		unMove.Special2 = Rook
		unMove.Special2BB = s.Bitboards[color][Rook]
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.HandleCastle(color, CastleLong, true)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A1), D1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A8), D8)
		}

	case EnpassantCapture:
		unMove.Special1 = Pawn //move.Capture
		unMove.Special1BB = s.Bitboards[!color][Pawn]
		s.Bitboards[color][move.Piece] = SetPopBit(s.Bitboards[color][move.Piece], move.To, move.From)
		//s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		if color {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To-8)
		} else {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To+8)
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
		if s.CanCastle(color, CastleLong) && ((move.From == A1 && s.WhiteToMove) || (move.From == A8 && !s.WhiteToMove)) {
			s.HandleCastle(color, CastleLong, true)
		}
		if s.CanCastle(color, CastleShort) && ((move.From == H1 && s.WhiteToMove) || (move.From == H8 && !s.WhiteToMove)) {
			s.HandleCastle(color, CastleShort, true)
		}
	}
	s.WhiteToMove = !color
	s.friends = s.GetAll(s.WhiteToMove)
	s.enemies = s.GetAll(!s.WhiteToMove)
	s.occupied = s.friends | s.enemies
	return &unMove
}

func (s *Board) UnMakeMove(move *UnMove) {
	var color bool = !s.WhiteToMove
	s.WhiteToMove = color
	s.Enpassant = move.Enpassant
	s.Castling = move.Castling
	s.friends = move.Friends
	s.enemies = move.Enemies
	s.occupied = move.Occupied
	s.Bitboards[color][move.Piece] = move.PieceBB

	switch move.Flag {
	case Capture, EnpassantCapture:
		s.Bitboards[!color][move.Special1] = move.Special1BB
	case Promotion, KingCastle, QueenCastle:
		s.Bitboards[color][move.Special2] = move.Special2BB
	case CapturePromotion:
		s.Bitboards[!color][move.Special1] = move.Special1BB
		s.Bitboards[color][move.Special2] = move.Special2BB
	}
	//s.friends = s.GetAll(s.WhiteToMove)
	//s.enemies = s.GetAll(!s.WhiteToMove)
	//s.occupied = s.friends | s.enemies
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

/*func (s *Board) IsMoveLegal(move *Move) bool {
	var color bool = s.WhiteToMove
	s.MakeMove(move)
	var kingBB uint64 = s.Bitboards[color][King]
	var result bool
	if !s.IsUnderAttack(kingBB, !color) {
		result = true
	}
	s.UnMakeMove()
	return result
}
*/
