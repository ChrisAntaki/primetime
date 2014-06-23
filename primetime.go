package main

import (
	"bufio"
	"flag"
	"math"
	"os"
	"strconv"
)

var largest_previous_prime = math.Inf(1)

func Generate(out chan<- int) {
	i := LoadDataFile(out)
	i |= 1

	largest_previous_prime = float64(i)

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
	var nth = flag.Int("nth", 100, "How many primes should be found, before stopping?")
	flag.Parse()
	return *nth
}

func GetAppendableFile() *os.File {
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
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
	new_primes := []int{}
	for i = 0; i < length; i++ {
		prime = <-channel
		if float64(prime) > largest_previous_prime {
			new_primes = append(new_primes, prime)
		}
		new_channel := make(chan int)
		go Filter(channel, new_channel, prime)
		channel = new_channel
	}

	SavePrimes(new_primes)

	print("prime", i, " := ", prime)
}
