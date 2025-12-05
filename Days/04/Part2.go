package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanLines(lines []string) ([]string, int) {
	final := ""
	total := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			item := string(lines[y][x])
			if item == "@" {
				seroundings := []string{}

				// Top
				if y-1 >= 0 && lines[y-1][x] == '@' {
					seroundings = append(seroundings, string(lines[y-1][x]))
				}
				// Bottom
				if y+1 < len(lines) && lines[y+1][x] == '@' {
					seroundings = append(seroundings, string(lines[y+1][x]))
				}
				// Left
				if x-1 >= 0 && lines[y][x-1] == '@' {
					seroundings = append(seroundings, string(lines[y][x-1]))
				}
				// Right
				if x+1 < len(lines[y]) && lines[y][x+1] == '@' {
					seroundings = append(seroundings, string(lines[y][x+1]))
				}
				//Top Right
				if y-1 >= 0 && x+1 < len(lines[y-1]) && lines[y-1][x+1] == '@' {
					seroundings = append(seroundings, string(lines[y-1][x+1]))
				}
				//Top Left
				if y-1 >= 0 && x-1 >= 0 && lines[y-1][x-1] == '@' {
					seroundings = append(seroundings, string(lines[y-1][x-1]))
				}
				//Bottom Right
				if y+1 < len(lines) && x+1 < len(lines[y+1]) && lines[y+1][x+1] == '@' {
					seroundings = append(seroundings, string(lines[y+1][x+1]))
				}
				//Bottom Left
				if y+1 < len(lines) && x-1 >= 0 && lines[y+1][x-1] == '@' {
					seroundings = append(seroundings, string(lines[y+1][x-1]))
				}
				// fmt.Printf("Target (%d, %d) S: %s\n", x, y, seroundings)
				if len(seroundings) < 4 {
					final += "."
					total += 1
				} else {
					final += "@"
				}
			} else {
				final += "."
			}
		}
		final += "\n"
	}

	return strings.Split(final, "\n")[:len(strings.Split(final, "\n"))-1], total
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	total := 0

	lines := strings.Split(string(data), "\n")
	newLines, totalNum := cleanLines(lines)
	total += totalNum
	// fmt.Println(newLines, lines)
	for totalNum != 0 {
		newLines, totalNum = cleanLines(newLines)
		total += totalNum
	}
	fmt.Println(total)
}
