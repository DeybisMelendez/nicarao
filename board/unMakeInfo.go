package board

type UnMakeInfo struct {
	Enpassant Square
	Castling  uint8
	Hash      uint64
}

func (s *Board) pushUnMakeInfo() {
	s.unMakeInfoIndex++
}

func (s *Board) saveUnMakeInfo() {
	s.unMakeInfoList[s.unMakeInfoIndex].Enpassant = s.Enpassant
	s.unMakeInfoList[s.unMakeInfoIndex].Castling = s.Castling
	s.unMakeInfoList[s.unMakeInfoIndex].Hash = s.Hash
}

func (s *Board) popUnMakeInfo() {
	s.unMakeInfoIndex--
	s.Enpassant = s.unMakeInfoList[s.unMakeInfoIndex].Enpassant
	s.Castling = s.unMakeInfoList[s.unMakeInfoIndex].Castling
	s.Hash = s.unMakeInfoList[s.unMakeInfoIndex].Hash
}

/*func clearUnMakeInfo() {
	unMakeInfoIndex = 0
}*/
