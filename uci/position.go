package uci

import (
	"fmt"
	"nicarao/board"
	"strings"
)

func handlePosition(tokens []string) {
	if len(tokens) < 2 {
		fmt.Println("Error: Missing parameters for position command")
		return
	}

	switch tokens[1] {
	case "startpos":
		boardUCI.ParseFEN(board.StartingPos)
	case "fen":
		if len(tokens) < 8 {
			fmt.Println("Error: Invalid FEN string for position command")
			return
		}
		fen := strings.Join(tokens[2:], " ")
		boardUCI.ParseFEN(fen)
	default:
		fmt.Println("Error: Invalid parameters for position command")
	}
}
