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
