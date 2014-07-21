#include <stdio.h>

int main(int argc, char *argv[])
{
	int nth = 100;
	
	if (argc == 2)
	{
		nth = atoi(argv[1]);
	}

	int primesFound = 1;
	int primes[nth];
	primes[0] = 2;

	int i, isPrime;
	int candidate = 3;
	while (primesFound < nth)
	{
		isPrime = 1;

		for (i = 0; i < primesFound && primes[i]*primes[i] <= candidate; i++)
		{
			if (candidate % primes[i] == 0)
			{
				isPrime = 0;
				break;
			}
		}

		if (isPrime)
		{
			primes[primesFound] = candidate;
			primesFound++;
		}

		candidate += 2;
	}

	printf("%d\n", primes[primesFound - 1]);

	return 0;
}
