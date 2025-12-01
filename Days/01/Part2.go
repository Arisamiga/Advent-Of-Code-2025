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

		steps := 1

		if data[0] == 'L' {
			steps = -1
		}

		for i := 0; i < num; i++ {
			// current = (current + ((i + 1) * steps)) % 100
			current = current + steps

			if current < 0 {
				current = 99
			} else if current > 99 {
				current = 0
			}

			if current == 0 {
				times += 1
			}
		}

	}

	fmt.Println(times)
}
