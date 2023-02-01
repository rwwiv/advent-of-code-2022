package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func uniqueRuneSlice(rs []rune, length int) bool {
	if len(rs) != length {
		return false
	}
	rm := make(map[rune]int, len(rs))
	for _, r := range rs {
		rm[r]++
	}
	for _, i := range rm {
		if i > 1 {
			return false
		}
	}
	return true
}

func main() {
	// Accept file name as arg
	filename := flag.String("f", "./input.txt", "Relative or absolute path to input file")
	flag.Parse()

	// Open file
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("unable to read input: %v", err)
	}
	defer f.Close()

	// Create buffer scanner
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanRunes)

	var ps []rune
	var ms []rune

	pCounter, mCounter := 1, 1
	pM, mM := false, false
	for fs.Scan() {
		if pM && mM {
			break
		}
		r := rune(fs.Text()[0])

		/* Part 1 */
		ps = append(ps, r)
		if len(ps) > 4 {
			ps = ps[1:]
		}

		if !uniqueRuneSlice(ps, 4) && !pM {
			pCounter++
		} else {
			pM = true
		}

		/* Part 2 */

		ms = append(ms, r)
		if len(ms) > 14 {
			ms = ms[1:]
		}

		if !uniqueRuneSlice(ms, 14) && !mM {
			mCounter++
		} else {
			mM = true
		}

	}
	/* Part 1 */
	fmt.Println(pCounter)

	/* Part 2 */
	fmt.Println(mCounter)
}
