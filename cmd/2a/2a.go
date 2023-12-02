package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Red   int = 0
	Green     = 1
	Blue      = 2
)

func main() {
	names := map[string]int{
		"red":   Red,
		"green": Green,
		"blue":  Blue,
	}
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Let's not go with regex parsing just yet
		parts := strings.Split(line, ":")
		gameNumber, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
		if err != nil {
			log.Fatalf("%s: %s", line, err)
		}
		possible := true
		draws := strings.Split(parts[1], ";")

		for _, draw := range draws {
			drawn := []int{0, 0, 0}
			stones := strings.Split(draw, ",")

			for _, stone := range stones {
				cc := strings.Split(strings.TrimSpace(stone), " ")

				count, err := strconv.Atoi(cc[0])
				if err != nil {
					log.Fatalf("%s: '%s': %s", line, stone, err)
				}
				color, ok := names[cc[1]]
				if !ok {
					log.Fatalf("%s: bad colour '%s'", line, cc[1])
				}
				if drawn[color] != 0 {
					log.Fatalf("dupe: %s", line)
				}
				drawn[color] = count
			}
			if drawn[Red] > 12 || drawn[Green] > 13 || drawn[Blue] > 14 {
				possible = false
			}
		}
		if possible {
			sum += gameNumber
		}
	}
	fmt.Println(sum)
}
