package board

import (
	"math/bits"
)

func (s *Board) GetBishopAttacks(square Square) uint64 {
	var attacks uint64
	if Rays["northWest"][square]&s.occupied != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["northWest"][square] & s.occupied))
		attacks |= Rays["northWest"][square] & ^Rays["northWest"][bitscanForward]
	} else {
		attacks |= Rays["northWest"][square]
	}
	if Rays["northEast"][square]&s.occupied != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["northEast"][square] & s.occupied))
		attacks |= Rays["northEast"][square] & ^Rays["northEast"][bitscanForward]
	} else {
		attacks |= Rays["northEast"][square]
	}
	if Rays["southWest"][square]&s.occupied != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["southWest"][square]&s.occupied))
		attacks |= Rays["southWest"][square] & ^Rays["southWest"][bitscanForward]
	} else {
		attacks |= Rays["southWest"][square]
	}
	if Rays["southEast"][square]&s.occupied != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["southEast"][square]&s.occupied))
		attacks |= Rays["southEast"][square] & ^Rays["southEast"][bitscanForward]
	} else {
		attacks |= Rays["southEast"][square]
	}
	return attacks & ^s.friends
}
func (s *Board) GetRookAttacks(square Square) uint64 {
	var attacks uint64
	if Rays["north"][square]&s.occupied != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["north"][square] & s.occupied))
		attacks |= Rays["north"][square] & ^Rays["north"][bitscanForward]
	} else {
		attacks |= Rays["north"][square]
	}
	if Rays["east"][square]&s.occupied != 0 {
		var bitscanForward = Square(bits.TrailingZeros64(Rays["east"][square] & s.occupied))
		attacks |= Rays["east"][square] & ^Rays["east"][bitscanForward]
	} else {
		attacks |= Rays["east"][square]
	}
	if Rays["south"][square]&s.occupied != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["south"][square]&s.occupied))
		attacks |= Rays["south"][square] & ^Rays["south"][bitscanForward]
	} else {
		attacks |= Rays["south"][square]
	}
	if Rays["west"][square]&s.occupied != 0 {
		var bitscanForward = 63 - Square(bits.LeadingZeros64(Rays["west"][square]&s.occupied))
		attacks |= Rays["west"][square] & ^Rays["west"][bitscanForward]
	} else {
		attacks |= Rays["west"][square]
	}
	return attacks & ^s.friends
}

func (s *Board) GetKnightAttacks(square Square) uint64 {
	return KnightMasks[square] & ^s.friends
}

func (s *Board) GetKingAttacks(square Square) uint64 {
	return KingMasks[square] & ^s.friends
}

func (s *Board) GetPawnPushes(square Square) uint64 {
	if s.WhiteToMove {
		return PawnWhitePushMasks[square] & ^s.occupied
	}
	return PawnBlackPushMasks[square] & ^s.occupied
}
func (s *Board) GetPawnAttacks(square Square) uint64 {
	if !s.WhiteToMove { //Revisar por qué pasa esto, si quito el ! los peones dejan de atacar correctamente (invertido)
		return PawnWhiteAttackMasks[square] & s.enemies
	}
	return PawnBlackAttacksMasks[square] & s.enemies
}

func (s *Board) GenerateAttacksForPiece(piece Piece, from Square) uint64 {
	//var attacks uint64

	switch piece {
	case Pawn:
		return s.GetPawnAttacks(from) | s.GetPawnPushes(from)
	case Knight:
		return s.GetKnightAttacks(from)
	case Bishop:
		return s.GetBishopAttacks(from)
	case Rook:
		return s.GetRookAttacks(from)
	case Queen:
		return s.GetBishopAttacks(from) | s.GetRookAttacks(from)
	case King:
		return s.GetKingAttacks(from)
	}
	return 0
}

func (s *Board) IsUnderAttack(pieceBB uint64) bool {

	for _, piece := range pieceTypes {
		/*if piece == King {
			continue
		}*/

		enemyPieces := s.Bitboards[!s.WhiteToMove][piece]
		for enemyPieces != 0 {
			from := Square(bits.TrailingZeros64(enemyPieces))
			attacks := s.GenerateAttacksForPiece(piece, from)
			//PrintBitboard(attacks)
			if attacks&pieceBB != 0 {
				return true
			}
			enemyPieces &= enemyPieces - 1
		}
	}
	return false
}
func (s *Board) AnyUnderAttack(squares ...Square) bool {
	for _, square := range squares {
		pieceBB := SetBit(0, square)
		if s.IsUnderAttack(pieceBB) {
			return true
		}
	}
	return false
}
