package primenumbers

import (
	"testing"
)

func TestFindPrimes(t *testing.T) {
	primesTest := []struct {
		limit    int
		expected []int
	}{
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
	}

	for _, test := range primesTest {
		primes := FindPrimes(test.limit)
		if len(primes) != len(test.expected) {
			t.Fatalf("Expected %d primes, got %d", len(test.expected), len(primes))
		}
		for i, p := range primes {
			if p != test.expected[i] {
				t.Errorf("Expected %d, got %d", test.expected[i], p)
			}
		}
	}
}

func TestHCF(t *testing.T) {
	hcfTest := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},
		{10, 3, 1},
		{3, 10, 1},
		{270, 192, 6},
		{10, 0, 10},
		{0, 10, 10},
		{0, 0, 0},
		{1, 1, 1},
	}

	for _, test := range hcfTest {
		hcf := HCF(test.a, test.b)
		if hcf != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, hcf)
		}
	}
}

func TestLCM(t *testing.T) {
	lcmTest := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 10},
		{10, 3, 30},
		{3, 10, 30},
		{270, 192, 8640},
		{10, 0, 0},
		{0, 10, 0},
		{0, 0, 0},
		{1, 1, 1},
	}

	for _, test := range lcmTest {
		lcm := LCM(test.a, test.b)
		if lcm != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, lcm)
		}
	}
}

func TestIsPrime(t *testing.T) {
	primeTest := []struct {
		n        int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
	}

	for _, test := range primeTest {
		prime := IsPrime(test.n)
		if prime != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, prime)
		}
	}
}

func TestIsPerfect(t *testing.T) {
	perfectTest := []struct {
		n        int
		expected bool
	}{
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{78, false},
		{1000, false},
		{4096, false},
	}

	for _, test := range perfectTest {
		perfect := IsPerfectNumber(test.n)
		if perfect != test.expected {
			t.Errorf("%v Expected %t, got %t", test, test.expected, perfect)
		}
	}
}

func TestFindFactors(t *testing.T) {
	factorsTest := []struct {
		n        int
		expected []int
	}{
		{1, []int{1}},
		{6, []int{1, 2, 3, 6}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{496, []int{1, 2, 4, 8, 16, 31, 62, 124, 248, 496}},
		{8128, []int{1, 2, 4, 8, 16, 32, 64, 127, 254, 508, 1016, 2032, 4064, 8128}},
	}

	for _, test := range factorsTest {
		factors := FindFactors(test.n)
		if len(factors) != len(test.expected) {
			t.Fatalf("Expected %d factors, got %d", len(test.expected), len(factors))
		}
		for i, f := range factors {
			if f != test.expected[i] {
				t.Errorf("Expected %d, got %d", test.expected[i], f)
			}
		}
	}
}

func TestFindPerfectNumbers(t *testing.T) {
	perfectTest := []struct {
		limit    int
		expected []int
	}{
		{10, []int{6}},
		{1000, []int{6, 28, 496}},
		{10000, []int{6, 28, 496, 8128}},
		// {100000000, []int{6, 28, 496, 8128, 33550336}},
	}

	for _, test := range perfectTest {
		perfects := FindPerfectNumbers(test.limit)
		if len(perfects) != len(test.expected) {
			t.Fatalf("Expected %d perfect numbers, got %d", len(test.expected), len(perfects))
		}
		for i, p := range perfects {
			if p != test.expected[i] {
				t.Errorf("Expected %d, got %d", test.expected[i], p)
			}
		}
	}
}