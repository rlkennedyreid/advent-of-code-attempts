package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INPUT_FILE = "../input.txt"
const SUBSTR_PIVOT = 7

var (
	ROW_MAPPING = BinaryKey{lower: "F", upper: "B"}
	COL_MAPPING = BinaryKey{lower: "L", upper: "R"}
)

type BinaryKey struct {
	lower string
	upper string
}

func main() {
	inputLines := readFileToStringSlice(INPUT_FILE)

	rowSequences, colSequences := splitSliceStrings(inputLines, SUBSTR_PIVOT)

	rowIDs := MapSequencesToDecimals(rowSequences, ROW_MAPPING)
	colIDs := MapSequencesToDecimals(colSequences, COL_MAPPING)

	seatIDs := getSeatIDsFrom(rowIDs, colIDs)

	sort.Ints(seatIDs)

	getAnswersAndPrintOutputMessage(seatIDs)
}

func readFileToStringSlice(fileDirectory string) []string {

	file, err := os.Open(fileDirectory)

	if err != nil {
		log.Fatal("Could not open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
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

func substringsAt(fullString string, pivot int) (string, string) {
	if pivot > len(fullString) || pivot < 0 {
		log.Fatal("Pivot point for substrings is out of bounds")
	}

	return fullString[:pivot], fullString[pivot:]
}

func MapSequencesToDecimals(input []string, key BinaryKey) []int {

	binarySequences := MapToBinarySequenceSlice(input, key)

	decimalSlice := BinaryStringSliceToInts(binarySequences)

	return decimalSlice
}

func MapToBinarySequenceSlice(input []string, key BinaryKey) []string {
	output := make([]string, 0, len(input))

	for _, value := range input {
		output = append(output, CharSequenceToBinarySequence(value, key))
	}

	return output
}

func CharSequenceToBinarySequence(input string, key BinaryKey) string {
	var output string

	temp := strings.Replace(input, key.upper, "1", -1)
	temp = strings.Replace(temp, key.lower, "0", -1)

	output = temp

	return output
}

func BinaryStringSliceToInts(input []string) []int {
	inputSize := len(input)
	intSlice := make([]int, 0, inputSize)

	for _, value := range input {
		parsedInt, err := strconv.ParseInt(value, 2, 64)
		if err != nil {
			log.Fatal()
		}
		intSlice = append(intSlice, int(parsedInt))
	}

	return intSlice
}

func getSeatIDsFrom(rows, cols []int) []int {
	size := len(rows)
	seatIDs := make([]int, 0, size)

	for i := 0; i < size; i++ {
		seatIDs = append(seatIDs, seatID(rows[i], cols[i]))
	}

	return seatIDs
}

func seatID(row, column int) int {
	return (row * 8) + column
}

func getAnswersAndPrintOutputMessage(seatIDs []int) {

	fmt.Printf("Seat IDs: Min, Max = %v, %v\n", seatIDs[0], seatIDs[len(seatIDs)-1])

	for index, value := range seatIDs {
		if index > 0 {
			if value-seatIDs[index-1] != 1 {
				fmt.Printf("Missing Seat ID between %v and %v\n", seatIDs[index-1], seatIDs[index])
			}
		}
	}
}
