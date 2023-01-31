package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func containEither(start1 int64, end1 int64, start2 int64, end2 int64) bool {
	if start1 <= start2 && end1 >= end2 {
		return true
	}
	if start2 <= start1 && end2 >= end1 {
		return true
	}
	return false
}

func overlap(start1 int64, end1 int64, start2 int64, end2 int64) bool {
	if start1 <= start2 && end1 >= start2 {
		return true
	}
	if start2 <= start1 && end2 >= start1 {
		return true
	}
	if containEither(start1, end1, start2, end2) {
		return true
	}
	return false
}

func convertAssignmentSlice(aSlice []string) []int64 {
	slice := []int64{}
	for _, s := range aSlice {
		num, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Fatalf("could not parse int for pair %v", aSlice)
		}
		slice = append(slice, num)
	}
	return slice
}

func main() {
	// Accept file name as arg
	filename := flag.String("f", "./input", "Relative or absolute path to input file")
	flag.Parse()

	// Open file
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("unable to read input: %v", err)
	}
	defer f.Close()

	// Create buffer scanner
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	numFullContains, numOverlaps := 0, 0

	counter := 1
	for fs.Scan() {
		line := fs.Text()
		pair := strings.Split(line, ",")
		sar1 := strings.Split(pair[0], "-")
		sar2 := strings.Split(pair[1], "-")
		ar1 := convertAssignmentSlice(sar1)
		ar2 := convertAssignmentSlice(sar2)

		/* Part 1 */
		if containEither(ar1[0], ar1[1], ar2[0], ar2[1]) {
			numFullContains++
		}

		/* Part 2 */
		if overlap(ar1[0], ar1[1], ar2[0], ar2[1]) {
			numOverlaps++
		}
		counter++
	}
	/* Part 1 */
	fmt.Printf("Number of overlaps: %d\n", numFullContains)

	/* Part 2 */
	fmt.Printf("Number of overlaps: %d\n", numOverlaps)

}
