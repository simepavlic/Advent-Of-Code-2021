package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoIntArray(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elements := make([]int, 0)
	var elem int
	for scanner.Scan() {
		elem, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		elements = append(elements, elem)
	}

	return elements, scanner.Err()
}

type Direction int64

const (
	Forward Direction = iota
	Up
	Down
)

type Command struct {
	Direction Direction
	Amount    int
}

func ReadFileCommandsIntoArray(filename string) ([]Command, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elements := make([]Command, 0)
	var elem Command
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")
		switch inputs[0] {
		case "forward":
			elem.Direction = Forward
		case "up":
			elem.Direction = Up
		case "down":
			elem.Direction = Down
		default:
			return nil, errors.New("unknown direction")
		}
		amount, err := strconv.Atoi(inputs[1])
		if err != nil {
			return nil, err
		}
		elem.Amount = amount
		elements = append(elements, elem)
	}

	return elements, scanner.Err()
}

func ReadBinaryIntoStringArray(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elements := make([]string, 0)
	for scanner.Scan() {
		elements = append(elements, scanner.Text())
	}

	return elements, scanner.Err()
}
