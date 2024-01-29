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
	arr := make([]int, b-1+1)
	for _, p := range primes {


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
	log.Println(GCD(155, 30))
	log.Println(UniqFactors(2*2*2*2*3*3*5))
	log.Println(Phi(77))
	log.Println(copyIntArray([]int{1,2,3,4}))
}
