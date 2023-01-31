package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func firstSimilarPt1(s1 []string, s2 []string) (string, error) {
	ms1 := make(map[string]bool, len(s1))
	for _, ks1 := range s1 {
		ms1[ks1] = true
	}
	for _, ks2 := range s2 {
		if ms1[ks2] {
			return ks2, nil
		}
	}
	return "", errors.New("no similar runes")
}

func firstSimilarPt2(s1 []string, s2 []string, s3 []string) (string, error) {
	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}
	if len(s3) > maxLen {
		maxLen = len(s3)
	}

	m := make(map[string]int, maxLen)
	for _, ks1 := range s1 {
		m[ks1] = 1
	}
	for _, ks2 := range s2 {
		if m[ks2] == 1 {
			m[ks2]++
		}
	}
	for _, ks3 := range s3 {
		if m[ks3] == 2 {
			return ks3, nil
		}
	}
	return "", errors.New("no common rune")
}

func calcScore(s string) (int, error) {
	r := rune(s[0])
	charCode := int(r)
	if charCode >= 65 && charCode <= 90 {
		return charCode - 38, nil
	}
	if charCode >= 97 && charCode <= 122 {
		return charCode - 96, nil
	}
	return -1, errors.New("invalid rune")
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

	itemPriTotal, badgePriTotal := 0, 0
	group := [][]string{}

	counter := 1
	for fs.Scan() {
		line := fs.Text()
		sack := strings.Split(line, "")
		half := len(sack) / 2
		c1 := sack[:half]
		c2 := sack[half:]

		/* Part 1 */
		fsi, err := firstSimilarPt1(c1, c2)
		if err != nil {
			log.Fatal("failed to find match between compartments")
		}
		score, err := calcScore(fsi)
		if err != nil {
			log.Fatalf("could not parse rune from %s", fsi)
		}
		itemPriTotal += score

		/* Part 2 */
		group = append(group, sack)
		if len(group) == 3 {
			fsb, err := firstSimilarPt2(group[0], group[1], group[2])
			if err != nil {
				log.Fatal("failed to find group badge")
			}
			score, err = calcScore(fsb)
			if err != nil {
				log.Fatalf("could not parse rune from %s", fsb)
			}
			badgePriTotal += score
			group = group[:0]
		}

		counter++
	}

	/* Part 1 */
	fmt.Printf("Total item priorities: %d\n", itemPriTotal)

	/* Part 2 */
	fmt.Printf("Total badge priorities: %d\n", badgePriTotal)
}
