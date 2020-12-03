package main

import (
	"fmt"
	"os"
)

func main() {

	input, err := os.Open("two/input.txt")

	if err != nil {
		panic(err)
	}

	ans1, ans2 := 0, 0
	var a,b int
	var c rune
	var password string

	for {
		_, err := fmt.Fscanf(input, "%d-%d %c: %s", &a, &b, &c, &password)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		if ok1(password, c, a, b) {
			ans1++
		}

		if ok2(password, uint8(c), a - 1, b - 1) {
			ans2++
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func ok1(s string, c rune, a, b int)bool {
	count :=0
	for _, x := range s {
		if x == c {
			count++
			if count > b {
				return false
			}
		}
	}
	return count >= a
}

func ok2(s string, c uint8, a, b int)bool {
	return (s[a] == c) != (s[b] == c)
}