package main

import (
	"flag"
)

func main() {
	nth := GetNth()

	primes := []int{2}
	for i := 3; len(primes) < nth; i += 2 {
		prime := true

		for j := 0; j < len(primes) && primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
		}
	}

	prime := primes[len(primes)-1]

	print(prime)
}

func GetNth() int {
	var nth = flag.Int("nth", 100, "Find the Nth prime. (Default: 100)")
	flag.Parse()
	return *nth
}
