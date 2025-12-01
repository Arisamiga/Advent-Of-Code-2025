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

	if err != nil {
		log.Fatal(err)
	}

	current := 50

	times := 0

	turns := strings.Split(string(data), "\n")

	for i := 0; i < len(turns); i++ {
		data := turns[i]

		num, err := strconv.Atoi(data[1:])
		if err != nil {
			log.Fatal(err)
		}

		if data[0] == 'R' {
			current += num
		} else {
			current -= num
		}

		if current < 0 {
			for current < 99 && current < 0 {
				current += 100
			}
		} else if current > 99 {
			for current > 99 && current > 0 {
				current -= 100
			}
		}

		if current == 0 {
			times += 1
		}

		// fmt.Println(current)
	}

	fmt.Println(times)
}
