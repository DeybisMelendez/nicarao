package board

import "math/bits"

type Board struct {
	WhiteToMove    bool
	Bitboards      map[bool]map[Piece]uint64 // Mapa para almacenar bitboards por color y tipo de pieza
	Enpassant      uint8                     // square id (16-23 or 40-47) where en passant capture is possible
	CastlingRights uint8
	Halfmoveclock  uint8
}

func NewBoard() *Board {
	return &Board{
		Bitboards: map[bool]map[Piece]uint64{
			White: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
			Black: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
		},
	}
}

func (s *Board) GetOccupied() uint64 {
	return s.Bitboards[White][Pawn] | s.Bitboards[White][Bishop] | s.Bitboards[White][Knight] |
		s.Bitboards[White][Rook] | s.Bitboards[White][Queen] | s.Bitboards[White][King] |
		s.Bitboards[Black][Pawn] | s.Bitboards[Black][Bishop] | s.Bitboards[Black][Knight] |
		s.Bitboards[Black][Rook] | s.Bitboards[Black][Queen] | s.Bitboards[Black][King]
}
func (s *Board) GetAll(color bool) uint64 {
	var pieceTypes = []Piece{Pawn, Bishop, Knight, Rook, Queen, King}
	var total uint64
	for _, piece := range pieceTypes {
		total |= s.Bitboards[color][piece]
	}
	return total
}

func (board *Board) GetPiece(square Square, color bool) Piece {
	var mask uint64 = SetBit(0, square)

	for _, piece := range pieceTypes {
		if board.Bitboards[color][piece]&mask != 0 {
			return piece
		}
	}
	return None
}

func (s *Board) IsUnderAttack(pieceBB uint64, enemies, occupied, friends uint64) bool {

	//var pieceBB uint64 = SetBit(0, square)
	//var kingSquare Square = Square(bits.TrailingZeros64(kingBB))

	for _, piece := range pieceTypes {
		/*if piece == King {
			continue
		}*/

		enemyPieces := s.Bitboards[!s.WhiteToMove][piece]
		for enemyPieces != 0 {
			from := Square(bits.TrailingZeros64(enemyPieces))
			attacks := s.GenerateAttacksForPiece(piece, from, occupied, enemies, friends)

			if attacks&pieceBB != 0 {
				return true
			}
			enemyPieces &= enemyPieces - 1
		}
	}
	return false
}
func (b *Board) AnyUnderAttack(enemies uint64, occupied uint64, friends uint64, squares ...Square) bool {
	for _, square := range squares {
		pieceBB := SetBit(0, square)
		if b.IsUnderAttack(pieceBB, enemies, occupied, friends) {
			return true
		}
	}
	return false
}

func (board *Board) CanCastleShort(isWhite bool) bool {
	castleRight := CastlingWhiteShort
	if !isWhite {
		castleRight = CastlingBlackShort
	}
	return board.CastlingRights&castleRight != 0
}

func (board *Board) CanCastleLong(isWhite bool) bool {
	castleRight := CastlingWhiteLong
	if !isWhite {
		castleRight = CastlingBlackLong
	}
	return board.CastlingRights&castleRight != 0
}

func (board *Board) UpdateCastlingRights(castleRight uint8) {
	board.CastlingRights |= castleRight
}

func (board *Board) RemoveCastlingRights(castleRight uint8) {
	board.CastlingRights &= ^castleRight
}

func (s *Board) HandleCastle(isWhite bool, isShort bool, remove bool) {
	var castle uint8
	if isWhite {
		if isShort {
			castle = CastlingWhiteShort
		} else {
			castle = CastlingWhiteLong
		}
	} else {
		if isShort {
			castle = CastlingBlackShort
		} else {
			castle = CastlingBlackLong
		}
	}
	if remove {
		s.RemoveCastlingRights(castle)
	} else {
		s.UpdateCastlingRights(castle)
	}
}
