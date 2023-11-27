package search

const ttSize uint64 = 64 * MB / 18
const killerMovesMaxPly = 12
const LMRFullDepthMoves = 3 //En orden de la puntuación de movimientos, hasta Counter Moves como mínimo
const LMReductionLimit = 3
const deltaPruning int16 = 900
const internalIDminDepth = 2

//Switches
const LMRisActive bool = true
const orderingMoveIsActive bool = true
const deltaPruningisActive bool = true
const internalIDIsActive bool = true
const isNullMovePruningActive bool = true
