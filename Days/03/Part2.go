package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findHighestPair2(battery []int) string {
	final := ""
	index := 0
	for numbersLeft := 12; numbersLeft > 0; numbersLeft-- {
		newIndex := 0
		largest := 0

		for i := index; i < len(battery)-numbersLeft+1; i++ {
			if battery[i] > largest {
				largest = battery[i]
				newIndex = i
			}
		}
		// fmt.Println(battery[index : len(battery)-numbersLeft])
		final += strconv.Itoa(largest)
		index = newIndex + 1
	}
	return final

}

func turnToIntList(sliceStrings []string) []int {
	var newList []int
	for i := 0; i < len(sliceStrings); i++ {

		number, err := strconv.Atoi(sliceStrings[i])

		if err != nil {
			log.Fatal(err)
		}

		newList = append(newList, number)
	}
	return newList
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	batteryData := strings.Split(string(data), "\n")
	total := 0
	for i := 0; i < len(batteryData); i++ {
		battery := strings.Split(batteryData[i], "")
		number, err := strconv.Atoi(findHighestPair2(turnToIntList(battery)))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(number)
		total += number
	}
	fmt.Println(total)
}
