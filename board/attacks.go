package board

import (
	"math/bits"
)

func (s *Board) GetBishopAttacks(square Square, color bool) uint64 {
	var attacks uint64
	var blockerIndex int
	var occluded uint64
	attacks |= Rays["northWest"][square]
	if occluded = Rays["northWest"][square] & s.occupied; occluded != 0 {
		blockerIndex = bits.TrailingZeros64(occluded)
		attacks &= ^Rays["northWest"][blockerIndex]
	}
	attacks |= Rays["northEast"][square]
	if occluded = Rays["northEast"][square] & s.occupied; occluded != 0 {
		blockerIndex = bits.TrailingZeros64(occluded)
		attacks &= ^Rays["northEast"][blockerIndex]
	}
	attacks |= Rays["southWest"][square]
	if occluded = Rays["southWest"][square] & s.occupied; occluded != 0 {
		blockerIndex = 63 - bits.LeadingZeros64(occluded)
		attacks &= ^Rays["southWest"][blockerIndex]
	}
	attacks |= Rays["southEast"][square]
	if occluded = Rays["southEast"][square] & s.occupied; occluded != 0 {
		blockerIndex = 63 - bits.LeadingZeros64(occluded)
		attacks &= ^Rays["southEast"][blockerIndex]
	}
	return s.filterAttacks(attacks, color)
}
func (s *Board) GetRookAttacks(square Square, color bool) uint64 {
	var attacks uint64
	var blockerIndex int
	var occluded uint64
	attacks |= Rays["north"][square]
	if occluded = Rays["north"][square] & s.occupied; occluded != 0 {
		blockerIndex = bits.TrailingZeros64(occluded)
		attacks &= ^Rays["north"][blockerIndex]
	}
	attacks |= Rays["east"][square]
	if occluded = Rays["east"][square] & s.occupied; occluded != 0 {
		blockerIndex = bits.TrailingZeros64(occluded)
		attacks &= ^Rays["east"][blockerIndex]
	}
	attacks |= Rays["south"][square]
	if occluded = Rays["south"][square] & s.occupied; occluded != 0 {
		blockerIndex = 63 - bits.LeadingZeros64(occluded)
		attacks &= ^Rays["south"][blockerIndex]
	}
	attacks |= Rays["west"][square]
	if occluded = Rays["west"][square] & s.occupied; occluded != 0 {
		blockerIndex = 63 - bits.LeadingZeros64(occluded)
		attacks &= ^Rays["west"][blockerIndex]
	}
	return s.filterAttacks(attacks, color)
}

func (s *Board) GetKnightAttacks(square Square, color bool) uint64 {
	return s.filterAttacks(KnightMasks[square], color)
}

func (s *Board) GetKingAttacks(square Square, color bool) uint64 {
	var attacks uint64 = KingMasks[square]
	var kingsideOK bool = s.CanCastle(color, CastleShort) &&
		s.occupied&castlingMask[color][CastleShort] == 0 &&
		!s.AnyUnderAttack(!s.WhiteToMove, castlingSquares[color][CastleShort]...)
	var queensideOK bool = s.CanCastle(color, CastleLong) &&
		s.occupied&castlingMask[color][CastleLong] == 0 &&
		!s.AnyUnderAttack(!s.WhiteToMove, castlingSquares[color][CastleLong]...)
	if kingsideOK {
		attacks = SetBit(attacks, castlingSquares[color][CastleShort][1])
	}
	if queensideOK {
		attacks = SetBit(attacks, castlingSquares[color][CastleLong][0])
	}
	return s.filterAttacks(attacks, color)
}

func (s *Board) GetPawnPushes(square Square, color bool) uint64 {
	var mask uint64 = PawnPushesMasks[color][square]
	if color {
		square += 8
	} else {
		square -= 8
	}
	if SetBit(0, square)&s.occupied != 0 {
		return 0
	}
	return mask & ^s.occupied
}

func (s *Board) GetPawnAttacks(square Square, color bool) uint64 {
	var enPassantMask uint64
	if s.Enpassant != 0 {
		enPassantMask = SetBit(0, s.Enpassant)
	}
	if color == s.WhiteToMove {
		return PawnAttacksMasks[color][square] & (s.enemies | enPassantMask)
	}
	return PawnAttacksMasks[color][square] & (s.friends | enPassantMask)
}

func (s *Board) GenerateAttacksForPiece(piece Piece, from Square, color bool) uint64 {

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
}

func (s *Board) IsUnderAttack(pieceBB uint64, color bool) bool {

	for _, piece := range pieceTypes {

		var attackerBB uint64 = s.Bitboards[color][piece]
		//TODO: Podría ser mas eficiente precalcular los ataques en la generación de movimientos y guardarlo en el board
		for attackerBB != 0 {
			var from Square = Square(bits.TrailingZeros64(attackerBB))
			var attacks uint64 = s.GenerateAttacksForPiece(piece, from, color)
			if attacks&pieceBB != 0 {
				return true
			}
			attackerBB &= attackerBB - 1
		}
	}
	return false
}
func (s *Board) AnyUnderAttack(color bool, squares ...Square) bool {
	for _, square := range squares {
		var pieceBB uint64 = SetBit(0, square)
		if s.IsUnderAttack(pieceBB, color) {
			return true
		}
	}
	return false
}
func (s *Board) filterAttacks(attacks uint64, color bool) uint64 {
	if color == s.WhiteToMove {
		return attacks & ^s.friends
	}
	return attacks & ^s.enemies
}
