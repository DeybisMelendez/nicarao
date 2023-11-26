package board

type MoveFlag uint8
type Move uint32

func NewMove(piece Piece, from Square, to Square, capture Piece, promo Piece, flag MoveFlag) Move {
	return Move((uint32(piece) & 0x7) | // 3 bits Representa piezas 0-7
		((uint32(from) & 0x3f) << 3) | // 6 bits Representa casillas 0-63
		((uint32(to) & 0x3f) << 9) | // 6 bits Representa casillas 0-63
		((uint32(capture) & 0x7) << 15) | // 3 bits Representa piezas 0-7
		((uint32(promo) & 0x7) << 18) | // 3 bits Representa piezas 0-7
		((uint32(flag) & 0x7) << 21)) // 3 bits Representa flags 0-7
	//8 bits restantes pueden aprovecharse para move ordering
}

//GETTERS
func (m *Move) Piece() Piece {
	return Piece(*m & 0x7)
}

func (m *Move) From() Square {
	var sq = Square((*m >> 3) & 0x3F)
	return sq
}

func (m *Move) To() Square {
	var sq = Square((*m >> 9) & 0x3F)
	return sq
}

func (m *Move) Capture() Piece {
	return Piece((*m >> 15) & 0x7)
}

func (m *Move) Promotion() Piece {
	return Piece((*m >> 18) & 0x7)
}

func (m *Move) Flag() MoveFlag {
	return MoveFlag((*m >> 21) & 0x7)
}

func (m *Move) GetScore() uint8 {
	return uint8((*m >> 24) & 0xFF)
}

//SETTERS

func (m *Move) SetPiece(piece Piece) {
	*m = Move(uint32(piece)&0x7 | (uint32(*m) &^ 0x7))
}

func (m *Move) SetFrom(square Square) {
	*m = Move((uint32(square) & 0x3F << 3) | (uint32(*m) &^ (0x3F << 3)))
}

func (m *Move) SetTo(square Square) {
	*m = Move((uint32(square) & 0x3F << 9) | (uint32(*m) &^ (0x3F << 9)))
}

func (m *Move) SetCapture(piece Piece) {
	*m = Move((uint32(piece) & 0x7 << 15) | (uint32(*m) &^ (0x7 << 15)))
}

func (m *Move) SetPromotion(piece Piece) {
	*m = Move((uint32(piece) & 0x7 << 18) | (uint32(*m) &^ (0x7 << 18)))
}

func (m *Move) SetFlag(flag MoveFlag) {
	*m = Move((uint32(flag) & 0x7 << 21) | (uint32(*m) &^ (0x7 << 21)))
}

func (m *Move) SetScore(score uint8) {
	*m = Move((uint32(score) & 0xff << 24) | (uint32(*m) &^ (0xff << 24)))
}

func (m Move) MoveToString() string {
	from := int(m.From())
	to := int(m.To())
	fromFile := string('a' + byte(from%8))
	fromRank := string('1' + byte(from/8))
	toFile := string('a' + byte(to%8))
	toRank := string('1' + byte(to/8))

	moveStr := fromFile + fromRank + toFile + toRank

	if m.Promotion() != None {
		promo := ""
		switch m.Promotion() {
		case Queen:
			promo = "q"
		case Rook:
			promo = "r"
		case Bishop:
			promo = "b"
		case Knight:
			promo = "n"
		}
		moveStr += promo
	}

	return moveStr
}
