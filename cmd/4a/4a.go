package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		s1 := strings.Split(line, ":")
		s2 := strings.Split(s1[1], "|")
		winners := map[string]struct{}{}
		for _, num := range strings.Fields(s2[0]) {
			winners[num] = struct{}{}
		}
		matches := 0
		for _, have := range strings.Fields(s2[1]) {
			_, ok := winners[have]
			if ok {
				matches++
			}
		}
		if matches > 0 {
			sum += 1 << (matches - 1)
		}
	}
	fmt.Println(sum)
}
