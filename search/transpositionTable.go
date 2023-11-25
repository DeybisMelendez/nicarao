package search

import "nicarao/board"

//http://web.archive.org/web/20070809015843/www.seanet.com/%7Ebrucemo/topics/hashing.htm
func recordHash(hash uint64, value int16, depth uint8, flag uint8, bestmove board.Move, age uint8) {
	ttEntry := &transpositionTable[hash%ttSize]
	if depth > ttEntry.Depth && age >= ttEntry.Age {
		ttEntry.Depth = depth
		ttEntry.Flag = flag
		ttEntry.Value = value
		ttEntry.BestMove = bestmove
		ttEntry.Hash = hash
		ttEntry.Age = age
	}
}

func probeHash(hash uint64, alpha int16, beta int16, depth uint8, move *board.Move, age uint8) int16 {
	var ttEntry TranspositionTable = transpositionTable[hash%ttSize]
	if ttEntry.Hash == hash && ttEntry.Age >= age {
		if ttEntry.Depth >= depth {
			if ttEntry.Flag == TTExact {
				return ttEntry.Value
			}
			if ttEntry.Flag == TTUpperBound && ttEntry.Value <= alpha {
				return alpha
			}
			if ttEntry.Flag == TTLowerBound && ttEntry.Value >= beta {
				return beta
			}
		}
		*move = ttEntry.BestMove
	}
	return NoHashEntry
}

/*func getPV(fen string, depth int8, bestmove board.Move) string {
	var pv string
	var board board.Board = chess.ParseFen(fen)

	for i := int8(0); i < depth; i++ {
		hash := board.Hash()
		entry := getEntry(hash)
		if entry.Hash == hash && entry.BestMove != 0 {
			if i == 0 {
				*bestmove = entry.BestMove
			}
			pv += entry.BestMove.String() + " "
			board.Apply(entry.BestMove)
		} else {
			break
		}
	}
	return pv
}

func getEntry(hash uint64) TranspositionTable {
	return transpositionTable[hash%ttSize]
}
*/
