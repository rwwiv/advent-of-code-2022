package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int64
}

type dir struct {
	parent *dir
	name   string
	size   int64
	dirs   []*dir
	files  []*file
}

func (d *dir) cd(dest string) (*dir, error) {
	nd := d
	if dest == "/" {
		for nd.name != "/" {
			nd = d.parent
		}
		return nd, nil
	}

	if dest == ".." {
		nd = d.parent
		return nd, nil
	}

	for _, chDir := range d.dirs {
		if chDir.name == dest {
			nd = chDir
			return nd, nil
		}
	}

	errMsg := fmt.Sprintf("dir %s does not contain dir %s", d.name, dest)

	log.Fatalf(errMsg)
	return nil, errors.New(errMsg)
}

func (d *dir) calcSize() int64 {
	var size int64 = 0
	for _, dir := range d.dirs {
		size += dir.calcSize()
	}
	for _, f := range d.files {
		size += f.size
	}
	d.size = size
	return size
}

func (d *dir) ls(r *bufio.Reader) {
	for {
		// Check if next line is a command
		nb, err := r.Peek(1)
		if err != nil {
			return
		}
		if rune(nb[0]) == '$' {
			break
		}

		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		strLine := string(line)
		contents := strings.Split(strLine, " ")
		if contents[0] == "dir" {
			chDir := &dir{d, contents[1], 0, []*dir{}, []*file{}}
			d.dirs = append(d.dirs, chDir)
			continue
		}
		size, err := strconv.ParseInt(contents[0], 10, 0)
		if err != nil {
			log.Fatalln("failed to read file")
		}
		f := &file{contents[1], size}
		d.files = append(d.files, f)
	}
}

func (d *dir) bigD() int64 {
	var sum int64 = 0
	for _, dir := range d.dirs {
		sum += dir.bigD()
	}
	if d.size <= 100_000 {
		sum += d.size
	}
	return sum
}

func (d *dir) allSizes() []int64 {
	size := []int64{}
	for _, dir := range d.dirs {
		size = append(size, dir.allSizes()...)
	}
	size = append(size, d.size)
	return size
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

	// Create buffer reader
	r := bufio.NewReader(f)

	pwd := &dir{
		nil, "/", 0, []*dir{}, []*file{},
	}

	for {
		token, _, err := r.ReadLine()
		if err != nil {
			break
		}
		line := string(token[:])
		input := strings.Split(line, " ")
		command := input[1]
		switch command {
		case "ls":
			pwd.ls(r)
		case "cd":
			pwd, err = pwd.cd(input[2])
			if err != nil {
				log.Fatalln("could not change dir")
			}
		}
	}

	root, err := pwd.cd("/")
	if err != nil {
		log.Fatalln("could not change dir")
	}
	root.calcSize()

	// Part 1
	sum := root.bigD()
	fmt.Printf("Part 1: %d\n", sum)

	// Part 2
	sizes := root.allSizes()
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] < sizes[j] })

	freeSpace := 70_000_000 - root.size
	for _, size := range sizes {
		if freeSpace+size >= 30_000_000 {
			fmt.Printf("Part 2: %d\n", size)
			break
		}
	}
}
