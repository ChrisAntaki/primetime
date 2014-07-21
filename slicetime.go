package main

import (
	"flag"
)

func main() {
	nth := GetNth()

	primes := make([]int, nth)
	primes[0] = 2
	primesFound := 1

	for i := 3; primesFound < nth; i += 2 {
		prime := true

		for j := 0; j < len(primes) && primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes[primesFound] = i
			primesFound++
		}
	}

	prime := primes[primesFound-1]

	print(prime)
	print("\n")
}

func GetNth() int {
	var nth = flag.Int("nth", 100, "Find the Nth prime. (Default: 100)")
	flag.Parse()
	return *nth
}
