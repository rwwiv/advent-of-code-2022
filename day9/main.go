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

type tuple[T, U any] struct {
	first  T
	second U
}

func moveHead(h *tuple[int, int], t *tuple[int, int], d rune) bool {
	hasMoved := false
	switch d {
	case 'R':
		h.first++
	case 'L':
		h.first--
	case 'U':
		h.second++
	case 'D':
		h.second--
	}
	// horiz
	if h.first > t.first+1 {
		t.first++
		hasMoved = true

		if h.second > t.second {
			t.second++
		}
		if h.second < t.second {
			t.second--
		}
	}

	//TODO: Make all like ^^^^

	if h.first < t.first {
		if h.second > t.second+1 {
			t.second++
			t.first--
			hasMoved = true
		}
		if h.second < t.second-1 {
			t.second--
			t.first--
			hasMoved = true
		}
		if h.first < t.first-1 {
			t.first--
			hasMoved = true
		}
	}
	// vert
	if h.second > t.second {
		if h.first > t.first+1 {
			t.first++
			t.second++
			hasMoved = true
		}
		if h.first < t.first-1 {
			t.first--
			t.second++
			hasMoved = true
		}
		if h.second > t.second+1 {
			t.second++
			hasMoved = true
		}
	}
	if h.second < t.second {
		if h.first > t.first+1 {
			t.first++
			t.second--
			hasMoved = true
		}
		if h.first < t.first-1 {
			t.first--
			t.second--
			hasMoved = true
		}
		if h.second < t.second-1 {
			t.second--
			hasMoved = true
		}
	}
	fmt.Printf("H: %v, T: %v, move: %v\n", h, t, hasMoved)
	return hasMoved
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

	h := &tuple[int, int]{0, 0}
	t := &tuple[int, int]{0, 0}
	tPos := 0
	counter := 0
	for fs.Scan() {
		if counter == 10 {
			break
		}
		line := fs.Text()
		fmt.Printf("---- %s ----\n", line)
		motion := strings.Split(line, " ")
		dir := motion[0]
		num, err := strconv.ParseInt(motion[1], 10, 0)
		if err != nil {
			log.Fatalln("invalid instruction")
		}
		tmpSum := 0
		for i := 0; i < int(num); i++ {
			if moveHead(h, t, rune(dir[0])) {
				tmpSum++
			}
		}
		fmt.Printf("  - Num spaces: %d\n", tmpSum)

		tPos += tmpSum
		fmt.Println()
		counter++
	}

	fmt.Println(tPos)
}
