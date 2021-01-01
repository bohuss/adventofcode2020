package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	fromA int
	toA int
	fromB int
	toB int
	name string
}

func main() {

	input, err := os.Open("sixteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 1
	inputType := 0
	fields := make ([]Field, 0)
	yourTicket := make ([]int, 0)
	nearbyTickets := make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)

		if len(sLine) == 0 {
			inputType++
			if inputType == 3 {
				break
			}
		} else {
			switch inputType {
			case 0:
				{
					fields = append(fields, parseField(sLine))
				}
			case 1:
				{
					if strings.HasPrefix(sLine, "your") {
						continue
					}
					for _, v := range strings.Split(sLine, ",") {
						num, err := strconv.Atoi(v)
						if err != nil {
							panic("unable to parse your ticket")
						}
						yourTicket = append(yourTicket, num)
					}
				}

			case 2:
				{
					if strings.HasPrefix(sLine, "nearby") {
						continue
					}
					ticket := make ([]int, 0)
					for _, v := range strings.Split(sLine, ",") {
						num, err := strconv.Atoi(v)
						if err != nil {
							panic("unable to parse nearby ticket")
						}
						ticket = append(ticket, num)
					}
					nearbyTickets = append(nearbyTickets, ticket)
				}
			}
		}
	}
	fieldMapping := make(map[int] []int)
	for i := range fields{
		fieldMapping[i] = make([]int , 0)
		for j := range fields {
			fieldMapping[i] = append(fieldMapping[i], j)
		}
	}

	for _, ticket := range nearbyTickets {
		valid := true
		for i := range fields {

			validOptions := 0
			for id := range fields {
				if canBe(ticket[i], fields[id]) {
					validOptions++
				}
			}

			if validOptions == 0 {
				ans1 += ticket[i]
				valid = false
				continue
			}
		}
		if valid {
			for i := range fields {
				filteredMapping := make([]int, 0)
				for _, v := range fieldMapping[i] {
					if canBe(ticket[i], fields[v]) {
						filteredMapping = append(filteredMapping, v)
					}
				}
				fieldMapping[i] = filteredMapping
			}
		}
	}

	assigned := make(map[int] bool)

	change := true
	for change {
		change = false
		for i, v := range fieldMapping {
			if len(v) == 1 {
				assignedId := v[0]
				if assigned[assignedId] {
					continue
				}
				assigned[assignedId] = true
				change = true
				for j, v := range fieldMapping {
					if j == i {
						continue
					}
					for id, x := range v {
						if x == assignedId {
							fieldMapping[j] = append(fieldMapping[j][0:id], fieldMapping[j][id+1:]...)
						}
					}
				}
			}
		}
	}

	for i := range fields {
		if strings.HasPrefix(fields[fieldMapping[i][0]].name, "departure"){
			ans2 *= yourTicket[i]
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func canBe(i int, field Field) bool {
	if i >= field.fromA && i <= field.toA {
		return true
	}
	if i >= field.fromB && i <= field.toB {
		return true
	}
	return false

}

func parseField(s string)Field {
	a := strings.Split(s, ":")
	name := a[0]
	b := strings.Split(a[1], " ")
	aLimit := strings.Split(b[1], "-")
	bLLimit := strings.Split(b[3], "-")
	fromA, err := strconv.Atoi(aLimit[0])
	if err != nil {
		panic("unable to parse fromA")
	}
	fromB, err := strconv.Atoi(bLLimit[0])
	if err != nil {
		panic("unable to parse fromA")
	}
	toA, err := strconv.Atoi(aLimit[1])
	if err != nil {
		panic("unable to parse fromA")
	}
	toB, err := strconv.Atoi(bLLimit[1])
	if err != nil {
		panic("unable to parse fromA")
	}
	return Field{
		fromA: fromA,
		toA:   toA,
		fromB: fromB,
		toB:   toB,
		name:  name,
	}

}
