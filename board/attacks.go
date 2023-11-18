package board

import (
	"math/bits"
)

func (s *Board) GetBishopAttacks(square Square, color bool) uint64 {
	var attacks uint64
	var blockerIndex Square
	attacks |= Rays["northWest"][square]
	if Rays["northWest"][square]&s.occupied != 0 {
		blockerIndex = Square(bits.TrailingZeros64(Rays["northWest"][square] & s.occupied))
		attacks &= ^Rays["northWest"][blockerIndex]
	}
	attacks |= Rays["northEast"][square]
	if Rays["northEast"][square]&s.occupied != 0 {
		blockerIndex = Square(bits.TrailingZeros64(Rays["northEast"][square] & s.occupied))
		attacks &= ^Rays["northEast"][blockerIndex]
	}
	attacks |= Rays["southWest"][square]
	if Rays["southWest"][square]&s.occupied != 0 {
		blockerIndex = 63 - Square(bits.LeadingZeros64(Rays["southWest"][square]&s.occupied))
		attacks &= ^Rays["southWest"][blockerIndex]
	}
	attacks |= Rays["southEast"][square]
	if Rays["southEast"][square]&s.occupied != 0 {
		blockerIndex = 63 - Square(bits.LeadingZeros64(Rays["southEast"][square]&s.occupied))
		attacks &= ^Rays["southEast"][blockerIndex]
	}
	return s.filterAttacks(attacks, color)
}
func (s *Board) GetRookAttacks(square Square, color bool) uint64 {
	var attacks uint64
	attacks |= Rays["north"][square]
	if Rays["north"][square]&s.occupied != 0 {
		var blockerIndex = Square(bits.TrailingZeros64(Rays["north"][square] & s.occupied))
		attacks &= ^Rays["north"][blockerIndex]
	}
	attacks |= Rays["east"][square]
	if Rays["east"][square]&s.occupied != 0 {
		var blockerIndex = Square(bits.TrailingZeros64(Rays["east"][square] & s.occupied))
		attacks &= ^Rays["east"][blockerIndex]
	}
	attacks |= Rays["south"][square]
	if Rays["south"][square]&s.occupied != 0 {
		var blockerIndex = 63 - Square(bits.LeadingZeros64(Rays["south"][square]&s.occupied))
		attacks &= ^Rays["south"][blockerIndex]
	}
	attacks |= Rays["west"][square]
	if Rays["west"][square]&s.occupied != 0 {
		var blockerIndex = 63 - Square(bits.LeadingZeros64(Rays["west"][square]&s.occupied))
		attacks &= ^Rays["west"][blockerIndex]
	}
	return s.filterAttacks(attacks, color)
}

func (s *Board) GetKnightAttacks(square Square, color bool) uint64 {
	return s.filterAttacks(KnightMasks[square], color)
}

func (s *Board) GetKingAttacks(square Square, color bool) uint64 {
	return s.filterAttacks(KingMasks[square], color)
}

func (s *Board) GetPawnPushes(square Square, color bool) uint64 {
	var mask uint64

	if color {
		mask = PawnWhitePushMasks[square]
		square += 8
		if SetBit(0, square)&s.occupied != 0 {
			mask = 0
		}
	} else {
		mask = PawnBlackPushMasks[square]
		square -= 8
		if SetBit(0, square)&s.occupied != 0 {
			mask = 0
		}
	}
	return mask & ^s.occupied
}
func (s *Board) GetPawnAttacks(square Square, color bool) uint64 {
	var mask uint64
	/*var enPassantMask uint64
	if s.Enpassant != 0 {
		enPassantMask = SetBit(0, s.Enpassant)
	}*/
	if color {
		mask |= PawnWhiteAttackMasks[square]
	} else {
		mask |= PawnBlackAttacksMasks[square]
	}
	if color == s.WhiteToMove {
		return mask & s.enemies
	}
	return mask & s.friends
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
