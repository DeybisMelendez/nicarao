package board

func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}

	var totalNodes int64
	moves := b.GeneratePseudoMoves()
	for _, move := range moves {
		var color bool = b.WhiteToMove
		b.MakeMove(&move)
		var kingBB uint64 = b.Bitboards[color][King]
		if !b.IsUnderAttack(kingBB, b.WhiteToMove) {
			totalNodes += Perft(b, depth-1)
		}
		b.UnMakeMove(&move)
	}
	return totalNodes
}

/*func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}
	var totalNodes int64
	moves := b.GeneratePseudoMoves()
	for _, move := range moves {
		if b.IsMoveLegal(&move) {
			b.MakeMove(&move)
			totalNodes += Perft(b, depth-1)
			b.UnMakeMove(&move)
		}
	}

	return totalNodes
}*/
