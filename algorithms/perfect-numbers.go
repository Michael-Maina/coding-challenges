package primenumbers

import (
	"math"
	"slices"
)

// IsPerfectNumber returns true if the given number is a perfect number
func IsPerfectNumber(n int) bool {
	if n <= 1 {
		return false
	}

	if sum(FindFactors(n)) == 2 * n {
		return true
	}
	return false
}

// FindPerfectNumbers returns a slice of perfect numbers up to the given limit
func FindPerfectNumbers(limit int) []int {
	var result []int
	for i := 2; i <= limit; i++ {
		if IsPerfectNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

// FindFactors returns a slice of all the factors of a number n
func FindFactors(n int) []int {
	if n == 1 {
		return []int{1}
	}

	var result []int
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrtN; i++ {
		if n % i == 0 {
			result = append(result, i)
			if n / i != i {
				result = append(result, n / i)
			}
		}
	}
	slices.Sort(result) // Sorted for testing purposes
	return result
}

// sum returns the sum of all the elements in a slice
func sum(factors []int) int {
	var sum int
	for _, f := range factors {
		sum += f
	}
	return sum
}
