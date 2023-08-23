package chess

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func NumOfSetBits(n uint64) int {
	count := 0
	for n != 0 {
		count += int(n) & 1
		n >>= 1
	}
	return count
}

func AbsDelta(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
