package board

import (
	"math/bits"
)

func (s *Board) GetBishopAttacks(square Square, color bool) uint64 {
	var attacks uint64 = Rays[northWest][square]
	var blocker int = bits.TrailingZeros64((Rays[northWest][square] & s.occupied) | 0x8000000000000000)
	attacks &= ^Rays[northWest][blocker]

	attacks |= Rays[northEast][square]
	blocker = bits.TrailingZeros64((Rays[northEast][square] & s.occupied) | 0x8000000000000000)
	attacks &= ^Rays[northEast][blocker]

	attacks |= Rays[southWest][square]
	blocker = 63 - bits.LeadingZeros64((Rays[southWest][square]&s.occupied)|1)
	attacks &= ^Rays[southWest][blocker]

	attacks |= Rays[southEast][square]
	blocker = 63 - bits.LeadingZeros64((Rays[southEast][square]&s.occupied)|1)
	attacks &= ^Rays[southEast][blocker]
	if color == s.WhiteToMove {
		return attacks & ^s.friends
	}
	return attacks & ^s.enemies
}
func (s *Board) GetRookAttacks(square Square, color bool) uint64 {
	var attacks uint64 = Rays[north][square]
	var blocker int = bits.TrailingZeros64((Rays[north][square] & s.occupied) | 0x8000000000000000)
	attacks &= ^Rays[north][blocker]

	attacks |= Rays[east][square]
	blocker = bits.TrailingZeros64((Rays[east][square] & s.occupied) | 0x8000000000000000)
	attacks &= ^Rays[east][blocker]

	attacks |= Rays[south][square]
	blocker = 63 - bits.LeadingZeros64((Rays[south][square]&s.occupied)|1)
	attacks &= ^Rays[south][blocker]

	attacks |= Rays[west][square]
	blocker = 63 - bits.LeadingZeros64((Rays[west][square]&s.occupied)|1)
	attacks &= ^Rays[west][blocker]

	if color == s.WhiteToMove {
		return attacks & ^s.friends
	}
	return attacks & ^s.enemies
}

func (s *Board) GetKnightAttacks(square Square, color bool) uint64 {
	if color == s.WhiteToMove {
		return KnightMasks[square] & ^s.friends
	}
	return KnightMasks[square] & ^s.enemies
}

func (s *Board) GetKingAttacks(square Square, color bool) uint64 {
	var attacks uint64 = KingMasks[square]
	var kingsideOK bool = s.CanCastle(color, CastleShort) &&
		s.occupied&castlingMask[color][CastleShort] == 0 &&
		!s.AnyUnderAttack(color, castlingSquares[color][CastleShort]...)
	var queensideOK bool = s.CanCastle(color, CastleLong) &&
		s.occupied&castlingMask[color][CastleLong] == 0 &&
		!s.AnyUnderAttack(color, castlingSquares[color][CastleLong]...)
	if kingsideOK {
		attacks |= (1 << castlingSquares[color][CastleShort][1]) //= SetBit(attacks, castlingSquares[color][CastleShort][1])
	}
	if queensideOK {
		attacks |= (1 << castlingSquares[color][CastleLong][0]) //= SetBit(attacks, castlingSquares[color][CastleLong][0])
	}
	if color == s.WhiteToMove {
		return attacks & ^s.friends
	}
	return attacks & ^s.enemies
}

func (s *Board) GetPawnPushes(square Square, color bool) uint64 {
	var mask uint64 = PawnPushesMasks[color][square]
	if color {
		square += 8
	} else {
		square -= 8
	}
	if (1<<square)&s.occupied != 0 {
		return 0
	}
	return mask & ^s.occupied
}

func (s *Board) GetPawnAttacks(square Square, color bool) uint64 {
	var enPassantMask uint64
	if s.Enpassant != 0 && color == s.WhiteToMove {
		enPassantMask = 1 << s.Enpassant
	}
	if color == s.WhiteToMove {
		return PawnAttacksMasks[color][square] & (s.enemies | enPassantMask)
	}
	return PawnAttacksMasks[color][square] & (s.friends | enPassantMask)
}

/*func (s *Board) GenerateAttacksForPiece(piece Piece, from Square, color bool) uint64 {

	switch piece {
	case Pawn:
		return s.GetPawnAttacks(from, color) | s.GetPawnPushes(from, color)
	case Knight:
		return s.GetKnightAttacks(from, color)
	case Bishop:
		return s.GetBishopAttacks(from, color)
	case Rook:
		return s.GetRookAttacks(from, color)
	case Queen:
		return s.GetBishopAttacks(from, color) | s.GetRookAttacks(from, color)
	case King:
		return s.GetKingAttacks(from, color)
	}
	return 0
}*/

func (s *Board) IsUnderAttack(square Square, color bool) bool {
	if s.GetBishopAttacks(square, color)&(s.Bitboards[!color][Bishop]|s.Bitboards[!color][Queen]) != 0 {
		return true
	}
	if s.GetRookAttacks(square, color)&(s.Bitboards[!color][Rook]|s.Bitboards[!color][Queen]) != 0 {
		return true
	}
	if s.GetKnightAttacks(square, color)&s.Bitboards[!color][Knight] != 0 {
		return true
	}
	if s.GetPawnAttacks(square, color)&s.Bitboards[!color][Pawn] != 0 {
		return true
	}
	return KingMasks[square]&s.Bitboards[!color][King] != 0
}
func (s *Board) AnyUnderAttack(color bool, squares ...Square) bool {
	for _, square := range squares {
		if s.IsUnderAttack(square, color) {
			return true
		}
	}
	return false
}
