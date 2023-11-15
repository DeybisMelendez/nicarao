package board

import "math/bits"

func GetBishopAttacks(square Square, blockers uint64, friends uint64) uint64 {
	var attacks uint64
	if Rays["northWest"][square]&blockers != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["northWest"][square] & blockers))
		attacks |= Rays["northWest"][square] & ^Rays["northWest"][bitscanForward]
	} else {
		attacks |= Rays["northWest"][square]
	}
	if Rays["northEast"][square]&blockers != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["northEast"][square] & blockers))
		attacks |= Rays["northEast"][square] & ^Rays["northEast"][bitscanForward]
	} else {
		attacks |= Rays["northEast"][square]
	}
	if Rays["southWest"][square]&blockers != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["southWest"][square]&blockers))
		attacks |= Rays["southWest"][square] & ^Rays["southWest"][bitscanForward]
	} else {
		attacks |= Rays["southWest"][square]
	}
	if Rays["southEast"][square]&blockers != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["southEast"][square]&blockers))
		attacks |= Rays["southEast"][square] & ^Rays["southEast"][bitscanForward]
	} else {
		attacks |= Rays["southEast"][square]
	}
	return attacks & ^friends
}
func GetRookAttacks(square Square, blockers uint64, friends uint64) uint64 {
	var attacks uint64
	if Rays["north"][square]&blockers != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["north"][square] & blockers))
		attacks |= Rays["north"][square] & ^Rays["north"][bitscanForward]
	} else {
		attacks |= Rays["north"][square]
	}
	if Rays["east"][square]&blockers != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["east"][square] & blockers))
		attacks |= Rays["east"][square] & ^Rays["east"][bitscanForward]
	} else {
		attacks |= Rays["east"][square]
	}
	if Rays["south"][square]&blockers != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["south"][square]&blockers))
		attacks |= Rays["south"][square] & ^Rays["south"][bitscanForward]
	} else {
		attacks |= Rays["south"][square]
	}
	if Rays["west"][square]&blockers != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["west"][square]&blockers))
		attacks |= Rays["west"][square] & ^Rays["west"][bitscanForward]
	} else {
		attacks |= Rays["west"][square]
	}
	return attacks & ^friends
}

func GetKnightAttacks(square Square, friends uint64) uint64 {
	return KnightMasks[square] & ^friends
}

func GetKingAttacks(square Square, friends uint64) uint64 {
	return KingMasks[square] & ^friends
}

func GetPawnPushes(isWhite bool, square Square, blockers uint64) uint64 {
	if isWhite {
		return PawnWhitePushMasks[square] & ^blockers
	}
	return PawnBlackPushMasks[square] & ^blockers
}
func GetPawnAttacks(isWhite bool, square Square, enemies uint64) uint64 {
	if isWhite {
		return PawnWhiteAttackMasks[square] & enemies
	}
	return PawnBlackAttacksMasks[square] & enemies
}
