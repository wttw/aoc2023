package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	digits := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstPos := len(line) + 1
		var firstVal int
		lastPos := -1
		var lastVal int
		for k, v := range digits {
			first := strings.Index(line, k)
			if first != -1 && first < firstPos {
				firstPos = first
				firstVal = v
			}
			last := strings.LastIndex(line, k)
			if last != -1 && last > lastPos {
				lastPos = last
				lastVal = v
			}

		}
		fmt.Printf("%s -> %d %d\n", line, firstVal, lastVal)
		val := firstVal*10 + lastVal
		sum += val
	}
	fmt.Println(sum)
}
