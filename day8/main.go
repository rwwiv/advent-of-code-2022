package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type tree struct {
	h int
	// pt1
	vt bool
	vb bool
	vr bool
	vl bool
	// pt2
	st int
	sb int
	sl int
	sr int
}

func setVis(trees [][]*tree, outerIndex int, innerIndex int) bool {
	t := trees[outerIndex][innerIndex]

	if outerIndex == 0 || innerIndex == 0 {
		return true
	}

	if outerIndex == len(trees)-1 || innerIndex == len(trees[outerIndex])-1 {
		return true
	}

	// vt
	for i := outerIndex - 1; i >= 0; i-- {
		t.st++
		comp := trees[i][innerIndex].h
		if comp >= t.h {
			t.vt = false
			break
		}
	}
	// vb
	for i := outerIndex + 1; i < len(trees); i++ {
		t.sb++
		comp := trees[i][innerIndex].h
		if comp >= t.h {
			t.vb = false
			break
		}
	}
	// vl
	for i := innerIndex - 1; i >= 0; i-- {
		t.sl++
		comp := trees[outerIndex][i].h
		if comp >= t.h {
			t.vl = false
			break
		}
	}
	// vr
	for i := innerIndex + 1; i < len(trees[outerIndex]); i++ {
		t.sr++
		comp := trees[outerIndex][i].h
		if comp >= t.h {
			t.vr = false
			break
		}
	}

	return t.vt || t.vb || t.vl || t.vr
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
	fs.Split(bufio.ScanLines)

	treeMap := [][]*tree{}
	for fs.Scan() {
		treeLine := []*tree{}
		line := fs.Text()
		for _, r := range line {
			height := int(r - '0')
			t := &tree{height, true, true, true, true, 0, 0, 0, 0}
			treeLine = append(treeLine, t)
		}
		treeMap = append(treeMap, treeLine)
	}

	// Part 1
	allVis := 0

	// Part 2
	maxScore := 0

	for i, tl := range treeMap {
		for j := range tl {
			if setVis(treeMap, i, j) {
				allVis++
			}
			t := treeMap[i][j]
			score := t.st * t.sb * t.sl * t.sr

			if score > maxScore {
				maxScore = score
			}
		}
	}

	// Part 1
	fmt.Println(allVis)

	// Part 2
	fmt.Println(maxScore)
}
