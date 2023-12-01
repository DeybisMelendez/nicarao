package search

import (
	"math/bits"
	"nicarao/board"
)

/*
ORDEN DE LOS MOVIMIENTOS 0-255 (uint8):
1. 255 PV MOVE / HASH MOVE
2. 254 - PromoValue (min 1 - max 4) = 250 Promociones
3. 249 Capturas ganadoras
4. 248 Capturas igualadas
5. 247 Movimiento asesino (Killer move) #1
6. 246 Movimiento asesino (Killer move) #2
7. 245 Counter Move
8. 200->3 Movimientos tranquilos ordenados por History Moves o ¿¡Piece Square Table!?
9. 2 jugadas tranquilas sin evaluar pero posiblemente legales
10. 1 Capturas perdedoras
11. 0 jugadas evaluadas ilegales (capturas)

*/

func scoreMoves(b *board.Board, moves *board.MoveList, oldBestMove board.Move) {
	for i := 0; i < int(moves.Index); i++ {
		var move board.Move = moves.List[i]
		var flag board.MoveFlag = move.Flag()
		if oldBestMove == move { // PV Move / Hash Move
			moves.List[i].SetScore(255)
		} else if flag == board.Promotion || flag == board.CapturePromotion {
			moves.List[i].SetScore(255 - promotionValue[move.Promotion()])
		} else if flag == board.Capture || flag == board.EnpassantCapture { // Capturas
			captureValue, isCaptureLegal := seeCapture(b, move)
			if isCaptureLegal {
				if captureValue > 0 { // Capturas ganadoras
					moves.List[i].SetScore(249)
				} else if captureValue == 0 { // Capturas igualadas
					moves.List[i].SetScore(248)
				} else { // Capturas perdedoras
					moves.List[i].SetScore(1)
				}
			} else {
				moves.List[i].SetScore(0) // Capturas ilegales las ordenamos al final
			}
		} else if isKillerMove(b.Ply, move) > 0 { //Killer moves
			moves.List[i].SetScore(245 + isKillerMove(b.Ply, move))
		} else if isCounterMove(b.WhiteToMove, move) {
			moves.List[i].SetScore(245)
		} else { //History moves
			moves.List[i].SetScore(2 + getHistoryMove(b.WhiteToMove, move))
		}
	}
}

func scoreCaptures(moves *board.MoveList) {
	for i := 0; i < int(moves.Index); i++ {
		moves.List[i].SetScore(getMVV_LVA(moves.List[i]))
	}
}

/*

int see(int square, int side)
{
   value = 0;
   piece = get_smallest_attacker(square, side);
   // skip if the square isn't attacked anymore by this side
   if ( piece )
   {
      make_capture(piece, square);
      // Do not consider captures if they lose material, therefor max zero
      value = max (0, piece_just_captured() -see(square, other(side)) );
      undo_capture(piece, square);
   }
   return value;
}

*/
func see(b *board.Board, square board.Square) int8 {
	var value int8 = 0
	var captures board.MoveList
	var smallestAttacker board.Move
	var captureMove board.Move
	var kingSquare board.Square
	var color uint8 = b.WhiteToMove // La legalidad debe ser probada en el turno enemigo
	var firstCapture bool
	b.GeneratePseudoCaptureSquare(&captures, square)
	//get smallest attacker
	for i := 0; i < int(captures.Index); i++ {
		captureMove = captures.List[i]
		if !firstCapture {
			b.MakeMove(captureMove)
			//Evaluar capturas legales
			kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
			if !b.IsUnderAttack(kingSquare, color) {
				smallestAttacker = captureMove
				firstCapture = true
			}
			b.UnMakeMove(captureMove)
		} else if pieceCaptureValue[smallestAttacker.Piece()] > pieceCaptureValue[captureMove.Piece()] {
			b.MakeMove(smallestAttacker)
			//Evaluar capturas legales
			kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
			if !b.IsUnderAttack(kingSquare, color) {
				smallestAttacker = captureMove
			}
			b.UnMakeMove(smallestAttacker)
		}
	}
	// Si no hay mas capturas
	if smallestAttacker.Piece() != board.None {
		b.MakeMove(smallestAttacker)
		value = pieceCaptureValue[smallestAttacker.Capture()] - see(b, square)
		/*if value < 0 {
			value = 0
		}*/
		b.UnMakeMove(smallestAttacker)
	}
	return value
}

func seeCapture(b *board.Board, move board.Move) (int8, bool) {
	var value int8
	var kingSquare board.Square
	var color uint8 = b.WhiteToMove // La legalidad debe ser probada en el turno enemigo
	b.MakeMove(move)
	kingSquare = board.Square(bits.TrailingZeros64(b.Bitboards[color][board.King]))
	if !b.IsUnderAttack(kingSquare, color) { // evaluar solo capturas legales
		value = pieceCaptureValue[move.Capture()] - see(b, move.To())
		b.UnMakeMove(move)
		return value, true
	}
	b.UnMakeMove(move)
	return 0, false
}
