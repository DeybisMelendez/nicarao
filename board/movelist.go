package board

//type Move uint32

type MoveList struct {
	List  [MaxMoveList]Move //Cantidad de movimientos máximos definidos por MaxMoveList uint8 = 128
	Index uint8             // Considerando que la máxima cantidad de jugadas legales teóricas es 218 segun chessprogramming
}

//Agrega un nuevo movimiento a la lista sin utilizar append
func (s *MoveList) Push(move Move) {
	s.List[s.Index] = move
	s.Index++
}
