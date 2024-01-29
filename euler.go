package main

import (
	"log"
	"math"
)

var initSieve = 100000 
var primes []int

// sieve creates the primes below initSieve and populates primes
func sieve(n int) {
	arr := make([]int, n+1)
	arr[0], arr[1] = 1, 1
	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			for j := 2*i; j <= n; j += i {
				arr[j] = 1
			}
		}
	}
	for i := range arr {
		if arr[i] == 0 {
			primes = append(primes, i)
		}
	}
}

func copyIntArray(arr []int) []int {
	ret := make([]int, len(arr))
	copy(ret, arr)
	return ret
}

func PrimesInRange(a, b int) []int {
	arr := make([]int, b-a+1)
	if a == 1 {
		arr[0] = 1
	}
	if a == 0 {
		arr[0], arr[1] = 1, 1
	}
	m := IntSqrt(b)
	for _, p := range primes {
		if p > m {
			break
		}
		start := 0
		if p >= a && p <= b {
			start = p+p
		} else if a%p == 0 {
			start = a
		} else {
			start = a + p - a%p
		}
		for i := start; i <= b; i += p {
			arr[i-a] = 1
		}
	}
	ret := []int{}
	for i, n := range arr {
		if n == 0 {
			ret = append(ret, i+a)
		}
	}
	return ret
}


// IntSqrt returns the integer square root of n
func IntSqrt(n int) int {
        f := math.Sqrt(float64(n))
        return int(f)
}

// Factors returns the prime factors composing n
func Factors(n int) [] int {
	arr := []int{}
	for _, p := range primes {
		if p*p > n {
			break
		}
		for n%p == 0 {
			arr = append(arr, p)
			n /= p
		}
	}
	if n > 1 {
		arr = append(arr, n)
	}
	return arr
}

// UniqFactors returns the unique prime factors of n
func UniqFactors(n int) []int {
	f := Factors(n)
	arr := []int{}
	for _, p := range f {
		if len(arr) == 0 {
			arr = append(arr, p)
		}
		if arr[len(arr)-1] != p {
			arr = append(arr, p)
		}
	}
	return arr
}

// GCD returns the greatest common divisor of a, b
func GCD(a, b int) int {
	if b%a == 0 {
		return a
	}
	return GCD(b%a, a)
}

// Phi returns the euler totient of n
func Phi(n int) int {
	ret := n
	for _, p := range UniqFactors(n) {
		ret *= p-1
		ret /= p
	}
	return ret
}

func init() {
	sieve(initSieve)
}


func main() {
	log.Println(len(PrimesInRange(0,200)))
	log.Println(len(PrimesInRange(100, 200)))
}
