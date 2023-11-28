package uci

import (
	"bufio"
	"fmt"
	"nicarao/board"
	"os"
	"strings"
)

var boardUCI board.Board

// StartUCI inicia la comunicación con el GUI a través de UCI
func StartUCI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch tokens[0] {
		case "uci":
			fmt.Println("id name Nicarao")
			fmt.Println("id author Deybis Antonio Melendez Vargas")
			fmt.Println("uciok")
		case "isready":
			readyOK()
		case "ucinewgame":
			newGame()
		case "position":
			handlePosition(tokens)
		case "go":
			HandleGo(tokens)
		case "stop":
			// Lógica para detener la búsqueda
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command:", text)
		}
	}
}

func readyOK() {
	fmt.Println("readyok")
}
