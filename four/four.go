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

	input, err := os.Open("four/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	passportInfo := ""

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
		if sLine == "" {

			fields := strings.Split(strings.Trim(passportInfo, " "), " ")
			keys := make(map [string] string)
			for _, field := range fields {
				keyValues := strings.Split(field, ":")
				keys[keyValues[0]] = keyValues[1]
			}

			if len(keys) == 8 || (len(keys) == 7 && keys["cid"] == "") {
				ans1++
				ok := true
				if byr, err := strconv.Atoi(keys["byr"]); err != nil || byr < 1920 || byr > 2002 || keys["byr"][0]=='0'{
					ok = false
				}
				if iyr, err := strconv.Atoi(keys["iyr"]); err != nil || iyr < 2010 || iyr > 2020 || keys["iyr"][0]=='0'{
					ok = false
				}
				if eyr, err := strconv.Atoi(keys["eyr"]); err != nil || eyr < 2020 || eyr > 2030 || keys["eyr"][0]=='0'{
					ok = false
				}
				switch {
				case strings.HasSuffix(keys["hgt"], "cm"):
					{
						pref := strings.TrimSuffix(keys["hgt"], "cm")
						if v, err := strconv.Atoi(pref); err != nil || v < 150 || v > 193 || keys["hgt"][0]=='0'{
							ok = false
						}
					}
				case strings.HasSuffix(keys["hgt"], "in"):
					{
						pref := strings.TrimSuffix(keys["hgt"], "in")
						if v, err := strconv.Atoi(pref); err != nil || v < 59 || v > 76 || keys["hgt"][0]=='0'{
							ok = false
						}
					}
				default:
					{
						ok = false
					}
				}
				if m,err := regexp.MatchString("^#[0-9a-f]{6}$", keys["hcl"]); err!=nil || !m {
					ok = false
				}
				if m,err := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", keys["ecl"]); err!=nil || !m {
					ok = false
				}
				if m,err := regexp.MatchString("^[0-9]{9}$", keys["pid"]); err!=nil || !m {
					ok = false
				}
				if ok {
					ans2++
				}
			}
			passportInfo = ""
		} else {
			passportInfo = passportInfo + " " + sLine
		}
	}


	fmt.Println(ans1)
	fmt.Println(ans2)
}



