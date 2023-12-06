package board

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *Board) ParseFEN(fen string) error {
	parts := strings.Fields(fen)
	if len(parts) < 4 {
		return fmt.Errorf("error: FEN string is invalid, %s", fen)
	}
	var rank Square = 7
	var file Square = 0
	if parts[1] == "w" {
		s.WhiteToMove = White
	} else if parts[1] == "b" {
		s.WhiteToMove = Black
	} else {
		return fmt.Errorf("error: FEN string is invalid, %s", fen)
	}
	if len(parts) >= 5 {
		halfmoveCount, err := strconv.ParseUint(parts[4], 10, 8)
		if err != nil {
			return fmt.Errorf("error parsing HalfmoveClock from FEN: %s", err)
		}
		s.HalfmoveClock = uint8(halfmoveCount)
	}

	for _, char := range parts[0] {
		switch char {
		case '/':
			rank--
			file = 0
		case '1', '2', '3', '4', '5', '6', '7', '8':
			file += Square(char - '0')
		default:
			piece := charToPiece(char)
			if piece <= 7 {
				s.Bitboards[White][piece%7] |= SetBit(0, rank*8+file)
			} else {
				s.Bitboards[Black][piece%7] |= SetBit(0, rank*8+file)
			}
			file++
		}
	}
	if parts[2] != "-" {
		castlingRights := parts[2]
		s.parseCastling(castlingRights)
	}
	if parts[3] != "-" {
		enpassantSquare, _ := StringToSquare(parts[3])
		if enpassantSquare != 64 {
			s.Enpassant = enpassantSquare
		}
	}
	s.friends = s.GetAll(s.WhiteToMove)
	s.enemies = s.GetAll(s.GetEnemyColor())
	s.occupied = s.friends | s.enemies
	s.calcZobristHash()
	s.Nodes = 0
	return nil
}

func charToPiece(char rune) Piece {
	switch char {
	case 'P':
		return Pawn
	case 'N':
		return Knight
	case 'B':
		return Bishop
	case 'R':
		return Rook
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'p':
		return Pawn + 7
	case 'n':
		return Knight + 7
	case 'b':
		return Bishop + 7
	case 'r':
		return Rook + 7
	case 'q':
		return Queen + 7
	case 'k':
		return King + 7
	default:
		return None
	}
}

func (s *Board) parseCastling(rights string) {
	for _, char := range rights {
		switch char {
		case 'K':
			s.UpdateCastling(castling[White][CastleShort])
		case 'Q':
			s.UpdateCastling(castling[White][CastleLong])
		case 'k':
			s.UpdateCastling(castling[Black][CastleShort])
		case 'q':
			s.UpdateCastling(castling[Black][CastleLong])
		}
	}
}

func pieceToEmoji(piece Piece) string {
	switch piece {
	case Pawn + 6:
		return "\u265F"
	case Knight + 6:
		return "\u265E"
	case Bishop + 6:
		return "\u265D"
	case Rook + 6:
		return "\u265C"
	case Queen + 6:
		return "\u265B"
	case King + 6:
		return "\u265A"
	case Pawn:
		return "\u2659"
	case Knight:
		return "\u2658"
	case Bishop:
		return "\u2657"
	case Rook:
		return "\u2656"
	case Queen:
		return "\u2655"
	case King:
		return "\u2654"
	case None:
		return "."
	default:
		return "?" // Espacio para casillas vacías
	}
}

func StringToSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("string is invalid: %s", s)
	}

	file := s[0]
	rank := s[1]

	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return 0, fmt.Errorf("string is invalid: %s", s)
	}

	file -= 'a'
	rank -= '1'

	return Square(rank*8 + file), nil
}
func ParseMove(b *Board, movestr string) Move {
	var move Move
	var promo Piece
	var moves MoveList
	/*if len(movestr) < 4 || len(movestr) > 5 {
		return mv, errors.New("Invalid move to parse.")
	}*/
	from, _ := StringToSquare(movestr[0:2])
	to, _ := StringToSquare(movestr[2:4])

	/*if errf != nil || errto != nil {
		return move //, errors.New("Invalid move to parse.")
	}*/
	if len(movestr) == 5 {
		switch movestr[4] {
		case 'n':
			promo = Knight
		case 'b':
			promo = Bishop
		case 'q':
			promo = Queen
		case 'r':
			promo = Rook
			/*default:
			return move, errors.New("Invalid promotion symbol in move.")*/
		}
	}
	b.GeneratePseudoMoves(&moves)
	for i := 0; i < int(moves.Index); i++ {
		if moves.List[i].From() == from && moves.List[i].To() == to && moves.List[i].Promotion() == promo {
			move = moves.List[i]
			break
		}
	}
	if move == 0 {
		panic("error: movimiento no encontrado válido")
	}
	return move //, nil
}
