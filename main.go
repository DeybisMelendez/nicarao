package main

import (
	"fmt"
	"nicarao/board"
	"nicarao/search"
)

func main() {
	var test board.Board
	test.ParseFEN("8/1P5B/8/2P5/8/6K1/NkP3p1/RN6 w - - 0 1")
	for i := uint8(1); i < 16; i++ {
		var score int16 = search.PVSearch(&test, -search.Inf, search.Inf, i)
		fmt.Println(i, score)
	}
}
