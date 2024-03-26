package primenumbers

// FindPrimes returns a slice of prime numbers up to the given limit
func FindPrimes(limit int) []int {
	primes := make([]bool, limit + 1)
	for i := 2; i <= limit; i++ {
		primes[i] = true
	}

	for p := 2; p * p <= limit; p++ {
		if primes[p] {
			for i := p * p; i <= limit; i += p {
				primes[i] = false
			}
		}
	}

	var result []int
	for p := 2; p <= limit; p++ {
		if primes[p] {
			result = append(result, p)
		}
	}
	return result
}

// IsPrime returns true if n is a prime number
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
