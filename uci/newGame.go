package uci

import "nicarao/board"

func newGame() {
	boardUCI = board.Board{}
	readyOK()
}
