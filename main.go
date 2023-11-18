package main

import (
	"fmt"
	"nicarao/board"
)

func main() {
	//board.Init()
	var test board.Board = *board.NewBoard()
	test.ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - ")
	test.Print()
	fmt.Println(test.Castling)
	fmt.Println(test.CanCastle(board.White, false))
	moves := test.GeneratePseudoMoves()
	for _, move := range moves {
		var color bool = test.WhiteToMove
		var unMove = test.MakeMove(&move)
		var kingBB uint64 = test.Bitboards[color][board.King]
		if !test.IsUnderAttack(kingBB, test.WhiteToMove) {
			fmt.Println(move)
		}
		test.UnMakeMove(unMove)
	}

}
