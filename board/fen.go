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
		s.ParseCastlingRights(castlingRights)
	}
	if parts[3] != "-" {
		enpassantSquare := StringToSquare(parts[3])
		if enpassantSquare != 64 {
			s.Enpassant = enpassantSquare
		}
	}
	s.friends = s.GetAll(s.WhiteToMove)
	s.enemies = s.GetAll(!s.WhiteToMove)
	s.occupied = s.GetOccupied()
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
	fmt.Printf("   a b c d e f g h\n")
	for rank := Square(7); rank < 8; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := Square(0); file < 8; file++ {
			var square Square = rank*8 + file
			piece := s.GetPiece(square, White)
			if piece == None {
				piece = s.GetPiece(square, Black)
				if piece != None {
					piece += 6
				}
			}
			fmt.Printf("%s ", PieceToChar(piece))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
func PieceToChar(piece Piece) string {
	pieceChars := ".PNBRQKpnbrqk"
	return string(pieceChars[piece])
}

func (s *Board) ParseCastlingRights(rights string) {
	for _, char := range rights {
		switch char {
		case 'K':
			s.UpdateCastlingRights(CastlingWhiteShort)
		case 'Q':
			s.UpdateCastlingRights(CastlingWhiteLong)
		case 'k':
			s.UpdateCastlingRights(CastlingBlackShort)
		case 'q':
			s.UpdateCastlingRights(CastlingBlackLong)
		}
	}
}
func StringToSquare(s string) Square {
	if len(s) != 2 {
		return 64 // Indicar que la cadena no es válida
	}

	file := s[0]
	rank := s[1]

	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return 64 // Indicar que la cadena no es válida
	}

	file -= 'a'
	rank -= '1'

	return Square(rank*8 + file)
}
