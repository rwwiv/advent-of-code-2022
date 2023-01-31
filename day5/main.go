package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseCrateSetup(buff []string) [][]rune {
	stacks := make([][]rune, 0)
	sbuff := buff[:len(buff)-1]

	for _, line := range sbuff {
		linebuf := []rune(line)

		counter := 0
		for i := 1; i < len(linebuf); i += 4 {
			if len(stacks) < counter+1 {
				stacks = append(stacks, make([]rune, 0))
			}
			if string(linebuf[i]) != " " {
				stacks[counter] = append(stacks[counter], linebuf[i])
			}
			counter++
		}
	}
	return stacks
}

func parseCommand(command string) (int, int, int, error) {
	linebuf := strings.Split(command, " ")
	// num, source, dest := , linebuf[3], linebuf[5]
	num, err := strconv.ParseInt(linebuf[1], 10, 0)
	if err != nil {
		return -1, -1, -1, errors.New("could not parse crate num")
	}
	source, err := strconv.ParseInt(linebuf[3], 10, 0)
	if err != nil {
		return -1, -1, -1, errors.New("could not parse source")
	}
	dest, err := strconv.ParseInt(linebuf[5], 10, 0)
	if err != nil {
		return -1, -1, -1, errors.New("could not parse destination")
	}
	return int(num), int(source - 1), int(dest - 1), nil
}

func moveCratesPt1(buff []string, stacks [][]rune) string {
	si := make([][]rune, len(stacks))
	copy(si, stacks)
	out := ""
	for index, line := range buff {
		num, source, dest, err := parseCommand(line)
		if err != nil {
			log.Fatalf("could not parse command %d", index+1)
		}
		for i := 0; i < num; i++ {
			tmp := si[source][0]
			si[source] = si[source][1:]
			si[dest] = append([]rune{tmp}, si[dest]...)
		}
	}
	for _, ra := range si {
		out += string(ra[0])
	}
	return out
}

func moveCratesPt2(buff []string, stacks [][]rune) string {
	si := make([][]rune, len(stacks))
	copy(si, stacks)
	out := ""
	for index, line := range buff {
		num, source, dest, err := parseCommand(line)
		if err != nil {
			log.Fatalf("could not parse command %d", index+1)
		}
		tmp := make([]rune, num)
		copy(tmp, si[source][:num])
		si[source] = si[source][num:]
		si[dest] = append(tmp, si[dest]...)
	}
	for _, ra := range si {
		out += string(ra[0])
	}
	return out
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

	var buff []string
	var stacks [][]rune
	for fs.Scan() {
		line := fs.Text()
		if len(line) == 0 {
			stacks = parseCrateSetup(buff)
			buff = buff[:0]
		}
		if len(line) != 0 {
			buff = append(buff, line)
		}
	}
	fmt.Printf("Part 1 answer: %s\n", moveCratesPt1(buff, stacks))
	fmt.Printf("Part 2 answer: %s\n", moveCratesPt2(buff, stacks))

}
