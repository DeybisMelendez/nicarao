package board

import "math/bits"

type Board struct {
	WhiteToMove bool
	Bitboards   map[bool]map[Piece]uint64 // Mapa para almacenar bitboards por color y tipo de pieza
}

func NewBoard() *Board {
	return &Board{
		Bitboards: map[bool]map[Piece]uint64{
			true: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
			false: {
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
	return s.Bitboards[true][Pawn] | s.Bitboards[true][Bishop] | s.Bitboards[true][Knight] |
		s.Bitboards[true][Rook] | s.Bitboards[true][Queen] | s.Bitboards[true][King] |
		s.Bitboards[false][Pawn] | s.Bitboards[false][Bishop] | s.Bitboards[false][Knight] |
		s.Bitboards[false][Rook] | s.Bitboards[false][Queen] | s.Bitboards[false][King]
}
func (s *Board) GetAll(isWhite bool) uint64 {
	var pieceTypes = []Piece{Pawn, Bishop, Knight, Rook, Queen, King}
	var total uint64
	for _, piece := range pieceTypes {
		total |= s.Bitboards[isWhite][piece]
	}
	return total
}
func (s *Board) MakeMove(move Move) {
	color := s.WhiteToMove
	piece := s.Bitboards[color][move.Piece]
	capture := s.Bitboards[!color][move.Capture]

	piece = SetBit(PopBit(piece, move.From), move.To)

	if move.Capture != None {
		s.Bitboards[!color][move.Capture] = PopBit(capture, move.To)
	}

	s.Bitboards[color][move.Piece] = piece
	s.WhiteToMove = !s.WhiteToMove
}

func (s *Board) UnMakeMove(move Move) {
	color := !s.WhiteToMove
	piece := s.Bitboards[color][move.Piece]
	capture := s.Bitboards[!color][move.Capture]

	piece = SetBit(PopBit(piece, move.To), move.From)

	if move.Capture != None {
		s.Bitboards[!color][move.Capture] = SetBit(capture, move.To)
	}

	s.Bitboards[color][move.Piece] = piece
	s.WhiteToMove = !s.WhiteToMove
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

func (s *Board) generateAttacksForPiece(piece Piece, from Square, isWhite bool, occupied, friends, enemies uint64) uint64 {
	//var attacks uint64

	switch piece {
	case Pawn:
		return GetPawnAttacks(isWhite, from, enemies)
	case Knight:
		return GetKnightAttacks(from, friends)
	case Bishop:
		return GetBishopAttacks(from, occupied, friends)
	case Rook:
		return GetRookAttacks(from, occupied, friends)
	case Queen:
		return GetBishopAttacks(from, occupied, friends) | GetRookAttacks(from, occupied, friends)
	case King:
		return GetKingAttacks(from, friends)
	}

	return 0
}

func (board *Board) GeneratePseudoLegalMoves() []Move {
	var moves []Move
	isWhite := board.WhiteToMove
	friends := board.GetAll(isWhite)
	enemies := board.GetAll(!isWhite)
	occupied := board.GetOccupied()

	for _, piece := range pieceTypes {
		pieceBoard := board.Bitboards[isWhite][piece]

		for pieceBoard != 0 {
			from := Square(bits.TrailingZeros64(pieceBoard))
			attacks := board.generateAttacksForPiece(piece, from, isWhite, occupied, friends, enemies)
			if piece == Pawn {
				attacks |= GetPawnPushes(isWhite, from, occupied)
			}
			for attacks != 0 {
				to := Square(bits.TrailingZeros64(attacks))
				capture := board.GetPiece(to, !board.WhiteToMove)
				if capture != King {
					move := Move{From: from, To: to, Piece: piece, Capture: capture, IsWhite: isWhite}
					board.MakeMove(move)
					if !board.IsKingInCheck(isWhite, enemies, occupied, friends) {
						// TODO: Agregar una evaluación del movimiento
						moves = append(moves, move)
					}
					board.UnMakeMove(move)
				}
				attacks &= attacks - 1
			}
			pieceBoard &= pieceBoard - 1
		}
	}

	return moves
}

func (s *Board) IsKingInCheck(isWhite bool, enemies, occupied, friends uint64) bool {

	var kingBB uint64 = s.Bitboards[isWhite][King]
	var kingSquare Square = Square(bits.TrailingZeros64(kingBB))

	for _, piece := range pieceTypes {
		/*if piece == King {
			continue
		}*/

		enemyPieces := s.Bitboards[!isWhite][piece]
		for enemyPieces != 0 {
			from := Square(bits.TrailingZeros64(enemyPieces))
			attacks := s.generateAttacksForPiece(piece, from, !isWhite, occupied, enemies, friends)

			if attacks&(SetBit(0, kingSquare)) != 0 {
				return true
			}
			enemyPieces &= enemyPieces - 1
		}
	}
	return false
}
