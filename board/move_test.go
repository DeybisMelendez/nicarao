package board

import (
	"testing"
)

func TestMoveFunctions(t *testing.T) {
	// Crear un movimiento
	move := NewMove(King, H8, H8, King, 0, EnpassantCapture)

	// Verificar las funciones Piece(), From(), To(), Capture(), Promotion(), Flag()
	if move.Piece() != King {
		t.Errorf("Expected piece to be 6, got %v", move.Piece())
	}

	if move.From() != H8 {
		t.Errorf("Expected from square to be 63, got %v", move.From())
	}

	if move.To() != H8 {
		t.Errorf("Expected to square to be 63, got %v", move.To())
	}

	if move.Capture() != King {
		t.Errorf("Expected capture piece to be 6, got %v", move.Capture())
	}

	if move.Promotion() != None {
		t.Errorf("Expected promotion piece to be 6, got %v", move.Promotion())
	}

	if move.Flag() != EnpassantCapture {
		t.Errorf("Expected flag to be 7, got %v", move.Flag())
	}
}
func TestMoveToString(t *testing.T) {
	m := NewMove(Pawn, A2, A4, None, None, QuietMoves)
	result := m.MoveToString()
	if result != "a2a4" {
		t.Errorf("Error: Movimiento esperado: %s, Resultado: %s", "a2a4", result)
	}
	type testCase struct {
		from   Square
		to     Square
		promo  Piece
		result string
	}

	testCases := []testCase{
		{G1, F3, None, "g1f3"},
		{E2, E4, None, "e2e4"},
		{A1, H8, Queen, "a1h8q"},
		{G1, G8, Rook, "g1g8r"},
	}

	for _, tc := range testCases {
		m := NewMove(None, tc.from, tc.to, None, tc.promo, QuietMoves)
		result := m.MoveToString()
		if result != tc.result {
			t.Errorf("Error: Movimiento esperado: %s, Resultado: %s", tc.result, result)
		}
	}
}
