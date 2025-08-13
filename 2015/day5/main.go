package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getInput() ([]string, error) {
	cwd, _ := os.Getwd()
	inputFile, err := os.ReadFile(cwd + "/2015/day5/input.txt")

	data := strings.Split(string(inputFile), "\n")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func CheckDoubleLetter(str string) bool {
	for i, l := range str {
		if i+1 == len(str) {
			return false
		}

		l := string(l)
		ln := string(str[i+1])
		if l == ln {
			return true
		}
	}

	return false
}

func CheckPairTwoLetter(str string) bool {
	pairIndexes := make(map[string][]int)

	for i := 0; i < len(str)-1; i++ {
		pair := str[i : i+2]
		pairIndexes[pair] = append(pairIndexes[pair], i)
	}

	for _, indexes := range pairIndexes {
		if len(indexes) >= 2 {
			for i := range indexes {
				for j := i + 1; j < len(indexes); j++ {
					if indexes[j]-indexes[i] >= 2 {
						return true
					}
				}
			}
		}
	}

	return false
}

func CheckOneRepeatOne(str string) bool {
	for i, l := range str {
		if i+2 == len(str) {
			return false
		}

		l := string(l)
		ln := string(str[i+2])

		if l == ln {
			return true
		}
	}

	return false
}

func NaughtyCheck(str string) string {
	vowelPattern := regexp.MustCompile(`[aeiou]`)
	Pattern := regexp.MustCompile(`ab|cd|pq|xy`)

	vowels := vowelPattern.FindAllString(str, -1)
	lice := Pattern.FindAllString(str, -1)

	isVowels := len(vowels) > 2
	isMulti := CheckDoubleLetter(str)
	isClean := len(lice) == 0

	if !(isVowels && isMulti && isClean) {
		return ""
	}

	return str
}

func NaughtyCheck_b(str string) string {
	isPairTwo := CheckPairTwoLetter(str)
	isRepeatOne := CheckOneRepeatOne(str)

	if !(isPairTwo && isRepeatOne) {
		return ""
	}

	return str
}

func main() {
	input, err := getInput()

	if err != nil {
		panic(err)
	}

	var niceStrings []string

	// This is for the part one that i have been solved
	// for _, v := range input {
	// 	str := NaughtyCheck(v)
	// 	if str != "" {
	// 		niceStrings = append(niceStrings, str)
	// 	}
	// }

	for _, v := range input {
		str := NaughtyCheck_b(v)
		if str != "" {
			niceStrings = append(niceStrings, str)
		}
	}

	fmt.Println(niceStrings)
	fmt.Println(len(niceStrings))
}
