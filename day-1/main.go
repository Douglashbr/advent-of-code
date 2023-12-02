package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	json, _ := getJsonValues("input.json")
	getSum(json)
}

func getSum(values string) {
	splitedValue := strings.Fields(values)

	var accumulatedValue uint64 = 0

	for _, word := range splitedValue {
		splitedWord := strings.Split(word, "")
		var sliceNumbers []string

		for _, letter := range splitedWord {
			match := regexp.MustCompile("[0-9]")
			findNumber := match.FindAllString(letter, -1)
			if len(findNumber) != 0 {
				sliceNumbers = append(sliceNumbers, letter)
			}
		}

		firstElement := sliceNumbers[0]
		lastElement := sliceNumbers[len(sliceNumbers)-1]
		concatNumber, err := strconv.ParseUint(firstElement+lastElement, 10, 0)
		if err != nil {
			panic(err)
		}

		accumulatedValue = accumulatedValue + concatNumber
	}

	fmt.Println("Final value => ", accumulatedValue)
}

type Input struct {
	Value string `json:"values"`
}

func getJsonValues(fileName string) (string, error) {
	jsonFile, err := os.Open(fileName)

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var input Input

	json.Unmarshal(byteValue, &input)

	return input.Value, err
}
