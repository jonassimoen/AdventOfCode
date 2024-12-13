package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Input file: %s\n", *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	out := part1(ls)
	fmt.Printf("Output P1: %d\n", out)
	out = part2(ls)
	fmt.Printf("Output P2: %d\n", out)
}

type Block struct {
	id    int // -1 if free space
	size  int
	moved bool
}

func parseInput(l string) []Block {
	var blocks []Block
	idx := 0
	for ii, c := range l {
		v := int(c - '0')
		blockId := -1
		if ii%2 == 0 {
			blockId = idx
			idx++
		}
		if v != 0 {
			block := Block{
				blockId, v, false,
			}

			blocks = append(blocks, block)
		}
	}
	return blocks
}

func printBlocks(blocks []Block, spaces bool) {
	for _, b := range blocks {
		for i := 0; i < b.size; i++ {
			if b.id == -1 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", b.id)
			}
		}
		if spaces {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func findFreeSpaces(blocks []Block) int {
	for idx, b := range blocks {
		if b.id == -1 {
			return idx
		}
	}
	return -1
}

func replaceSingleFileBlock(blocks []Block) []Block {
	idxSpace := findFreeSpaces(blocks)
	idxLastBlock := len(blocks) - 1
	for idx := len(blocks) - 1; idx > idxSpace; idx-- {
		if blocks[idx].id != -1 {
			idxLastBlock = idx
			break
		}
	}
	fileBlock := blocks[idxLastBlock]
	spaceBlock := blocks[idxSpace]

	blocksNew := []Block{}
	if fileBlock.size < spaceBlock.size {
		// Fileblock kan slechts deel opvullen van spaceblock
		blocks[idxSpace].size -= fileBlock.size
		blocksNew = append(blocksNew, blocks[:idxSpace]...)
		blocksNew = append(blocksNew, Block{fileBlock.id, fileBlock.size, false})
		blocksNew = append(blocksNew, Block{-1, spaceBlock.size - fileBlock.size, false})
		if idxSpace+1 < len(blocks) {
			blocksNew = append(blocksNew, blocks[idxSpace+1:idxLastBlock]...)
		}
		blocksNew = append(blocksNew, Block{-1, fileBlock.size, false})
		return blocksNew
	} else {
		blocks[idxLastBlock].size -= spaceBlock.size
		if blocks[idxLastBlock].size == 0 {
			blocksNew = append(blocksNew, blocks[:idxLastBlock]...)
			blocksNew = append(blocksNew, blocks[idxLastBlock+1:]...)
		} else {
			blocksNew = append(blocksNew, blocks...)
		}
		blocksNew[idxSpace].id = fileBlock.id
		blocksNew = append(blocksNew, Block{-1, spaceBlock.size, false})
		return blocksNew
	}
}

func mergeFreeSpaces(blocks []Block) []Block {
	freeSpaceSize := 0
	newBlocks := []Block{}
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].id == -1 {
			freeSpaceSize += blocks[i].size
		} else {
			newBlocks = append(newBlocks, blocks[:i+1]...)
			newBlocks = append(newBlocks, Block{-1, freeSpaceSize, false})
			return newBlocks
		}
	}
	return append(newBlocks, blocks...)
}

func calculateChecksum(blocks []Block) int {
	sum := 0
	idx := 0
	for _, block := range blocks {
		for i := 0; i < block.size; i++ {
			if block.id != -1 {
				sum += ((idx + i) * block.id)
			}
		}
		idx += block.size
	}
	return sum
}

func part1(ls []string) int {
	blocks := parseInput(ls[0])

	for findFreeSpaces(blocks) != (len(blocks) - 1) {
		blocks = replaceSingleFileBlock(blocks)
		//printBlocks(blocks, true)
		blocks = mergeFreeSpaces(blocks)
		//printBlocks(blocks, true)
		//fmt.Println()
	}
	return calculateChecksum(blocks)
}

func replaceSingleFileBlockFull(blocks []Block) []Block {
	i := len(blocks) - 1
	for i >= 0 {
		block := blocks[i]
		if (blocks[i].id == -1) || block.moved {
			i--
			continue
		}
		moved := false
		for j := 0; j < i; j++ {
			if blocks[j].id == -1 {
				spaceBlock := blocks[j]
				fileBlock := blocks[i]

				//fmt.Printf("\tempty space block of length %d\n", blocks[j].size)
				if fileBlock.size <= spaceBlock.size {
					moved = true
					// On space location: |...| ==> |xx.| or |xxx|
					blocks[j] = Block{fileBlock.id, fileBlock.size, true}
					// Old FileBlock becames fully SpaceBlock
					blocks[i] = Block{-1, fileBlock.size, false}

					// Creating new block sequence
					var newBlocks []Block
					// All blocks until new filled are kept...
					newBlocks = append(newBlocks, blocks[:j+1]...)
					// If SpaceBlock was not fully filled by FileBlock (thus FileSize < SpaceSize)
					//    a block with the remaining SpaceSize needs to be created
					if fileBlock.size < spaceBlock.size {
						newBlocks = append(newBlocks, Block{-1, spaceBlock.size - fileBlock.size, false})
					}
					// All blocks after new SpaceBlock need to be added
					newBlocks = append(newBlocks, blocks[j+1:]...)

					blocks = newBlocks
					break
				}
			}
		}
		if !moved {
			i--
		}
	}
	return blocks
}

func part2(ls []string) int {
	blocks := parseInput(ls[0])
	blocks = replaceSingleFileBlockFull(blocks)

	return calculateChecksum(blocks)
}
