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

func firstSimilar(s1 []string, s2 []string) (string, error) {
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

func calcScore(s string) (int, error) {
	r := rune(s[0])
	charCode := int(r)
	if charCode >= 65 && charCode <= 90 {
		return charCode - 38, nil
	} else if charCode >= 97 && charCode <= 122 {
		return charCode - 96, nil
	} else {
		return -1, errors.New("invalid rune")
	}
}

func main() {
	/*
		Part 1
	*/

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

	totalPriorities := 0

	counter := 1
	for fs.Scan() {
		line := fs.Text()
		sack := strings.Split(line, "")
		half := len(sack) / 2
		c1 := sack[:half]
		c2 := sack[half:]
		fsm, err := firstSimilar(c1, c2)
		if err != nil {
			log.Fatal("Failed to find match between compartments")
		}
		score, err := calcScore(fsm)
		if err != nil {
			log.Fatalf("Could not parse rune from %s", fsm)
		}
		totalPriorities += score
		counter++
	}
	fmt.Printf("Total priorities: %d\n", totalPriorities)
}
