package main

import (
	"fmt"
	"log"
	"path/filepath"

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
	if depth < 0 {
		depth *= -1
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
	if depth < 0 {
		depth *= -1
	}
	position := horizontal * depth
	return position
}

func main() {
	absPath, _ := filepath.Abs("input.txt")
	commands, err := util.ReadFileCommandsIntoArray(absPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calculatePosition(commands))
	fmt.Println(calculatePositionWithAim(commands))
}
