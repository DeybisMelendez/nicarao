package main

import (
	"nicarao/board"
)

func main() {
	var test board.Board = *board.NewBoard()
	test.ParseFEN(board.StartingPos)
	test.Print()
}
