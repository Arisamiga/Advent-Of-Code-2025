package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

			if len(stringJ) < 2 {
				continue
			}

			for e := 2; e < len(stringJ); e++ {
				// if stringJ[:len(stringJ)/e] == stringJ[len(stringJ)/e:] {
				// 	invalid = append(invalid, j)
				// 	continue
				// }
				if len(stringJ)%e != 0 {
					// fmt.Printf("NOT ALLOWED %d %d (%d)\n", e, len(stringJ), len(stringJ)%e)
					continue
				}

				match := true
				for f := e; f <= len(stringJ)+1; f += e {
					if f+e <= len(stringJ) && stringJ[:e] != stringJ[f:f+e] {
						// invalid = append(invalid, j)
						// fmt.Printf("--- %s - %d + %d len(%d)\n", stringJ, f, e, len(stringJ))
						// fmt.Println(stringJ[:e], stringJ[f:f+e])
						match = false
						break
					}
				}

				if match && !slices.Contains(invalid, j) {
					invalid = append(invalid, j)
				}
			}

			theSame := true

			for e := 0; e < len(stringJ); e++ {
				if stringJ[0] != stringJ[e] {
					theSame = false
				}
			}

			if theSame && !slices.Contains(invalid, j) {
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
