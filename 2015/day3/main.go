package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Log []int
type Logs []Log

func modify_log(log Log, direction string) Log {
	switch direction {
	case "^":
		log[0] = log[0] + 1
	case "v":
		log[0] = log[0] - 1
	case ">":
		log[1] = log[1] + 1
	case "<":
		log[1] = log[1] - 1
	}

	copyLog := make(Log, len(log))
	copy(copyLog, log)

	return copyLog
}

func data_fision(log_marks []string) ([]string, []string) {
	var directions_a []string
	var directions_b []string

	for i, log_mark := range log_marks {
		if i%2 == 0 {
			directions_a = append(directions_a, log_mark)
		} else {
			directions_b = append(directions_b, log_mark)
		}
	}

	return directions_a, directions_b
}

func main() {
	dat, err := os.ReadFile("/media/Datacenter/CS/adventofcode-marathon/2015/day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(dat), "")
	data_a, data_b := data_fision(data)

	current_log := Log{0, 0}
	current_log_a := Log{0, 0}
	current_log_b := Log{0, 0}

	var data_log Logs
	var data_log_a Logs
	var data_log_b Logs

	for _, v := range data {
		data_log = append(data_log, modify_log(current_log, v))
	}
	for _, v := range data_a {
		data_log_a = append(data_log_a, modify_log(current_log_a, v))
	}
	for _, v := range data_b {
		data_log_b = append(data_log_b, modify_log(current_log_b, v))
	}

	slices.SortFunc(data_log, func(a, b Log) int {
		return slices.Compare(a, b)
	})
	slices.SortFunc(data_log_a, func(a, b Log) int {
		return slices.Compare(a, b)
	})
	slices.SortFunc(data_log_b, func(a, b Log) int {
		return slices.Compare(a, b)
	})

	data_log = slices.CompactFunc(data_log, slices.Equal)
	data_log_a = slices.CompactFunc(data_log_a, slices.Equal)
	data_log_b = slices.CompactFunc(data_log_b, slices.Equal)

	fmt.Println(len(data_log) + 1)

	data_log_ab := slices.Concat(data_log_a, data_log_b)
	slices.SortFunc(data_log_ab, func(a, b Log) int {
		return slices.Compare(a, b)
	})
	data_log_ab = slices.CompactFunc(data_log_ab, slices.Equal)

	fmt.Println(len(data_log_ab))
}
