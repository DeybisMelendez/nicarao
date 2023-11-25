package board

import "fmt"

//var unMakeInfoList [MaxPly]UnMakeInfo
//var unMakeInfoIndex int

//Board contiene todos los elementos necesarios para representar un tablero de ajedrez
type Board struct {
	Hash uint64
	//WhiteToMove indica el turno del jugador, true para blancas y false para negras
	WhiteToMove uint8
	//Bitboards es un Array que contiene un bitboard por cada pieza de ajedrez en el tablero
	Bitboards [2][7]uint64
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
	//Historial que contiene información irreversible como Castling y Enpassant
	unMakeInfoList [MaxPly]UnMakeInfo
	//Index del último cambio realizado en el historial UnMakeInfoList
	unMakeInfoIndex int

	//Indica la profundidad alcanzada por el tablero
	Ply int16
}

func NewBoard() *Board {
	return &Board{
		/*Bitboards: map[bool]map[Piece]uint64{
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
		},*/
	}
}

func (s *Board) Print() {
	var maxRank int = 8
	fmt.Printf("  a b c d e f g h\n")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
		for file := 0; file < maxRank; file++ {
			square := rank*maxRank + file
			piece := s.GetPiece(Square(square), White)
			if piece == None {
				piece = s.GetPiece(Square(square), Black)
				fmt.Printf("%s ", pieceToEmoji(piece))
			} else {
				fmt.Printf("%s ", pieceToEmoji(piece+6))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
