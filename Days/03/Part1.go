package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func findHighestPair(battery []string) string {
	sorted := make([]string, len(battery))
	copy(sorted, battery[:len(battery)-1])
	sort.Strings(sorted)
	slices.Reverse(sorted)

	first := slices.Index(battery, sorted[0])
	seconds := "0"

	for i := first + 1; i < len(battery); i++ {
		if battery[i] > seconds {
			seconds = battery[i]
		}
	}
	return sorted[0] + seconds
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
		number, err := strconv.Atoi(findHighestPair(battery))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(number)
		total += number
	}
	fmt.Println(total)
}
