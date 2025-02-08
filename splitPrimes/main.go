package main

import (
	"fmt"
	"strconv"
)

const (
	MOD    = 1000000007
	MAXVAL = 1000000
)

// sieve generates a boolean slice of length MAXVAL+1
// where prime[i] is true if i is prime, false otherwise.
func sieve() []bool {
	prime := make([]bool, MAXVAL+1)
	// Initially, mark all as prime.
	for i := 0; i <= MAXVAL; i++ {
		prime[i] = true
	}
	prime[0], prime[1] = false, false

	for p := 2; p*p <= MAXVAL; p++ {
		if prime[p] {
			for multiple := p * p; multiple <= MAXVAL; multiple += p {
				prime[multiple] = false
			}
		}
	}
	return prime
}

// countPrimeSplits computes the number of ways to split the string s
// such that every segment (without leading zeros) is a prime number
// within the range [2, MAXVAL]. The result is returned modulo MOD.
func countPrimeSplits(s string, prime []bool) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1 // There's one way to split an empty string

	// For each index i (1-indexed for dp, corresponding to the first i characters of s)
	for i := 1; i <= n; i++ {
		// Check only the last 6 digits (or fewer if i < 6)
		for l := 1; l <= 6; l++ {
			if i-l < 0 {
				break
			}
			// Get the substring s[i-l:i]
			if s[i-l] == '0' {
				// Skip segments with a leading zero.
				continue
			}

			// Convert the substring to an integer.
			num, err := strconv.Atoi(s[i-l : i])
			if err != nil {
				continue
			}
			// If the number exceeds our allowed maximum, skip it.
			if num > MAXVAL {
				continue
			}

			// If the number is prime, update dp.
			if prime[num] {
				dp[i] = (dp[i] + dp[i-l]) % MOD
			}
		}
	}
	return dp[n]
}

func main() {
	// Precompute primes up to MAXVAL.
	prime := sieve()

	// Calculate and print the number of valid splits.
	result := countPrimeSplits("11375", prime)
	fmt.Println(result)
}
