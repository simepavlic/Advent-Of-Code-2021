package main

import (
	"fmt"
	"log"

	"github.com/simepavlic/advent-of-code-2021/internal/util"
)

func calculatePosition(commands []util.Command) int {
	horizontal, depth := 0, 0
	for _, c := range commands {
		switch c.Direction {
		case util.Forward:
			horizontal += c.Amount
		case util.Up:
			depth -= c.Amount
		case util.Down:
			depth += c.Amount
		}
	}
	position := horizontal * depth
	return position
}

func calculatePositionWithAim(commands []util.Command) int {
	var horizontal, depth, aim int
	for _, c := range commands {
		switch c.Direction {
		case util.Forward:
			horizontal += c.Amount
			depth += aim * c.Amount
		case util.Up:
			aim -= c.Amount
		case util.Down:
			aim += c.Amount
		}
	}
	position := horizontal * depth
	return position
}

func main() {
	commands, err := util.ReadFileCommandsIntoArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calculatePosition(commands))
	fmt.Println(calculatePositionWithAim(commands))
}
