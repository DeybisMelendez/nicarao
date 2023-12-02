package search

import (
	"fmt"
	"math"
	"nicarao/board"
	"time"
)

var Stopped bool
var StopTime int64 = -1

func SearchWithStopTime(b *board.Board, stopTime int64) {
	var start int64 = time.Now().UnixMilli()
	var bestmove string
	var flag uint8
	var scoreType string = "cp"
	restartSearch()
	StopTime = stopTime

	for depth := uint8(1); depth < math.MaxUint8; depth++ {
		var score int16 = PVSearch(b, -Inf, Inf, depth)
		if isTimeToStop(b.Nodes) {
			break
		}
		var time_elapsed int64 = time.Now().UnixMilli() - start
		bestmove = GetBestMove(b.Hash).MoveToString()
		flag = GetFlagScore(b.Hash)
		if flag == TTLowerBound {
			scoreType = "lowerbound"
		} else if flag == TTUpperBound {
			scoreType = "upperbound"
		}
		fmt.Printf("info depth %d score %s %d nodes %d time %d pv %s\n", depth, scoreType, score, b.Nodes, time_elapsed, bestmove)
	}
	fmt.Println("bestmove ", bestmove)
}

func SearchWithDepth(b *board.Board, depth uint8) {
	//TODO agregar busqueda con depth
}

func restartSearch() {
	historyMoves = [2][7][64]uint8{}
	counterMoves = [2][7][64]board.Move{}
	Stopped = false
	StopTime = -1
}

func isTimeToStop(nodes int) bool {
	if Stopped {
		return Stopped
	}
	if StopTime != -1 {
		if nodes&2047 == 0 {
			if time.Now().UnixMilli() >= StopTime {
				Stopped = true
			}
		}
	}
	return Stopped
}
