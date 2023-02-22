package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const topN int = 3

func main() {
	/*
		Part 1
	*/

	// Open file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("unable to read input: %v", err)
	}
	defer f.Close()

	// Create buffer scanner
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	// Create destination variables
	elfCalories := []int64{}
	var currentTotal int64 = 0

	// Scan file until EOF
	counter := 1
	for fs.Scan() {
		line := fs.Text()
		if line == "" {
			elfCalories = append(elfCalories, currentTotal)
			currentTotal = 0
		} else {
			num, err := strconv.ParseInt(line, 10, 0)
			if err != nil {
				log.Fatalf("unable to parse int at line %d: %v", counter, err)
			}
			currentTotal += num
		}
		counter++
	}
	// Append last total to calorie list
	elfCalories = append(elfCalories, currentTotal)

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	fmt.Printf("Part 1 answer: %d\n", elfCalories[0])

	/*
		Part 2
	*/

	var topNCalories int64 = 0
	for i := 0; i < topN; i++ {
		topNCalories += elfCalories[i]
	}

	fmt.Printf("Part 2 answer: %d\n", topNCalories)
}
