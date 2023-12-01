package search

const ttSize uint64 = 32 * MB / 18 //18 es el tamaño de bytes que ocupa cada entry de tt
const killerMovesMaxPly = 12
const LMRFullDepthMoves = 20 //En orden de la puntuación de movimientos, hasta Counter Moves como mínimo
const LMReductionLimit = 3
const deltaPruning int16 = 900
const internalIDminDepth = 2
const nullMovePruningLimit = 4
const nullMovePruningR = 3

//Switches
const LMRisActive bool = true
const orderingMoveIsActive bool = true
const deltaPruningisActive bool = true
const internalIDIsActive bool = true
const isNullMovePruningActive bool = true
const isTranspositionTableActive bool = true
