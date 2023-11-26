package search

const ttSize uint64 = 64 * MB / 18
const killerMovesMaxPly = 12
const LMRFullDepthMoves = 3
const LMReductionLimit = 3
const deltaPruning int16 = 900

//Switches
const LMRisActive bool = true
const orderingMoveIsActive bool = true
const deltaPruningisActive bool = true
