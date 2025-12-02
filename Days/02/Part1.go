package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")

	var invalid []int

	if err != nil {
		log.Fatal(err)
	}
	dataList := strings.Split(string(data), ",")

	for i := 0; i < len(dataList); i++ {
		ranges := strings.Split(dataList[i], "-")

		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			log.Fatal(err)
		}

		for j := start; j < end+1; j++ {
			stringJ := strconv.FormatInt(int64(j), 10)
			// fmt.Println(stringJ)
			// fmt.Println(stringJ[:len(stringJ)/2] == stringJ[len(stringJ)/2:])
			if len(stringJ)%2 == 0 && stringJ[:len(stringJ)/2] == stringJ[len(stringJ)/2:] {
				invalid = append(invalid, j)
			}
		}
	}
	total := 0
	for i := 0; i < len(invalid); i++ {
		total += invalid[i]
	}
	fmt.Println(total)
}
