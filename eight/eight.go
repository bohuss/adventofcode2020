package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("eight/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	instructions := make([]string, 0)
	arguments := make([]int, 0)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := strings.Split(string(line), " ")
		instruction := sLine[0]
		argument, err := strconv.Atoi(sLine[1])
		if err != nil {
			panic("uable to parse argument as number")
		}
		instructions = append(instructions, instruction)
		arguments = append(arguments, argument)

	}

	ip := 0
	mem := 0
	seen := make(map [int] bool)
	for seen[ip] != true {
		seen[ip] = true
		switch instructions[ip] {
		case "nop": {
			ip++
		}
		case "acc": {
			mem += arguments[ip]
			ip++
		}
		case "jmp": {
			ip += arguments[ip]
		}
		}
	}
	ans1 := mem
	ans2 := -1
	for change := range instructions {
		switch instructions[change] {
		case "acc": {
			continue
		}
		case "nop": {
			instructions[change] = "jmp"
		}
		case "jmp": {
			instructions[change] = "nop"
		}
		}

		ip := 0
		mem := 0
		seen := make(map [int] bool)
		for seen[ip] != true && ip < len(instructions) {
			seen[ip] = true
			switch instructions[ip] {
			case "nop": {
				ip++
			}
			case "acc": {
				mem += arguments[ip]
				ip++
			}
			case "jmp": {
				ip += arguments[ip]
			}
			}
		}
		if ip == len(instructions) {
			if ans2 != -1 {
				panic("more than one correct answer found")
			}
			ans2 = mem
		}

		switch instructions[change] {
		case "nop": {
			instructions[change] = "jmp"
		}
		case "jmp": {
			instructions[change] = "nop"
		}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

