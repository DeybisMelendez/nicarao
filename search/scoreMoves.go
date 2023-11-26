package search

import (
	"math/bits"
	"nicarao/board"
)

/*
ORDEN DE LOS MOVIMIENTOS 0-255 (uint8):
1. 255 PV MOVE / HASH MOVE
2. 254 - 18 = 236 Capturas/Promociones* ganadoras (pieza capturada + Promoción - pieza recapturada = saldo positivo)
3. 236 Capturas/Promociones** igualadas (pieza capturada + Promoción - pieza recapturada = 0)
4. 235 Movimiento asesino (Killer move) #1
5. 234 Movimiento asesino (Killer move) #2
6. 233 Counter Move
7. 232 - 223 = 9 Movimientos tranquilos ordenados por History Moves o ¿¡Piece Square Table!?
8. 8 - 8 = 0 Capturas/Promociones perdedoras** (pieza capturada + Promoción - pieza recapturada = saldo negativo)

* El maximo valor (positivo) conseguible es Queen(9) + Queen(9) - None(0) = 18
** El mínimo valor (negativo) conseguible sería Pawn(1) + None(0) - Queen(9) = -8
*/

func scoreMoves(b *board.Board, moves *board.MoveList, oldBestMove board.Move) {
	for i := 0; i < int(moves.Index); i++ {
		var move board.Move = moves.List[i]
		var flag board.MoveFlag = move.Flag()
		if oldBestMove == move { // PV Move / Hash Move
			moves.List[i].SetScore(255)
		} else if flag == board.Capture || flag == board.EnpassantCapture || flag == board.CapturePromotion ||
			flag == board.Promotion { // Capturas
			var captureValue int8 = int8(pieceCaptureValue[move.Capture()]) - int8(recapturedValue(b, move))
			if flag == board.CapturePromotion || flag == board.Promotion {
				captureValue += int8(pieceCaptureValue[move.Promotion()]) - 1 //Se resta 1 por el peón
			}
			if captureValue >= 0 {
				moves.List[i].SetScore(236 + uint8(captureValue))
			} else {
				moves.List[i].SetScore(8 + uint8(captureValue)) // Se suma porque el captureValue es negativo
			}
		} else if isKillerMove(b.Ply, move) > 0 { //Killer moves
			moves.List[i].SetScore(233 + isKillerMove(b.Ply, move))
		} else if isCounterMove(b.GetEnemyColor(), move) { //Como se hizo make, el color de la pieza sería el del enemigo
			moves.List[i].SetScore(233)
		} else { //History moves
			moves.List[i].SetScore(9 + getHistoryMove(b.GetEnemyColor(), move))
		}
	}
}

func recapturedValue(b *board.Board, move board.Move) uint8 {
	b.MakeMove(move)
	var captures board.MoveList
	var captureMove board.Move
	var kingSquare board.Square
	var color uint8 = b.WhiteToMove // La legalidad debe ser probada en el turno enemigo
	var valueCapture uint8
	b.GeneratePseudoCaptureSquare(&captures, move.To())
	for i := 0; i < int(captures.Index); i++ {
		captureMove = captures.List[i]
		b.MakeMove(captureMove)
		kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
		if !b.IsUnderAttack(kingSquare, color) {
			b.UnMakeMove(captureMove)
			valueCapture = pieceCaptureValue[move.Piece()]
			break
		} else {
			b.UnMakeMove(captureMove)
		}
	}
	b.UnMakeMove(move)
	return valueCapture
}
