package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, err := os.Open("twentyfive/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	line, _, err := reader.ReadLine()
	card, err := strconv.Atoi(string(line))

	if err != nil {
		panic("unable to parse card public key as number")
	}

	line, _, err = reader.ReadLine()
	door, err := strconv.Atoi(string(line))

	if err != nil {
		panic("unable to parse door public key as number")
	}

	cardLoopSize := 0
	val := 1

	for val != card {
		val *= 7
		val %= 20201227
		cardLoopSize++
	}

	doorLoopSize := 0
	val = 1

	for val != door {
		val *= 7
		val %= 20201227
		doorLoopSize++
	}

	val = 1
	for i := 0; i < cardLoopSize; i++ {
		val *= door
		val %= 20201227
	}

	fmt.Println(val)
}
