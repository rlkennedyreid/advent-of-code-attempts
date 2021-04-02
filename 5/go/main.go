package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

const inputFile = "../input.txt"

type BinarySpacePartitioner struct {
	Power int
}

func (partitioner BinarySpacePartitioner) initialRange() int {
	floatPower := float64(partitioner.Power)

	return int(math.Pow(2.0, floatPower))
}

func (partitioner BinarySpacePartitioner) indexFromTree(treeString string, key BinaryKey) int {

	tree := strings.Split(treeString, "")

	if len(tree) != partitioner.Power {
		log.Fatal("The tree string parsed into too many characters")
	}

	currentRange := partitioner.initialRange()
	currentMin := 0

	for i := 0; i < partitioner.Power; i++ {
		currentRange /= 2

		if tree[i] == key.upper {
			currentMin += currentRange
		}
	}

	return currentMin
}

type BinaryKey struct {
	lower string
	upper string
}

func readFileToStringSlice(fileDirectory string) []string {

	file, err := os.Open(fileDirectory)
	defer file.Close()

	if err != nil {
		log.Fatal("Could not open input file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func substringsAt(fullString string, pivot int) (string, string) {
	if pivot > len(fullString) || pivot < 0 {
		log.Fatal("Pivot point for substrings is out of bounds")
	}

	return fullString[:pivot], fullString[pivot:]
}

func splitSliceStrings(input []string, pivot int) ([]string, []string) {
	inputSize := len(input)
	leftSubstrings := make([]string, 0, inputSize)
	rightSubstrings := make([]string, 0, inputSize)

	for _, value := range input {
		left, right := substringsAt(value, pivot)

		leftSubstrings = append(leftSubstrings, left)
		rightSubstrings = append(rightSubstrings, right)
	}
	return leftSubstrings, rightSubstrings
}

func (partitioner BinarySpacePartitioner) getPartitionIndexesFrom(input []string, key BinaryKey) []int {
	indexSlice := make([]int, 0, len(input))

	for _, tree := range input {
		index := partitioner.indexFromTree(tree, key)
		indexSlice = append(indexSlice, index)
	}
	return indexSlice
}

func seatID(row, column int) int {
	return (row * 8) + column
}

func getSeatIDsFrom(rows, cols []int) []int {
	size := len(rows)
	seatIDs := make([]int, 0, size)

	for i := 0; i < size; i++ {
		seatIDs = append(seatIDs, seatID(rows[i], cols[i]))
	}

	return seatIDs
}

func main() {
	// Parse input into slice of rows
	rows := readFileToStringSlice(inputFile)

	// Parse each row into slice of substrings
	rowTrees, colTrees := splitSliceStrings(rows, 7)

	rowPartitioner := BinarySpacePartitioner{7}
	colPartitioner := BinarySpacePartitioner{3}

	// Iterate over slices, producing a slice of row numbers and columns
	rowNums := rowPartitioner.getPartitionIndexesFrom(rowTrees, BinaryKey{lower: "F", upper: "B"})
	colNums := colPartitioner.getPartitionIndexesFrom(colTrees, BinaryKey{lower: "L", upper: "R"})
	seatIDs := getSeatIDsFrom(rowNums, colNums)
	sort.Ints(seatIDs)
	// fmt.Println(seatIDs)
	for index, value := range seatIDs {
		if index != 0 {
			if value-seatIDs[index-1] != 1 {
				fmt.Println(seatIDs[index-1], seatIDs[index])
			}
		}
	}
	// From column and row numbers, produce a slice of seatIDs
}
