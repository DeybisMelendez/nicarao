package search

import (
	"math"
	"nicarao/board"
)

var pieceCaptureValue [7]uint8 = [7]uint8{0, 1, 3, 3, 5, 9, 0}
var transpositionTable [ttSize]TranspositionTable

type TranspositionTable struct {
	Hash     uint64
	Age      uint8
	Depth    uint8
	Flag     uint8
	Value    int16
	BestMove board.Move
}

const Inf int16 = math.MaxInt16
const MateValue int16 = 32700
const (
	TTExact      uint8 = iota //PV Move https://www.chessprogramming.org/Exact_Score
	TTLowerBound              //Cut-Node  https://www.chessprogramming.org/Lower_Bound
	TTUpperBound              // Upper-Node https://www.chessprogramming.org/Upper_Bound
)
const MB = 1024 * 1024 // Tamaño en bytes de 1 MB
// 16 bytes suman los bytes contenidos teóricamente por el struct TranspositionTable
const ttSize uint64 = 64 * MB / 18
const NoHashEntry = math.MinInt16
