package main

import (
	"fmt"
	"nicarao/board"
	"nicarao/search"
)

func main() {
	var test board.Board
	test.ParseFEN(board.StartingPos)
	for i := uint8(1); i < 8; i++ {
		var score int16 = search.PVSearch(&test, -search.Inf, search.Inf, i)
		fmt.Println(i, score)
	}
}
