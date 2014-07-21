#include <iostream>
#include <cstdlib>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
	int nth = 100;

	if (argc == 2) {
		nth = atoi(argv[1]);
	}

	vector<int> primes;
	primes.push_back(2);

	int i;
	for (i = 3; primes.size() < nth; i += 2) {
		bool prime = true;

		for (int j = 0; j < primes.size() && primes[j]*primes[j] <= i; j++) {
			if (i % primes[j] == 0) {
				prime = false;
				break;
			}
		}

		if (prime) {
			primes.push_back(i);
		}
	}

	int final = primes.back();
	cout << final << "\n";
}
