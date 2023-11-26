package main

import (
	"fmt"
	"nicarao/board"
	"nicarao/search"
)

func main() {
	var test board.Board
	test.ParseFEN("8/1p3rk1/p6p/6p1/2P3P1/2bQ3K/P2p2P1/3Rq3 w - - 3 36")
	for i := uint8(1); i < 16; i++ {
		var score int16 = search.PVSearch(&test, -search.Inf, search.Inf, i)
		fmt.Println(i, score)
	}
}
