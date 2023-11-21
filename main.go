package main

import (
	"fmt"
	"nicarao/board"
	"nicarao/utils"
)

func main() {
	for i := board.Square(0); i < 64; i++ {
		bitboard := utils.BitscanReverse(board.Rays["southWest"][i])
		/*if bitboard == 64 {
			bitboard = 0
		}*/
		fmt.Printf("%d, ", bitboard)
	}
}
