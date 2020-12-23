package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contains map[string] int

type Bag struct {
	color string
	Contains Contains
}

type Bags map[string]*Bag

func main() {

	input, err := os.Open("seven/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	colors := make([]string, 0)
	insides := make([][]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)

		parts := strings.Split(sLine, " bags contain ")
		colors = append(colors, parts[0])
		insides = append(insides, strings.Split(parts[1], ", "))
	}

	bags := make(Bags)
	for id, color := range colors {
		bags[color] = parseInsides(color, insides[id])
	}

	canContainShinyGold := make(map[string]bool)
	toProcess := make([]string, 0)
	for _, bag := range bags {
		if bag.Contains["shiny gold"] != 0 {
			canContainShinyGold[bag.color] = true
			toProcess = append(toProcess, bag.color)
		}
	}

	for len(toProcess) > 0 {
		peek := len(toProcess) - 1
		color := toProcess[peek]
		toProcess = toProcess[:peek]
		canContainShinyGold[color] = true

		for _, bag := range bags {
			if bag.Contains[color] != 0 {
				canContainShinyGold[bag.color] = true
				toProcess = append(toProcess, bag.color)
			}
		}
	}

	ans1 := len(canContainShinyGold)
	ans2 := 0

	count := make([]int, 0)
	toProcess = append(toProcess, "shiny gold")
	count = append(count, 1)

	for len(toProcess) > 0 {
		peek := len(toProcess) - 1
		color := toProcess[peek]
		toProcess = toProcess[:peek]
		num := count[peek]
		count = count[:peek]
		ans2 += num
		for color, num2 := range bags[color].Contains {
			count = append(count, num * num2)
			toProcess = append(toProcess, color)
		}
	}
	ans2--

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func parseInsides(color string, insides []string) *Bag {
	contains := make(Contains)
	for _, inside := range insides{
		if strings.HasPrefix(inside, "no") {
			continue
		}
		split := strings.Split(inside, " ")
		num, err := strconv.Atoi(split[0])
		if err != nil {
			panic("unable to parse number of bags")
		}
		color := split[1] + " " + split[2]
		contains[color] = num
	}

	return &Bag {
		color: color,
		Contains: contains,
	}
}

