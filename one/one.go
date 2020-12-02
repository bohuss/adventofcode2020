package main

import (
	"fmt"
	"os"
)

const targetResult = 2020

func main() {

	fmt.Println(os.Getwd())
	input, err := os.Open("one/input.txt")

	if err != nil {
		panic(err)
	}

	var inNum int
	all := make([]int, 0)

	for {
		_, err := fmt.Fscanln(input, &inNum)
		if err != nil {
			break
		}

		for _,x := range all {
			if x+inNum == targetResult {
				fmt.Printf("%d * %d = %d\n", x, inNum, x * inNum)
				break
			}
		}

		for _,x := range all {
			for _,y := range all {
				if x+y+inNum == targetResult {
					fmt.Printf("%d * %d * %d = %d\n", x, y, inNum, x*y*inNum)
					break
				}
			}
		}

		all = append(all, inNum)
	}

	fmt.Printf("Done, scanned %d numbers", len(all))

}