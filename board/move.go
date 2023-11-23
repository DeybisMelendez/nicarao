package board

type MoveFlag uint8
type Move uint32

func NewMove(piece Piece, from Square, to Square, capture Piece, promo Piece, flag MoveFlag) Move {
	var move Move = Move((uint32(piece) & 0x7) | // 3 bits Representa piezas 0-7
		((uint32(from) & 0x3f) << 3) | // 6 bits Representa casillas 0-63
		((uint32(to) & 0x3f) << 9) | // 6 bits Representa casillas 0-63
		((uint32(capture) & 0x7) << 15) | // 3 bits Representa piezas 0-7
		((uint32(promo) & 0x7) << 18) | // 3 bits Representa piezas 0-7
		((uint32(flag) & 0x7) << 21)) // 3 bits Representa flags 0-7
	return move
}

func (m *Move) Piece() Piece {
	return Piece(*m & 0x7)
}

func (m *Move) From() Square {
	var sq = Square((*m >> 3) & 0x3F)
	if sq > 63 {
		panic(sq)
	}
	return sq
}

func (m *Move) To() Square {
	var sq = Square((*m >> 9) & 0x3F)
	if sq > 63 {
		panic(sq)
	}
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
