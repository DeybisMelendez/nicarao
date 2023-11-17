package main

import (
	"fmt"
	"nicarao/board"
)

func main() {
	var test board.Board = *board.NewBoard()
	test.ParseFEN("k7/8/8/8/8/3ppp2/2p3p1/4K3 w - - 0 1")
	test.Print()
	//fmt.Println(len(test.GeneratePseudoMoves()))
	fmt.Println(board.Perft(&test, 1))
	//board.PrintBitboard(board.PawnBlackAttacksMasks[board.C6])
	//board.PrintBitboard(board.PawnBlackAttacksMasks[board.H3])
}
