package main

import (
	"nicarao/board"
)

func main() {
	var test *board.Board = board.NewBoard()
	test.WhiteToMove = true
	/*test.Bitboards[true][board.Pawn] = board.SetBit(0, board.E2)
	test.Bitboards[false][board.King] = board.SetBit(0, board.E8)
	test.Bitboards[false][board.Pawn] = board.SetBit(0, board.D3)
	test.Bitboards[true][board.King] = board.SetBit(0, board.G6)
	test.Bitboards[true][board.Queen] = board.SetBit(0, board.A1)*/
	test.ParseFEN(board.StartingPos)
	test.Print()
}
