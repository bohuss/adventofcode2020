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

	input, err := os.Open("nineteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0
	phase := 1
	rules := make(map[int] string)
	rules2 := make(map[int] string)
	reg := ""
	reg2 := ""
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
			phase = 2

			reg = rules[0]
			for pos := strings.Index(reg, "#"); pos > -1; pos = strings.Index(reg, "#"){
				e := pos + 1
				for reg[e] >='0' && reg[e] <= '9' {
					e++
				}

				num, err := strconv.Atoi(reg[pos+1:e])
				if err != nil {
					panic("unable to parse id after #")
				}
				replacement := rules[num]

				if replacement == "" {
					panic("empty replacement")
				}
				reg = fmt.Sprintf("%s%s%s", reg[:pos], replacement, reg[e:])

			}
			reg = "^" + reg + "$"

			reg2 = rules2[0]
			rules2[8] = "(#42|#42#8)"
			rules2[11] = "(#42#31|#42#11#31)"
			num8 := 0
			num11 := 0
			const MaxRecursion = 4
			for pos := strings.Index(reg2, "#"); pos > -1; pos = strings.Index(reg2, "#"){
				e := pos + 1
				for reg2[e] >='0' && reg2[e] <= '9' {
					e++
				}

				num, err := strconv.Atoi(reg2[pos+1:e])
				if err != nil {
					panic("unable to parse id after #")
				}
				replacement := rules2[num]

				if replacement == "" {
					panic("empty replacement")
				}
				reg2 = fmt.Sprintf("%s%s%s", reg2[:pos], replacement, reg2[e:])

				if num == 8 {
					num8++
					if num8 == MaxRecursion {
						rules2[8] = "#42"
					}
				}
				if num == 11 {
					num11++
					if num11 == MaxRecursion {
						rules2[11] = "#42#31"
					}
				}
			}
			reg2 = "^" + reg2 + "$"

		} else {
			if phase == 1 {
				ruleId, rule := parseRule(sLine)
				rules[ruleId] = rule
				rules2[ruleId] = rule
			} else {
				matched, err := regexp.MatchString(reg, sLine)
				if err != nil {
					panic("error while executing regex")
				}
				if matched {
					ans1++
				}

				matched2, err := regexp.MatchString(reg2, sLine)
				if err != nil {
					panic("error while executing regex")
				}
				if matched2 {
					ans2++
				}
			}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func parseRule(line string) (int, string) {
	s := strings.Split(line, ":")
	id, err := strconv.Atoi(s[0])
	if err != nil {
		panic("error parsing rule id")
	}
	poss := strings.Split(strings.Trim(s[1], " "), "|")
	op := make ([]string, 0)
	for i := range poss {
		now := strings.Trim(poss[i], " \"")
		if now == "a" || now == "b" {
			op = append(op, now)
		} else {
			items := strings.Split(now, " ")
			for i := range items {
				items[i] = fmt.Sprintf("#%s", items[i])
			}
			op = append(op,fmt.Sprintf("(%s)", strings.Join(items, "")))
		}
	}
	return id, "(" + strings.Join(op, "|") + ")"
}

