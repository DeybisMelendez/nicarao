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
	var move board.Move
	switch tokens[1] {
	case "startpos":
		boardUCI = board.Board{}
		boardUCI.ParseFEN(board.StartingPos)
	case "fen":
		if len(tokens) < 8 {
			fmt.Println("Error: Invalid FEN string for position command")
			return
		}
		fen := strings.Join(tokens[2:], " ")
		boardUCI = board.Board{}
		boardUCI.ParseFEN(fen)
	default:
		fmt.Println("Error: Invalid parameters for position command")
	}
	if len(tokens) > 3 && tokens[2] == "moves" {
		movesList := tokens[3:]
		for _, moveStr := range movesList {
			move = board.ParseMove(&boardUCI, moveStr)
			boardUCI.MakeMove(move)
		}
	}
}
