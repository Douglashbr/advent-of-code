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
	json2, _ := getJsonValues("example2.json")
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
	numberEnum := map[string]string{
		"one":   "o1ne",
		"two":   "t2wo",
		"three": "th3ree",
		"four":  "f4our",
		"five":  "fi5ve",
		"six":   "si6x",
		"seven": "se7ven",
		"eight": "ei8ght",
		"nine":  "ni9ne",
	}

	splitedValue := strings.Fields(values)

	var accumulatedValue uint64 = 0

	for _, word := range splitedValue {
		matchTextNumbers := regexp.MustCompile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")
		replaceStringToNumber := matchTextNumbers.ReplaceAllStringFunc(word, func(v string) string {
			return numberEnum[v]
		})
		fmt.Println("spl: ", replaceStringToNumber)
		splitedWord := strings.Split(replaceStringToNumber, "")
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
