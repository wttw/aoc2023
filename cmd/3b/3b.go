package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{""}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, "."+line+".")
	}
	lines[0] = strings.Repeat(".", 2+len(lines[1]))
	lines = append(lines, lines[0][:])

	digitsRe := regexp.MustCompile(`[0-9]+`)
	gears := map[string][]int{}
	for y, line := range lines {
		matches := digitsRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				log.Fatalln(err)
			}
			for x := match[0]; x < match[1]; x++ {
				for _, adj := range adjacent(lines, x, y) {
					gears[adj] = append(gears[adj], num)
				}
			}
		}
	}
	sum := 0
	for _, v := range gears {
		slices.Sort(v)
		v = slices.Compact(v)
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	fmt.Println(sum)
}

func adjacent(lines []string, x, y int) []string {
	var ret []string
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			c := lines[j+y][i+x]
			if c != '*' {
				continue
			}
			ret = append(ret, fmt.Sprintf("%d,%d", x+i, y+j))
		}
	}
	return ret
}
