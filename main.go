package main

import (
	"fmt"
	"nicarao/board"
	"nicarao/search"
)

func main() {
	var test board.Board
	test.ParseFEN("1n2R3/2N5/3bBp2/1Kp1p3/1p1Pkpp1/3N4/3PPQ1n/3r4 w - - 0 1")
	for i := uint8(1); i < 16; i++ {
		var score int16 = search.PVSearch(&test, -search.Inf, search.Inf, i)
		fmt.Println(i, score)
	}
}
