package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/simepavlic/advent-of-code-2021/internal/util"
)

func GammaEpsilonRate(array []string) (int, error) {
	bitAppearance := make([]int, len(array[0]))
	for _, b := range array {
		for i := range b {
			if b[i] == '1' {
				bitAppearance[i]++
			}
		}
	}
	var gamma, epsilon bytes.Buffer
	for _, num := range bitAppearance {
		var mostCommonValue rune
		if num > len(array)/2 {
			mostCommonValue = '1'
		} else {
			mostCommonValue = '0'
		}
		gamma.WriteRune(mostCommonValue)
		epsilon.WriteRune(mostCommonValue ^ 1)
	}
	gammaInt, err := strconv.ParseInt(gamma.String(), 2, 64)
	if err != nil {
		return -1, err
	}
	epsilonInt, err := strconv.ParseInt(epsilon.String(), 2, 64)
	if err != nil {
		return -1, err
	}
	return int(gammaInt) * int(epsilonInt), nil
}

func OxygenCO2Rate(array []string) (int, error) {
	var zeroArray, oneArray, oxygenArray, co2Array []string
	for _, b := range array {
		if b[0] == '1' {
			oneArray = append(oneArray, b)
		} else {
			zeroArray = append(zeroArray, b)
		}
	}
	if len(oneArray) > len(zeroArray) {
		oxygenArray = oneArray
		co2Array = zeroArray
	} else {
		oxygenArray = zeroArray
		co2Array = oneArray
	}

	for i := 1; len(oxygenArray) > 1; i++ {
		oneArray, zeroArray = []string{}, []string{}
		for _, b := range oxygenArray {
			if b[i] == '1' {
				oneArray = append(oneArray, b)
			} else {
				zeroArray = append(zeroArray, b)
			}
		}
		if len(oneArray) >= len(zeroArray) {
			oxygenArray = oneArray
		} else {
			oxygenArray = zeroArray
		}
	}

	for i := 1; len(co2Array) > 1; i++ {
		oneArray, zeroArray = []string{}, []string{}
		for _, b := range co2Array {
			if b[i] == '1' {
				oneArray = append(oneArray, b)
			} else {
				zeroArray = append(zeroArray, b)
			}
		}
		if len(oneArray) < len(zeroArray) {
			co2Array = oneArray
		} else {
			co2Array = zeroArray
		}
	}
	oxygenInt, err := strconv.ParseInt(oxygenArray[0], 2, 64)
	if err != nil {
		return -1, err
	}
	co2Int, err := strconv.ParseInt(co2Array[0], 2, 64)
	if err != nil {
		return -1, err
	}
	return int(oxygenInt) * int(co2Int), nil
}

func main() {
	binaryArray, err := util.ReadBinaryIntoStringArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	gammaEpsilonRate, err := GammaEpsilonRate(binaryArray)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gammaEpsilonRate)
	oxygenCo2Rate, err := OxygenCO2Rate(binaryArray)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oxygenCo2Rate)
}
