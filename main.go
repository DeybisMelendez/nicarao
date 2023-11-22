package main

import (
	"fmt"
	"nicarao/board"
)

func main() {
	for i := board.Square(0); i < 64; i++ {
		fmt.Printf("0x%x, ", board.SetBit(0, i))
	}
}
