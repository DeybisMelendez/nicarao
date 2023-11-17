package board

//Board contiene todos los elementos necesarios para representar un tablero de ajedrez
type Board struct {
	//WhiteToMove indica el turno del jugador, true para blancas y false para negras
	WhiteToMove bool
	//Bitboards es un mapa que contiene un bitboard por cada pieza de ajedrez en el tablero
	Bitboards map[bool]map[Piece]uint64
	//Enpassant indica la casilla donde la captura al paso puede suceder,
	//su estado es 0 cuando no existe captura al paso
	Enpassant Square
	//Castling indica el derecho a enrocar de cada bando,
	//guarda 1 byte por cada derecho de enroque,
	//largo, corto, para blancas o negras
	Castling uint8
	//HalfmoveClock es un contador de movimientos que se resetea cuando
	//se realiza un movimiento de peón o captura,
	//sirve para aplicar la regla de los 50 movimientos
	HalfmoveClock uint8
	//friends es un bitboard útil para calcular movimientos posibles,
	//se debe actualizar al realizar un movimiento nuevo
	friends uint64
	//enemies es un bitboard útil para calcular movimientos posibles,
	//se debe actualizar al realizar un movimiento nuevo
	enemies uint64
	//occupied es un bitboard útil para calcular movimientos posibles,
	//se debe actualizar al realizar un movimiento nuevo
	occupied uint64
}

func NewBoard() *Board {
	return &Board{
		Bitboards: map[bool]map[Piece]uint64{
			White: {
				Pawn:   0,
				Knight: 0,
				Bishop: 0,
				Rook:   0,
				Queen:  0,
				King:   0,
			},
			Black: {
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
func (s *Board) CloneBoard() *Board {
	clonedBoard := &Board{
		WhiteToMove:   s.WhiteToMove,
		Enpassant:     s.Enpassant,
		Castling:      s.Castling,
		HalfmoveClock: s.HalfmoveClock,
	}

	// Copiar bitboards
	clonedBoard.Bitboards = make(map[bool]map[Piece]uint64)
	for color, pieces := range s.Bitboards {
		clonedBoard.Bitboards[color] = make(map[Piece]uint64)
		for piece, bb := range pieces {
			clonedBoard.Bitboards[color][piece] = bb
		}
	}

	return clonedBoard
}

//GetAll devuelve un bitboard con todas las casillas ocupadas del jugador indicado
func (s *Board) GetAll(color bool) uint64 {
	var total uint64
	for _, piece := range pieceTypes {
		total |= s.Bitboards[color][piece]
	}
	return total
}

//GetPiece devuelve la pieza que está controlando la casilla indicada
func (s *Board) GetPiece(square Square, color bool) Piece {
	var mask uint64 = SetBit(0, square)

	for _, piece := range pieceTypes {
		if s.Bitboards[color][piece]&mask != 0 {
			return piece
		}
	}
	return None
}

//CanCastle devuelve true si el jugador indicado tiene derecho a enrocar corto
func (s *Board) CanCastle(color bool, isShort bool) bool {
	return castling[color][isShort]&s.Castling != 0
}

//UpdateCastling cambia el derecho a enroque al nuevo estado
func (s *Board) UpdateCastling(Castling uint8) {
	s.Castling |= Castling
}

//RemoveCastling elimina un derecho de enroque según se indique
func (s *Board) RemoveCastling(Castling uint8) {
	s.Castling &= ^Castling
}

//HandleCastle controla la modificación del derecho a enroque
func (s *Board) HandleCastle(color bool, isShort bool, remove bool) {
	if remove {
		s.RemoveCastling(castling[color][isShort])
	} else {
		s.UpdateCastling(castling[color][isShort])
	}
}
