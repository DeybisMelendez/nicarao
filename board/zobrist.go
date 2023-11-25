package board

import (
	"math/bits"
	"math/rand"
)

func generateZobristConstants() {
	whiteToMoveZobrist = rand.Uint64()
	for i := 0; i < 2; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 64; k++ {
				pieceSquareZobrist[i][j][k] = rand.Uint64()
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			castleRightsZobrist[i][j] = rand.Uint64()
		}
	}
}

func (b *Board) calcZobristHash() {
	var hash uint64 = 0
	var square int
	var pieces uint64
	if b.WhiteToMove == White {
		hash ^= whiteToMoveZobrist
	}
	if b.CanCastle(White, CastleShort) {
		hash ^= castleRightsZobrist[White][CastleShort]
	}
	if b.CanCastle(White, CastleLong) {
		hash ^= castleRightsZobrist[White][CastleLong]
	}
	if b.CanCastle(Black, CastleShort) {
		hash ^= castleRightsZobrist[Black][CastleShort]
	}
	if b.CanCastle(Black, CastleLong) {
		hash ^= castleRightsZobrist[Black][CastleLong]
	}
	hash ^= uint64(b.Enpassant)
	for color := 0; color < 2; color++ {
		for _, piece := range pieceTypes {
			pieces = b.Bitboards[color][piece]
			for pieces != 0 {
				square = bits.TrailingZeros64(pieces)
				hash ^= pieceSquareZobrist[color][piece][square]
				pieces &= pieces - 1
			}
		}
	}
	b.Hash = hash
}
