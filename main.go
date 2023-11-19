package main

import (
	"fmt"
	"nicarao/board"
)

func main() {
	//board.Init()
	var test board.Board = *board.NewBoard()
	test.ParseFEN("4k3/8/8/b7/1P6/8/8/4K3 w - - 0 1")
	test.Print()
	board.PrintBitboard(test.GetAll(test.WhiteToMove))
	board.PrintBitboard(test.GetAll(!test.WhiteToMove))
	moves := test.GeneratePseudoMoves()
	for _, move := range moves {
		var color bool = test.WhiteToMove
		var unMove = test.MakeMove(&move)
		var kingBB uint64 = test.Bitboards[color][board.King]
		if !test.IsUnderAttack(kingBB, test.WhiteToMove) {
			fmt.Println(move)
		}
		test.UnMakeMove(unMove)
	}
	board.PrintBitboard(test.GetAll(test.WhiteToMove))
	board.PrintBitboard(test.GetAll(!test.WhiteToMove))
	//FIXME:El rey se mueve y desaparece las piezas enemigas
}
