package board

func Perft(b *Board, depth int) int64 {
	if depth <= 0 {
		return 1
	}
	if depth == 1 {
		return int64(len(b.GenerateMoves()))
	}

	var totalNodes int64 = 0
	moves := b.GenerateMoves()
	for _, move := range moves {
		b.MakeMove(&move)
		totalNodes += Perft(b, depth-1)
		b.UnMakeMove(&move)
	}

	return totalNodes
}
