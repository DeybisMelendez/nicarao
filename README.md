# Nicarao

Nicarao is a basic chess engine written in Go by Deybis Melendez.

## Features

### UCI Protocol

-   commands:
    -   uci
    -   isready
    -   ucinewgame
    -   position startpos | fen
    -   wtime, btime, winc, binc
    -   depth (TODO)
    -   movetime (TODO)
    -   movestogo (TODO)
    -   stop (TODO)
    -   ponder (TODO)
    -   mate (TODO)
    -   infinite (TODO)

### Board

-   12 Bitboards
-   Classical Approach for move generation
-   Little endian file and rank mapping
-   Zobrist Hashing

### Search

-   Principal Variation Search with ZeroWindow Search
-   Iterative Deepening
-   32 MB Transposition Table

-   Move Ordering:

    -   Hash Move
    -   Internal Iterative Deepening
    -   Simple Recapture Evaluation
    -   Killer Heuristic
    -   Counter Move Heuristic
    -   History Heuristic

-   Selectivity:
    -   Late Move Reduction (LMR)
    -   Null Move Pruning
    -   Quiescence
    -   Delta Pruning

### Evaluation

-   Material
-   Mobility
-   Piece Square Table (TODO)
-   Evaluation of pieces (TODO)
-   Tapered eval (TODO)
-   Basic Draw Evaluation (TODO)
-   Mop-up Evaluation (TODO)

### Testers

-   
