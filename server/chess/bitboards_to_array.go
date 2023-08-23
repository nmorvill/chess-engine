package chess

func GetArrayFromBitboard(b Board) [64]int {
	var ret [64]int
	for i := 0; i < 6; i++ {
		wps := getAllIndexesOfBitBoard(b.getPieceSet(PieceType(i), White))
		for _, k := range wps {
			ret[k] = (i + 1) + 8
		}
		bps := getAllIndexesOfBitBoard(b.getPieceSet(PieceType(i), Black))
		for _, k := range bps {
			ret[k] = (i + 1) + 16
		}
	}
	return ret
}

func getAllIndexesOfBitBoard(u uint64) []int8 {
	var ret []int8
	for k := u != 0; k; k = u != 0 {
		i := bitScanForward(u)
		ret = append(ret, i)
		u &= u - 1
	}
	return ret
}
