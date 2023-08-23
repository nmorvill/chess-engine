package chess

const all1 = uint64(0xFFFFFFFFFFFFFFFF)

var inBetween [64][64]uint64

func inBetweenCalc(sq1 int8, sq2 int8) uint64 {
	a2a7 := uint64(0x0001010101010100)
	b2g7 := uint64(0x0040201008040200)
	h1b7 := uint64(0x0002040810204080)

	btwn := (all1 << sq1) ^ (all1 << sq2)
	file := uint64((sq2 & 7) - (sq1 & 7))
	rank := uint64(((sq2 | 7) - sq1) >> 3)
	line := ((file & 7) - 1) & a2a7
	line += 2 * (((rank & 7) - 1) >> 58)
	line += (((rank - file) & 15) - 1) & b2g7
	line += (((rank + file) & 15) - 1) & h1b7
	line *= btwn & -btwn
	return line & btwn
}

func InBetweenInit() {

	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			inBetween[i][j] = inBetweenCalc(int8(i), int8(j))
		}
	}

}
