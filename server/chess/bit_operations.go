package chess

/*
https://www.chessprogramming.org/Bitboard_Serialization
*/

var index64 = [64]int8{
	0, 1, 48, 2, 57, 49, 28, 3,
	61, 58, 50, 42, 38, 29, 17, 4,
	62, 55, 59, 36, 53, 51, 43, 22,
	45, 39, 33, 30, 24, 18, 12, 5,
	63, 47, 56, 27, 60, 41, 37, 16,
	54, 35, 52, 21, 44, 32, 23, 11,
	46, 26, 40, 15, 34, 20, 31, 10,
	25, 14, 19, 9, 13, 8, 7, 6}

/**
 * bitScanForward
 * @author Martin Läuter (1997)
 *         Charles E. Leiserson
 *         Harald Prokop
 *         Keith H. Randall
 * "Using de Bruijn Sequences to Index a 1 in a Computer Word"
 * @param bb bitboard to scan
 * @precondition bb != 0
 * @return index (0..63) of least significant one bit
 */
func bitScanForward(bb uint64) int8 {
	debruijn64 := uint64(0x03f79d71b4cb0a89)
	return index64[((bb&-bb)*debruijn64)>>58]
}

func nortOne(b uint64) uint64 {
	return b << 8
}

func soutOne(b uint64) uint64 {
	return b >> 8
}

func eastOne(b uint64) uint64 {
	return (b & ^uint64(0x8080808080808080)) << 1
}

func westOne(b uint64) uint64 {
	return (b & ^uint64(0x101010101010101)) >> 1
}

func nortEastOne(b uint64) uint64 {
	return nortOne(eastOne(b))
}

func nortWestOne(b uint64) uint64 {
	return nortOne(westOne(b))
}

func soutEastOne(b uint64) uint64 {
	return soutOne(eastOne(b))
}

func soutWestOne(b uint64) uint64 {
	return soutOne(westOne(b))
}

func fileDistance(sq1, sq2 int) int {
	return AbsDelta(file(sq1), file(sq2))
}

func file(sq int) int {
	return sq & 7
}

func rank(sq int) int {
	return sq >> 3
}

func rankDistance(sq1, sq2 int) int {
	return AbsDelta(rank(sq1), rank(sq2))
}

func squareDistance(sq1, sq2 int) int {
	return Max(fileDistance(sq1, sq2), rankDistance(sq1, sq2))
}
