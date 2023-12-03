package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Oi")

	i, _ := getJsonValues("input.json")

	result, result2 := getGames(i)

	fmt.Println("Result => ", result)
	fmt.Println("Result 2 => ", result2)
}

func getGames(values map[string]string) (int, int) {
	bags := "12 red 13 green 14 blue"
	bagsSplited := strings.Split(bags, " ")
	sumOfGames := 0
	multipleFewestCubes := 0

	for index, value := range values {
		canBeTheGame := true
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		totalCubesPerLine := 0

		replaceComma := strings.ReplaceAll(value, ",", "")
		replaceSemicolon := strings.ReplaceAll(replaceComma, ";", "")
		splitValues := strings.Split(replaceSemicolon, " ")

		for colorIndex, color := range splitValues {
			if colorIndex == 0 {
				continue
			}

			colorNumber, _ := strconv.Atoi(splitValues[colorIndex-1])

			if color == "blue" && colorNumber > maxBlue {
				maxBlue = colorNumber
			}
			if color == "green" && colorNumber > maxGreen {
				maxGreen = colorNumber
			}
			if color == "red" && colorNumber > maxRed {
				maxRed = colorNumber
			}

			for bagIndex, bagValue := range bagsSplited {
				if bagIndex == 0 {
					continue
				}

				if canBeTheGame == false {
					continue
				}

				bagNumber, _ := strconv.Atoi(bagsSplited[bagIndex-1])
				if bagValue == "blue" && color == "blue" && colorNumber > bagNumber {
					canBeTheGame = false
				}
				if bagValue == "green" && color == "green" && colorNumber > bagNumber {
					canBeTheGame = false
				}
				if bagValue == "red" && color == "red" && colorNumber > bagNumber {
					canBeTheGame = false
				}
			}
		}

		indexSplit := strings.Split(index, " ")
		indexNumber, _ := strconv.Atoi(indexSplit[1])
		if canBeTheGame {
			sumOfGames += int(indexNumber)
		}

		totalCubesPerLine = maxBlue * maxGreen * maxRed
		multipleFewestCubes += totalCubesPerLine
	}

	return sumOfGames, multipleFewestCubes
}

func getJsonValues(fileName string) (map[string]string, error) {
	input := make(map[string]string)
	jsonFile, err := os.Open(fileName)

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &input)

	return input, err
}
