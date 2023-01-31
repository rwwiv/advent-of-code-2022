package main

import (
	"bufio"
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
)

func roundScore(sa action, oa action) int {
	var wins = map[action]action{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}

	score := int(sa)
	defeat := wins[action(sa)]
	if action(oa) == defeat {
		score += 6
	} else if sa == oa {
		score += 3
	} else {
		score += 0
	}
	return score
}

func main() {
	/*
		Part 1
	*/

	// Open file
	f, err := os.Open("./input")
	if err != nil {
		log.Fatalf("unable to read input: %v", err)
	}
	defer f.Close()

	// Create buffer scanner
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	totalScore := 0

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
		totalScore += roundScore(sa, oa)
		counter++
	}
	fmt.Printf("Total score: %d\n", totalScore)
}
