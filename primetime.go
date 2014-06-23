package main

import (
	"bufio"
	"flag"
	"math"
	"os"
	"strconv"
)

var maximum_saved float64 = math.Inf(1)

func Generate(ch chan<- int) {
	i := 0

	loaded_channel := make(chan int)
	go LoadDataFile(loaded_channel)
	var number int
	for {
		number = <-loaded_channel
		if number == -1 {
			maximum_saved = float64(i)
			break
		}
		i = number
		ch <- i
	}

	if i < 2 {
		i = 2
	}

	for ; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func LoadDataFile(ch chan<- int) {
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		ch <- number
	}
	ch <- -1
}

func AppendToDataFile(data string) {
	filename := "data.txt"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data + "\n"); err != nil {
		panic(err)
	}
}

func GetNth() int {
	var nth = flag.Int("nth", 100, "How many primes should be found, before stopping?")
	flag.Parse()
	return *nth
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel
	go Generate(ch)      // Launch Generate goroutine.

	length := GetNth()

	var i, prime int
	for i = 0; i < length; i++ {
		prime = <-ch
		if float64(prime) > maximum_saved {
			AppendToDataFile(strconv.Itoa(prime))
			print(prime, "\n")
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	print("prime", i, " := ", prime)
}
