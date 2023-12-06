package main

import "nicarao/uci"

func main() {
	uci.StartUCI()
	/*var mask uint64 = 0
	var fileNumber = 4
	var rankNumber = 4
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			var square uint64 = uint64(rank)*8 + uint64(file)
			if fileNumber != -1 {
				if file == fileNumber {
					mask |= 1 << square
				}
			} else if rankNumber != -1 {
				if rank == rankNumber {
					mask |= 1 << square
				}
			}
		}
	}
	board.PrintBitboard(mask)*/
}
