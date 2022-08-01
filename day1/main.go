package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_U64 = uint64(18446744073709551615)
const INPUT_PATH = "input.txt"

func parseInput(inputPath string) []uint64 {
	var inputDepths []uint64

	input, err := os.Open(inputPath)
	if err != nil {
		panic("Input file not found")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		nextLine := scanner.Text()

		newDepth, err := strconv.ParseUint(nextLine, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Cannot parse value '%s'", nextLine))
		}

		inputDepths = append(inputDepths, newDepth)
	}

	return inputDepths
}

func measureDepthIncrease(depths []uint64) uint64 {
	lastDepth := MAX_U64
	var depthIncreaseCounter uint64

	for _, depth := range depths {
		if depth > lastDepth {
			depthIncreaseCounter++
		}
		lastDepth = depth
	}

	return depthIncreaseCounter
}

func measureDepthSlidingWindow(depths []uint64, width int) uint64 {
	lastDepth := MAX_U64
	var depthIncreaseCounter uint64

	for i := 0; i <= len(depths)-width; i++ {
		var sum uint64
		for _, v := range depths[i : i+width] {
			sum += v
		}
		if sum > lastDepth {
			depthIncreaseCounter++
		}
		lastDepth = sum
	}

	return depthIncreaseCounter
}

func main() {
	fmt.Print(
		"--- Advent of Code 2021 ---\n",
		"          Day   1          \n",
		"                           \n",
		"          Part  A          \n")

	inputDepths := parseInput(INPUT_PATH)
	fmt.Printf("Measurements larger than the previous one: %d\n", measureDepthIncrease(inputDepths))

	fmt.Printf("\n          Part  B          \n")
	fmt.Printf("Depth measurements from sliding window: %d\n", measureDepthSlidingWindow(inputDepths, 3))

	fmt.Printf("\n---------------------------\n")
}
