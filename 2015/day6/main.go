package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instructions := GetInput()

	lightsPartOne := make([][1000]int, 1000)
	lightsPartTwo := make([][1000]int, 1000)
	for _, v := range instructions {
		instruct := v[0]
		a := v[1]
		b := v[2]
		c := v[3]
		d := v[4]

		for i := a; i <= c; i++ {
			for j := b; j <= d; j++ {
				if instruct == 2 {
					lightsPartOne[i][j] ^= 1
				} else {
					lightsPartOne[i][j] = instruct
				}

				switch instruct {
				case 0:
					if lightsPartTwo[i][j] > 0 {
						lightsPartTwo[i][j] -= 1
					}
				case 1:
					lightsPartTwo[i][j] += 1
				case 2:
					lightsPartTwo[i][j] += 2
				}
			}
		}
	}

	lightsPartOneCount := 0
	for _, v := range lightsPartOne {
		for _, y := range v {
			lightsPartOneCount += y
		}
	}

	lightsPartTwoCount := 0
	for _, v := range lightsPartTwo {
		for _, y := range v {
			lightsPartTwoCount += y
		}
	}

	fmt.Println(lightsPartOneCount)
	fmt.Println(lightsPartTwoCount)
}

func GetInput() [][5]int {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	InputFile, err := os.ReadFile(cwd + "/2015/day6/input.txt")
	if err != nil {
		panic(err)
	}

	inputInstruction := strings.Split(strings.TrimSpace(string(InputFile)), "\n")
	inputData := make([][5]int, len(inputInstruction))

	digit := regexp.MustCompile(`\d{1,3}`)
	for i, v := range inputInstruction {
		digitString := digit.FindAllString(v, -1)
		instructTokenOn, err := regexp.MatchString(`^turn on`, v)
		instructTokenOff, err := regexp.MatchString(`^turn off`, v)
		instructTokenToggle, err := regexp.MatchString(`^toggle`, v)

		for j, v := range digitString {
			k, err := strconv.Atoi(string(v))
			if err == nil {
				inputData[i][j+1] = k
			}
		}

		if instructTokenOff {
			inputData[i][0] = 0
		}

		if instructTokenOn {
			inputData[i][0] = 1
		}

		if instructTokenToggle {
			inputData[i][0] = 2
		}

		if err != nil {
			inputData[i][0] = 999
		}
	}

	return inputData
}
