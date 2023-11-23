package board

type UnMakeInfo struct {
	Enpassant Square
	Castling  uint8
}

var unMakeInfoList [MaxPly]UnMakeInfo
var unMakeInfoIndex int

func pushUnMakeInfo(enpassant Square, castling uint8) {
	unMakeInfoList[unMakeInfoIndex].Enpassant = enpassant
	unMakeInfoList[unMakeInfoIndex].Castling = castling
	unMakeInfoIndex++
}
func popUnMakeInfo() (Square, uint8) {
	unMakeInfoIndex--
	return unMakeInfoList[unMakeInfoIndex].Enpassant, unMakeInfoList[unMakeInfoIndex].Castling
}

func clearUnMakeInfo() {
	//unMakeInfoList = [MaxPly]UnMakeInfo{}
	unMakeInfoIndex = 0
}
