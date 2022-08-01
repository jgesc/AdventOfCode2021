package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

const INPUT_PATH = "input.txt"
const DIAGNOSTIC_NUMBER_LENGTH = 12

type BitCriteria uint

const (
	BIT_CRITERIA_O2   = BitCriteria(iota)
	BIT_CRITERIA_CO2  = BitCriteria(iota)
	BIT_CRITERIA_NONE = BitCriteria(iota)
)

func parseInput(inputPath string) []uint16 {
	input, err := os.Open(inputPath)
	if err != nil {
		panic("Input file not found")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var diagnosticNumbers []uint16
	for scanner.Scan() {
		nextLine := scanner.Text()

		numberRaw, err := strconv.ParseUint(nextLine, 2, DIAGNOSTIC_NUMBER_LENGTH)
		if err != nil {
			panic(fmt.Sprintf("Could not parse line '%s'", nextLine))
		}

		number := bits.Reverse16(uint16(numberRaw)) >> (16 - DIAGNOSTIC_NUMBER_LENGTH)
		diagnosticNumbers = append(diagnosticNumbers, number)
	}

	return diagnosticNumbers
}

func getBit(number uint16, bit uint) uint16 {
	return (number >> bit) & 1
}

func calculateMostCommonBitCriteria(diagnosticNumbers []uint16, bit uint, criteria BitCriteria) uint {
	var counter uint16
	for _, number := range diagnosticNumbers {
		counter += getBit(number, bit)
	}

	elementCount := len(diagnosticNumbers)
	halfElementCount := uint16(elementCount / 2)
	if criteria != BIT_CRITERIA_NONE && counter == halfElementCount && elementCount%2 == 0 {
		switch criteria {
		case BIT_CRITERIA_O2:
			return 1
		case BIT_CRITERIA_CO2:
			return 1
		}
	}
	if counter > halfElementCount {
		return 1
	} else {
		return 0
	}
}

func calculateMostCommonBit(diagnosticNumbers []uint16, bit uint) uint {
	return calculateMostCommonBitCriteria(diagnosticNumbers, bit, BIT_CRITERIA_NONE)
}

func calculateGamma(diagnosticNumbers []uint16) uint {
	var gamma uint
	for bit := uint(0); bit < DIAGNOSTIC_NUMBER_LENGTH; bit++ {
		gamma = (gamma << 1) | calculateMostCommonBit(diagnosticNumbers, bit)
	}
	return gamma
}

func calculateEpsilon(gamma uint) uint {
	return (^gamma) << (bits.Len(^gamma) - DIAGNOSTIC_NUMBER_LENGTH) >> (bits.Len(^gamma) - DIAGNOSTIC_NUMBER_LENGTH)
}

func filterNumbers(diagnosticNumbers []uint16, bitCriteria BitCriteria) uint {
	remainingNumbers := make([]uint16, len(diagnosticNumbers))

	copy(remainingNumbers, diagnosticNumbers)
	for bit := uint(0); len(remainingNumbers) > 1; bit++ {
		mostCommonBit := calculateMostCommonBitCriteria(remainingNumbers, bit, bitCriteria)

		var filteredNumbers []uint16
		for _, number := range remainingNumbers {
			if bitCriteria == BIT_CRITERIA_O2 && getBit(number, bit) == uint16(mostCommonBit) {
				filteredNumbers = append(filteredNumbers, number)
			}
			if bitCriteria == BIT_CRITERIA_CO2 && getBit(number, bit) != uint16(mostCommonBit) {
				filteredNumbers = append(filteredNumbers, number)
			}
		}

		remainingNumbers = filteredNumbers
	}
	fmt.Printf("%016b\n", bits.Reverse16(remainingNumbers[0]))
	return uint(bits.Reverse16(remainingNumbers[0]) >> (16 - DIAGNOSTIC_NUMBER_LENGTH))
}

func main() {
	fmt.Print(
		"--- Advent of Code 2021 ---\n",
		"          Day   2          \n",
		"                           \n",
		"          Part  A          \n")

	diagnosticNumbers := parseInput(INPUT_PATH)

	gamma := calculateGamma(diagnosticNumbers)
	epsilon := calculateEpsilon(gamma)
	fmt.Printf("Gamma = %d, Epsilon = %d\n", gamma, epsilon)
	fmt.Printf("Gamma * Epsilon = %d\n", gamma*epsilon)

	fmt.Printf("\n          Part  B          \n")
	oxygenGeneratorRating := filterNumbers(diagnosticNumbers, BIT_CRITERIA_O2)
	co2ScrubberRating := filterNumbers(diagnosticNumbers, BIT_CRITERIA_CO2)
	fmt.Printf("Oxygen Generator Rating = %d, CO2 Scrubber Rating = %d\n", oxygenGeneratorRating, co2ScrubberRating)
	fmt.Printf("Oxygen Generator Rating * CO2 Scrubber Rating = %d\n", oxygenGeneratorRating*co2ScrubberRating)

	fmt.Printf("\n---------------------------\n")
}
