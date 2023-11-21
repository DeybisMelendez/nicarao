package board

func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}

	var totalNodes int64
	moves := b.GeneratePseudoMoves()
	for _, move := range moves {
		var color bool = b.WhiteToMove
		var unMakeInfo = b.MakeMove(&move)
		var kingBB uint64 = b.Bitboards[color][King]

		if !b.IsUnderAttack(kingBB, !color) {
			totalNodes += Perft(b, depth-1)
		}
		b.UnMakeMove(&move, &unMakeInfo)
	}
	return totalNodes
}
