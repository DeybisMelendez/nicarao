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
	//TODO: Intentar reducir el tamaño del struct al mínimo necesario
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
	// Ajuste del flag KingCastle y QueenCastle
	if move.Piece == King {
		var dist int = int(move.To) - int(move.From)
		if dist == 2 { // Enroque corto
			move.Flag = KingCastle
		} else if dist == -2 { // Enroque largo
			move.Flag = QueenCastle
		}
	}
	unMove.Flag = move.Flag
	switch move.Flag {
	case QuietMoves:
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
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
	case DoublePawnPush:
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		if color {
			s.Enpassant = move.To - 8
		} else {
			s.Enpassant = move.To + 8
		}
	case Capture:
		unMove.Special1 = move.Capture
		unMove.Special1BB = s.Bitboards[!color][move.Capture]
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
		if move.Capture == Rook {
			if move.To == A1 || move.To == A8 {
				s.HandleCastle(!color, CastleLong, true)
			} else if move.To == H1 || move.To == H8 {
				s.HandleCastle(!color, CastleShort, true)
			}
		}
	case Promotion:
		unMove.Special1 = move.Promotion
		unMove.Special1BB = s.Bitboards[color][move.Promotion]
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
	case CapturePromotion:
		unMove.Special1 = move.Promotion
		unMove.Special1BB = s.Bitboards[color][move.Promotion]
		unMove.Special2 = move.Capture
		unMove.Special2BB = s.Bitboards[!color][move.Capture]
		s.Bitboards[color][move.Piece] = PopBit(s.Bitboards[color][move.Piece], move.To)
		s.Bitboards[color][move.Promotion] = SetBit(s.Bitboards[color][move.Promotion], move.To)
		s.Bitboards[!color][move.Capture] = PopBit(s.Bitboards[!color][move.Capture], move.To)
	case KingCastle:
		unMove.Special1 = Rook
		unMove.Special1BB = s.Bitboards[color][Rook]
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.HandleCastle(color, CastleShort, true)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H1), F1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], H8), F8)
		}
	case QueenCastle:
		unMove.Special1 = Rook
		unMove.Special1BB = s.Bitboards[color][Rook]
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		s.HandleCastle(color, CastleLong, true)
		if color {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A1), D1)
		} else {
			s.Bitboards[color][Rook] = SetBit(PopBit(s.Bitboards[color][Rook], A8), D8)
		}

	case EnpassantCapture:
		unMove.Special1 = Pawn
		unMove.Special1BB = s.Bitboards[!color][Pawn]
		s.Bitboards[color][move.Piece] = SetBit(PopBit(s.Bitboards[color][move.Piece], move.From), move.To)
		if color {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To-8)
		} else {
			s.Bitboards[!color][Pawn] = PopBit(s.Bitboards[!color][Pawn], move.To+8)
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
		s.Bitboards[color][move.Special1] = move.Special1BB
	case CapturePromotion:
		s.Bitboards[color][move.Special1] = move.Special1BB
		s.Bitboards[!color][move.Special2] = move.Special2BB
	}

}

func (s *Board) GeneratePseudoMoves() []Move {
	var moves []Move
	var color bool = s.WhiteToMove
	for _, piece := range pieceTypes {
		var pieceBoard uint64 = s.Bitboards[color][piece]

		for pieceBoard != 0 {
			var from Square = Square(bits.TrailingZeros64(pieceBoard))
			var attacks uint64 = s.GenerateAttacksForPiece(piece, from, s.WhiteToMove)
			var flag MoveFlag
			if piece == King {
				var kingsideOK bool = s.occupied&castlingMask[color][CastleShort] == 0 &&
					s.CanCastle(color, CastleShort) &&
					!s.AnyUnderAttack(!s.WhiteToMove, castlingSquares[color][CastleShort]...)
				var queensideOK bool = s.occupied&castlingMask[color][CastleLong] == 0 &&
					s.CanCastle(color, CastleLong) &&
					!s.AnyUnderAttack(!s.WhiteToMove, castlingSquares[color][CastleLong]...)
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
				if capture == None {
					flag = QuietMoves
				} else {
					flag = Capture
				}
				if piece == Pawn {
					if to < A2 || to > H7 { //Promotion
						for _, promotion := range piecePromotions {
							if capture == None {
								flag = Promotion
							} else {
								flag = CapturePromotion
							}
							moves = append(moves, Move{From: from, To: to, Piece: piece,
								Capture: capture, Promotion: promotion, Flag: flag})
						}
					} else if color && to-from == 16 { //Enpassant blancas
						moves = append(moves, Move{From: from, To: to,
							Piece: piece, Flag: DoublePawnPush, Capture: capture})
					} else if !color && from-to == 16 { //Enpassant negras
						moves = append(moves, Move{From: from, To: to,
							Piece: piece, Flag: DoublePawnPush, Capture: capture})
					} else if s.Enpassant != 0 && s.Enpassant == to { //Enpassant Capture
						moves = append(moves, Move{From: from, To: to,
							Piece: piece, Capture: capture, Flag: EnpassantCapture})
					} else {
						moves = append(moves, Move{From: from, To: to, //Pawn move
							Piece: piece, Flag: flag, Capture: capture}) // en teoría no necesita Capture
					}
				} else {
					moves = append(moves, Move{From: from, To: to,
						Piece: piece, Capture: capture, Flag: flag})
				}
				attacks &= attacks - 1
			}
			pieceBoard &= pieceBoard - 1
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
