package board

import "fmt"

/*
Utiliza el Little-Endian Rank-File Mapping
https://www.chessprogramming.org/Square_Mapping_Considerations

enum enumSquare {
  a1, b1, c1, d1, e1, f1, g1, h1,
  a2, b2, c2, d2, e2, f2, g2, h2,
  a3, b3, c3, d3, e3, f3, g3, h3,
  a4, b4, c4, d4, e4, f4, g4, h4,
  a5, b5, c5, d5, e5, f5, g5, h5,
  a6, b6, c6, d6, e6, f6, g6, h6,
  a7, b7, c7, d7, e7, f7, g7, h7,
  a8, b8, c8, d8, e8, f8, g8, h8
};
*/

func PrintBitboard(bitboard uint64) {
	fmt.Printf("   a b c d e f g h\n")
	for rank := Square(7); rank < 64; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := Square(0); file < 8; file++ {
			var square Square = rank*8 + file
			if GetBit(bitboard, square) > 0 {
				fmt.Printf("@ ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Nota: Usar estos métodos en condiciones que no afecten a la búsqueda (Duplicar uint64 puede ser costoso)
// Preferiblemente usar solo para mantener legibilidad

func GetBit(bitboard uint64, square Square) uint64 {
	return bitboard & (1 << square)
}
func SetBit(bitboard uint64, square Square) uint64 {
	return bitboard | (1 << square)
}
func PopBit(bitboard uint64, square Square) uint64 {
	return bitboard &^ (1 << square)
}
func SetPopBit(bitboard uint64, set Square, pop Square) uint64 {
	return (1 << set) | bitboard&^(1<<pop)
}
