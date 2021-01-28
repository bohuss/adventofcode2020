package main

import (
	"bufio"
	"os"
)

type Cup struct {
	Next *Cup
	Val int32
}

func main() {

	input, err := os.Open("twentythree/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	line, _, err := reader.ReadLine()
	cups := make([]Cup, 0)

	for _, v := range string(line) {
		cups = append(cups, Cup{Val: v - '0'})
	}
	cupIds := make([]int, len(cups) + 1)
	for i, c := range cups {
		cupIds[c.Val] = i
	}
	for i := 1; i < len(cups); i++ {
		cups[i - 1].Next = &cups[i]
	}
	cups[len(cups) - 1].Next = &cups[0]

	currentCup := &cups[0]
	for i := 0; i <100; i++ {
		makeOneMove(currentCup, cups, cupIds)
		currentCup = currentCup.Next
	}
	printCup := cups[cupIds[1]].Next
	for printCup.Val != 1 {
		print(printCup.Val)
		printCup = printCup.Next
	}
	println()

	cups = make([]Cup, 1000000)
	for i := range cups {
		if i < len(line) {
			cups[i] = Cup{Val: int32(string(line)[i] - '0')}
		} else {
			cups[i] = Cup{Val: int32(i + 1)}
		}
	}

	cupIds = make([]int, len(cups) + 1)
	for i, c := range cups {
		cupIds[c.Val] = i
	}
	for i := 1; i < len(cups); i++ {
		cups[i - 1].Next = &cups[i]
	}
	cups[len(cups) - 1].Next = &cups[0]

	currentCup = &cups[0]
	for i := 0; i <10000000; i++ {
		makeOneMove(currentCup, cups, cupIds)
		currentCup = currentCup.Next
	}

	printCup = cups[cupIds[1]].Next
	println(int64(printCup.Val) * int64(printCup.Next.Val))

}

func makeOneMove(currentCup *Cup, cups []Cup, cupIds []int) {
	removed := make([]*Cup, 3)
	removed[0] = currentCup.Next
	removed[1] = currentCup.Next.Next
	removed[2] = currentCup.Next.Next.Next
	currentCup.Next = removed[2].Next
	destinationId := currentCup.Val - 1
	if destinationId == 0 {
		destinationId = int32(len(cups))
	}
	for j := 0; j < len(removed); j++ {
		for _, c := range removed {
			if c.Val == destinationId {
				destinationId--
				if destinationId == 0 {
					destinationId = int32(len(cups))
				}
			}
		}
	}
	destinationCup := &cups[cupIds[destinationId]]
	removed[2].Next = destinationCup.Next
	destinationCup.Next = removed[0]
}
