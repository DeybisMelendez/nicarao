package search

const ttSize uint64 = 32 * MB / 18 //18 es el tamaño de bytes que ocupa cada entry de tt
const killerMovesMaxPly = 12
const LMReductionLimit = 3
const deltaPruning int16 = 200
const internalIDminDepth = 2

//Switches
const LMRisActive bool = false
const orderingMoveIsActive bool = true
const deltaPruningisActive bool = true
const internalIDIsActive bool = true
const isTranspositionTableActive bool = true
