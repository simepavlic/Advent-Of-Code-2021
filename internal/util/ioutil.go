package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileIntoIntArray(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := make([]int, 0)
	var depth int
	for scanner.Scan() {
		depth, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		depths = append(depths, depth)
	}

	return depths, scanner.Err()
}
