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
	json2, _ := getJsonValues("input.json")
	getSum1(json)
	getSum2(json2)
}

func getSum1(values string) {
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

	fmt.Println("Final value 1 => ", accumulatedValue)
}

func getSum2(values string) {
	splitedValue := strings.Fields(values)

	var accumulatedValue uint64 = 0

	for _, word := range splitedValue {

		word = strings.ReplaceAll(word, "one", "o1ne")
		word = strings.ReplaceAll(word, "two", "t2wo")
		word = strings.ReplaceAll(word, "three", "th3ree")
		word = strings.ReplaceAll(word, "four", "f4our")
		word = strings.ReplaceAll(word, "five", "fi5ve")
		word = strings.ReplaceAll(word, "six", "si6x")
		word = strings.ReplaceAll(word, "seven", "se7ven")
		word = strings.ReplaceAll(word, "eight", "ei8ght")
		word = strings.ReplaceAll(word, "nine", "ni9ne")

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

	fmt.Println("Final value 2 => ", accumulatedValue)
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
