package primenumbers

func HCFEuclidean(a, b int) int {
	if a == 0 {
		return b
	}
	return HCFEuclidean(b % a, a)
}

func HCF(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}

	var result int
	if a < b {
		result = a
	} else {
		result = b
	}

	for result > 0 {
		if a % result == 0 && b % result == 0 {
			break
		}
		result--
	}
	return result
}
