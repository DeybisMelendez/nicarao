package board

func Perft(b *Board, depth int) int64 {
	if depth == 0 {
		return 1
	}

	var totalNodes int64
	moves := b.GeneratePseudoMoves()
	//var captures, ep, castles, promotions, checks int
	for _, move := range moves {
		var color bool = b.WhiteToMove
		var unMove = b.MakeMove(&move)
		var kingBB uint64 = b.Bitboards[color][King]

		if !b.IsUnderAttack(kingBB, b.WhiteToMove) {
			totalNodes += Perft(b, depth-1)
			/*if move.Flag == Capture || move.Flag == CapturePromotion {
				captures++
			}
			if move.Flag == Promotion || move.Flag == CapturePromotion {
				promotions++
			}
			if move.Flag == KingCastle || move.Flag == QueenCastle {
				castles++
			}
			if b.IsUnderAttack(b.Bitboards[b.WhiteToMove][King], b.WhiteToMove) {
				checks++
			}
			if move.Flag == EnpassantCapture {
				ep++
			}*/
		}
		b.UnMakeMove(unMove)
	}
	//fmt.Printf("Total Capturas: %d, Total Coronaciones: %d, Total Enroques: %d,\nTotal jaques: %d, Total Al Paso: %d\n\n", captures, promotions, castles, checks, ep)
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
