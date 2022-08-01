package main

import (
	"bufio"
	"fmt"
	"os"
)

const INPUT_PATH = "input.txt"

type Movement struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (movement *Movement) apply(nexMovement *Movement) {
	movement.Horizontal += nexMovement.Horizontal
	movement.Depth += nexMovement.Depth
}

func (movement *Movement) applyWithAim(nexMovement *Movement) {
	movement.Aim += nexMovement.Depth
	movement.Horizontal += nexMovement.Horizontal
	movement.Depth += movement.Aim * nexMovement.Horizontal
}

func parseInput(inputPath string) []Movement {
	input, err := os.Open(inputPath)
	if err != nil {
		panic("Input file not found")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var movements []Movement
	for scanner.Scan() {
		nextLine := scanner.Text()

		var direction string
		var distance int
		n, err := fmt.Sscanf(nextLine, "%s %d", &direction, &distance)
		if err != nil || n < 2 {
			panic(fmt.Sprintf("Cannot parse value '%s'", nextLine))
		}

		switch direction {
		case "forward":
			movements = append(movements, Movement{Horizontal: distance})
		case "up":
			movements = append(movements, Movement{Depth: -distance})
		case "down":
			movements = append(movements, Movement{Depth: distance})
		default:
			panic(fmt.Sprintf("Unknown direction '%s'", direction))
		}
	}

	return movements
}

func applyAllMovements(movements []Movement) Movement {
	var finalMovement Movement
	for _, movement := range movements {
		finalMovement.apply(&movement)
	}
	return finalMovement
}

func applyWithAimAllMovements(movements []Movement) Movement {
	var finalMovement Movement
	for _, movement := range movements {
		finalMovement.applyWithAim(&movement)
	}
	return finalMovement
}

func main() {
	fmt.Print(
		"--- Advent of Code 2021 ---\n",
		"          Day   2          \n",
		"                           \n",
		"          Part  A          \n")

	inputMovements := parseInput(INPUT_PATH)

	finalMovement := applyAllMovements(inputMovements)
	fmt.Printf("Submarine Depth * Horizontal Position = %d\n", finalMovement.Depth*finalMovement.Horizontal)

	fmt.Printf("\n          Part  B          \n")
	finalMovement = applyWithAimAllMovements(inputMovements)
	fmt.Printf("Submarine Depth * Horizontal Position With Aim = %d\n", finalMovement.Depth*finalMovement.Horizontal)

	fmt.Printf("\n---------------------------\n")
}
