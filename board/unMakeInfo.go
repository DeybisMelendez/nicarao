package board

type UnMakeInfo struct {
	Hash          uint64
	Castling      uint8
	Enpassant     Square
	HalfmoveClock uint8
}

func (s *Board) pushUnMakeInfo() {
	s.unMakeInfoIndex++
}

func (s *Board) saveUnMakeInfo() {
	s.unMakeInfoList[s.unMakeInfoIndex].Enpassant = s.Enpassant
	s.unMakeInfoList[s.unMakeInfoIndex].Castling = s.Castling
	s.unMakeInfoList[s.unMakeInfoIndex].Hash = s.Hash
	s.unMakeInfoList[s.unMakeInfoIndex].HalfmoveClock = s.HalfmoveClock
}

func (s *Board) popUnMakeInfo() {
	s.unMakeInfoIndex--
	s.Enpassant = s.unMakeInfoList[s.unMakeInfoIndex].Enpassant
	s.Castling = s.unMakeInfoList[s.unMakeInfoIndex].Castling
	s.Hash = s.unMakeInfoList[s.unMakeInfoIndex].Hash
	s.HalfmoveClock = s.unMakeInfoList[s.unMakeInfoIndex].HalfmoveClock
}

/*func clearUnMakeInfo() {
	unMakeInfoIndex = 0
}*/
