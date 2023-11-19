package board

type Square uint8
type Piece uint8

func init() {
	castling = map[bool]map[bool]uint8{
		White: {CastleShort: 2, CastleLong: 8},
		Black: {CastleShort: 1, CastleLong: 4},
	}
	castlingMask = map[bool]map[bool]uint64{
		White: {CastleShort: 0x60, CastleLong: 0xe},
		Black: {CastleShort: 0x6000000000000000, CastleLong: 0xe00000000000000},
	}
	castlingSquares = map[bool]map[bool][]Square{
		White: {CastleShort: {F1, G1}, CastleLong: {C1, D1}},
		Black: {CastleShort: {F8, G8}, CastleLong: {C8, D8}},
	}
}

var pieceTypes = []Piece{Pawn, Knight, Bishop, Rook, Queen, King}
var piecePromotions = []Piece{Queen, Rook, Bishop, Knight}
var castling map[bool]map[bool]uint8
var castlingMask map[bool]map[bool]uint64
var castlingSquares map[bool]map[bool][]Square

const White bool = true
const Black bool = false
const CastleShort bool = true
const CastleLong bool = false
const StartingPos string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

//const everything uint64 = ^(uint64(0))
const (
	None Piece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)
const (
	QuietMoves MoveFlag = iota
	DoublePawnPush
	Capture
	Promotion
	CapturePromotion
	KingCastle
	QueenCastle
	EnpassantCapture
)
const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

const FileA uint64 = 0x0101010101010101
const FileH uint64 = 0x8080808080808080
const Rank1 uint64 = 0x00000000000000FF
const Rank8 uint64 = 0xFF00000000000000
