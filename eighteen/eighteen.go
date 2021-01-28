package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("eighteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		ans1 += calc(sLine, false)
		ans2 += calc(sLine, true)

	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func calc(s string, priorityForAddition bool) int {
	pattern := regexp.MustCompile("\\([ \\d+*]+\\)")
	one := pattern.FindString(s)
	if one == "" {
		exp := strings.Split(s, " ")

		if priorityForAddition {
			exp2 := make([]string, 0)
			a, err := strconv.Atoi(exp[0])
			if err != nil {
				panic("unable to parse first number")
			}
			for i := 1; i < len(exp); i+=2 {
				arg, err := strconv.Atoi(exp[i + 1])
				if err != nil {
					panic("unable to parse number in expression")
				}
				switch exp[i] {
				case "+":
					a += arg
				case "*":
					exp2 = append(exp2, strconv.Itoa(a))
					exp2 = append(exp2, "*")
					a = arg
				}
			}
			exp2 = append(exp2, strconv.Itoa(a))
			exp = exp2
		}
		now, err := strconv.Atoi(exp[0])
		if err != nil {
			panic("unable to parse first number in expression")
		}
		for i := 1; i < len(exp); i+=2 {
			arg, err := strconv.Atoi(exp[i + 1])
			if err != nil {
				panic("unable to parse number in expression")
			}
			switch exp[i] {
			case "*":
				now *= arg
			case "+":
				now += arg
			}
		}
		return now
	}
	ones := strconv.Itoa(calc(one[1:len(one) - 1], priorityForAddition))
	return calc(strings.Replace(s, one, ones, 1), priorityForAddition)
}
