package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ls []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ls = append(ls, sc.Text())
	}
	return ls, sc.Err()
}

func main() {
	part1ptr := flag.Int("p", 1, "Part")
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Part %d - Input file: %s\n", *part1ptr, *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	var out int
	if *part1ptr == 1 {
		out = part1(ls)
	} else {
		out = part2(ls)
	}
	fmt.Printf("Output: %d\n", out)
}

type Directory struct {
	parent     *Directory
	name       string
	childDirs  []*Directory
	childFiles []*File
	size       int
}

type File struct {
	parent *Directory
	name   string
	size   int
}

func cd(parent *Directory, name string) *Directory {
	for _, dir := range parent.childDirs {
		if dir.name == name {
			return dir
		}
	}
	newDir := Directory{parent: parent, name: name}
	parent.childDirs = append(parent.childDirs, &newDir)
	return &newDir
}

func addFile(parent *Directory, name string, size int) *Directory {
	newFile := File{parent: parent, name: name, size: size}
	parent.childFiles = append(parent.childFiles, &newFile)
	recalculateSize(parent)
	return parent
}
func addDir(parent *Directory, name string) *Directory {
	newDir := Directory{parent: parent, name: name}
	parent.childDirs = append(parent.childDirs, &newDir)
	return parent
}

func recalculateSize(d *Directory) {
	for d.parent != nil {
		size := 0
		for _, child := range d.childFiles {
			size += child.size
		}
		for _, dir := range d.childDirs {
			size += dir.size
		}
		d.size = size
		d = d.parent
	}
}

func printTree(parent *Directory, index int) {
	for _ = range index {
		fmt.Printf("   ")
	}
	fmt.Printf("- %s (dir) [%d]\n", parent.name, parent.size)
	for _, childD := range parent.childDirs {
		printTree(childD, index+1)
	}
	for _, childF := range parent.childFiles {
		for _ = range index + 1 {
			fmt.Printf("   ")
		}
		fmt.Printf("- %s (file, size=%d)\n", childF.name, childF.size)
	}
}

func decode(ls []string) *Directory {
	homeDir := &Directory{name: "START"}
	currentDir := homeDir
	addFileToCurrentDir := false
	for _, l := range ls {
		prefix, prefixFound := strings.CutPrefix(l, "$ ")
		prefixSplit := strings.Split(prefix, " ")
		if prefixFound && prefixSplit[0] == "cd" {
			if prefixSplit[1] == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = cd(currentDir, prefixSplit[1])
			}
		} else if prefixFound && prefixSplit[0] == "ls" {
			addFileToCurrentDir = true
		} else if addFileToCurrentDir {
			lSplit := strings.Split(l, " ")
			if lSplit[0] == "dir" {
				currentDir = addDir(currentDir, lSplit[1])
			} else {
				size, _ := strconv.Atoi(lSplit[0])
				currentDir = addFile(currentDir, lSplit[1], size)
			}
		}
	}
	return homeDir.childDirs[0]
}

func calculateSizes(dir *Directory) int {
	if len(dir.childDirs) == 0 {
		if dir.size > 100000 {
			return 0
		}
		return dir.size
	}
	size := 0
	for _, child := range dir.childDirs {
		size += calculateSizes(child)
	}
	if dir.size > 100000 {
		return size
	}
	return dir.size + size
}

func part1(ls []string) int {
	d := decode(ls)
	//printTree(d, 0)
	a := calculateSizes(d)
	return a
}

func dirSizes(d *Directory, sizes []int) []int {
	sizes = append(sizes, d.size)
	for _, dir := range d.childDirs {
		sizes = dirSizes(dir, sizes)
	}
	return sizes
}

func findRightSize(d *Directory, sizes []int) int {
	unused := 70000000 - d.size
	for _, s := range sizes {
		fmt.Println(unused, s, "=", unused+s > 30000000)
		if unused+s > 30000000 {
			return s
		}
	}
	return 0
}

func part2(ls []string) int {
	d := decode(ls)
	s := dirSizes(d, []int{})
	slices.Sort(s)
	return findRightSize(d, s)
}
