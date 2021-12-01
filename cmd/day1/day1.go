package main

import (
	"fmt"
	"log"

	"github.com/simepavlic/advent-of-code-2021/internal/util"
)

func countMeasurementsSlidingWindow(window int, depths []int) int {
	count := 0
	for i := range depths[:len(depths)-window] {
		if depths[i+window] > depths[i] {
			count++
		}
	}
	return count
}

func main() {

	depths, err := util.ReadFileIntoIntArray("input/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := countMeasurementsSlidingWindow(1, depths)
	fmt.Println(count)
	count = countMeasurementsSlidingWindow(3, depths)
	fmt.Println(count)

}
