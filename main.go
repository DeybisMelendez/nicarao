package main

import (
	"fmt"
	"nicarao/board"
)

func main() {
	var test board.Board = *board.NewBoard()
	test.ParseFEN("k6K/rpbpnprp/P1P1P1P1/nppprpbn/1P1P1P1P/rqppqpbp/P1P1P1P1/8 w - - 0 1")
	moves := test.GeneratePseudoMoves()
	for _, move := range moves {
		//if test.IsMoveLegal(&move) {
		board.PrintBitboard(test.GenerateAttacksForPiece(board.Pawn, move.From, test.WhiteToMove))
		fmt.Println(move.From, move.To, move.Capture)
		fmt.Println("--------------------")
		//}
	}
}
