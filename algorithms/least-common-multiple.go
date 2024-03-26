package primenumbers

// LCM returns the least common multiple of two numbers
func LCMviaGCD(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / HCF(a, b)
}

// LCM returns the least common multiple of two numbers
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	var smaller, larger int
	if a > b {
		smaller, larger = b, a
	} else {
		smaller, larger = a, b
	}

	for i := larger; ; i += larger {
		if i % smaller == 0 {
			return i
		}
	}
}