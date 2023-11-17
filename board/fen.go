package board

import (
	"fmt"
	"strings"
)

const StartingPos string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func (s *Board) ParseFEN(fen string) error {
	parts := strings.Fields(fen)
	if len(parts) < 4 {
		return fmt.Errorf("error: FEN string is invalid, %s", fen)
	}
	var rank Square = 7
	var file Square = 0
	if parts[1] == "w" {
		s.WhiteToMove = true
	} else if parts[1] == "b" {
		s.WhiteToMove = false
	} else {
		return fmt.Errorf("error: FEN string is invalid, %s", fen)
	}

	for _, char := range parts[0] {
		switch char {
		case '/':
			rank--
			file = 0
		case '1', '2', '3', '4', '5', '6', '7', '8':
			file += Square(char - '0')
		default:
			piece := CharToPiece(char)
			s.Bitboards[piece <= 7][piece%7] |= SetBit(0, rank*8+file)
			file++
		}
	}
	if len(parts) >= 5 {
		castlingRights := parts[2]
		s.ParseCastling(castlingRights)
	}
	if parts[3] != "-" {
		enpassantSquare, _ := StringToSquare(parts[3])
		if enpassantSquare != 64 {
			s.Enpassant = enpassantSquare
		}
	}
	s.friends = s.GetAll(s.WhiteToMove)
	s.enemies = s.GetAll(!s.WhiteToMove)
	s.occupied = s.friends | s.enemies
	return nil
}

func CharToPiece(char rune) Piece {
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

func (s *Board) Print() {
	fmt.Printf("  a b c d e f g h\n")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
		for file := 0; file < 8; file++ {
			square := rank*8 + file
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
func PieceToChar(piece Piece) string {
	pieceChars := ".PNBRQK"
	return string(pieceChars[piece])
}

func (s *Board) ParseCastling(rights string) {
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
