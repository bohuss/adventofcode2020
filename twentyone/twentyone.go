package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	input, err := os.Open("twentyone/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := make([]string, 0)
	splitter := " (contains "

	allAllergens := make( map[string] bool)
	allAllergensList := make( []string, 0)
	allIngredients := make( map[string] bool)
	allergensMap := make(map[string] map[string] bool)

	lines := make([][]string, 0)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		s := strings.Split(string(line), splitter)
		lines = append(lines, s)

		ingredients := strings.Split(s[0], " ")
		allergens := strings.Split(s[1][:len(s[1]) - 1], ", ")
		for _, i := range ingredients {
			allIngredients[i] = true
		}
		for _, a := range allergens {
			allAllergens[a] = true
			if allergensMap[a] == nil {
				allergensMap[a] = make(map[string] bool)
				allAllergensList = append(allAllergensList, a)
			}
			for _, i := range ingredients {
				allergensMap[a][i] = true
			}
		}

	}

	for _, s := range lines {
		ingredients := make (map[string] bool)
		for _, i := range strings.Split(s[0], " ") {
			ingredients[i] = true
		}
		allergens := strings.Split(s[1][:len(s[1]) - 1], ", ")

		for _, a := range allergens {
			for i := range allergensMap[a] {
				if ingredients[i] != true {
					delete(allergensMap[a], i)
				}
			}
		}
	}

	change := true
	done := make(map[string] bool)
	badIngredients := make(map[string] bool)

	for change {
		change = false
		for a, v := range allergensMap {
			if len(v) > 1 || done[a]{
				continue
			}
			change = true
			done[a] = true
			identifiedIngredient := single(v)
			badIngredients[identifiedIngredient] = true
			for a2, v2 := range allergensMap {
				if a == a2 {
					continue
				}
				if v2[identifiedIngredient] {
					delete(v2, identifiedIngredient)
				}
			}
		}
	}

	for _, s := range lines {
		for _, i := range strings.Split(s[0], " ") {
			if badIngredients[i] != true {
				ans1++
			}
		}
	}
	sort.Strings(allAllergensList)
	for _, a := range allAllergensList {
		ans2 = append(ans2, single(allergensMap[a]))
	}

	fmt.Println(ans1)
	fmt.Println(strings.Join(ans2, ","))
}

func single(m map[string]bool) string { // TODO: is this the best way to get single key from map in GO?
	for k := range m {
		return k
	}
	return ""
}
