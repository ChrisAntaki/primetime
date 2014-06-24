package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

var loading_has_completed = false
var largest_previous_prime = 0

func Generate(out chan<- int) {
	i := LoadDataFile(out)
	i |= 1

	loading_has_completed = true
	largest_previous_prime = i

	for {
		i++
		out <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func LoadDataFile(ch chan<- int) int {
	var prime int

	file, _ := os.Open("data.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prime, _ = strconv.Atoi(scanner.Text())
		ch <- prime
	}

	return prime
}

func GetNth() int {
	var nth = flag.Int("nth", 100, "Find the Nth prime. (Default: 100)")
	flag.Parse()
	return *nth
}

func GetAppendableFile() *os.File {
	flag := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile("data.txt", flag, 0600)
	dealbreaker(err)
	return file
}

func SavePrimes(primes []int) {
	file := GetAppendableFile()
	defer file.Close()

	message := ""
	for _, prime := range primes {
		message += strconv.Itoa(prime) + "\n"
	}

	_, err := file.WriteString(message)
	dealbreaker(err)
}

func dealbreaker(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	channel := make(chan int)
	go Generate(channel)

	length := GetNth()

	var i, prime int
	var primes []int

	for i = 0; i < length; i++ {
		// Get next prime.
		prime = <-channel

		// Save "new" primes.
		if loading_has_completed && prime > largest_previous_prime {
			primes = append(primes, prime)
		}

		// Pick a daisy.
		new_channel := make(chan int)
		go Filter(channel, new_channel, prime)
		channel = new_channel
	}

	SavePrimes(primes)

	print("prime", i, " := ", prime)
}
