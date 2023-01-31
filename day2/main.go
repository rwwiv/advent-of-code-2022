package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type action int

const (
	_ action = iota
	ROCK
	PAPER
	SCISSORS
)

var (
	oaMap = map[string]action{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}
	saMap = map[string]action{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}
	wins = map[action]action{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}
	loses = map[action]action{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}
)

func roundScorePt1(oa action, sa action) int {
	score := int(sa)
	if action(oa) == wins[action(sa)] {
		score += 6
	} else if sa == oa {
		score += 3
	} else {
		score += 0
	}
	return score
}

func roundSCorePt2(oa action, res string) int {
	score := 0
	switch res {
	case "X":
		score += 0
		score += int(wins[oa])
	case "Y":
		score += 3
		score += int(oa)
	case "Z":
		score += 6
		score += int(loses[oa])
	}
	return score
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

	pt1TotalScore, pt2TotalScore := 0, 0

	counter := 1
	for fs.Scan() {
		line := fs.Text()
		round := strings.Split(line, " ")
		sa, ok := saMap[round[1]]
		if !ok {
			log.Fatal("Player action not in map")
		}
		oa, ok := oaMap[round[0]]
		if !ok {
			log.Fatal("Opponent action not in map")
		}
		pt1TotalScore += roundScorePt1(oa, sa)
		pt2TotalScore += roundSCorePt2(oa, round[1])
		counter++
	}
	fmt.Printf("Pt 1 total score: %d\n", pt1TotalScore)

	/*
		Part 2
	*/

	fmt.Printf("Pt 2 total score: %d\n", pt2TotalScore)

}
