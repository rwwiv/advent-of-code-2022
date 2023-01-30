package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Open file
	f, err := os.Open("./input")
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

	fmt.Println(elfCalories[0])
}
