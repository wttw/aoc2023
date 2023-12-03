package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	lines := []string{""}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, "."+line+".")
	}
	lines[0] = strings.Repeat(".", 2+len(lines[1]))
	lines = append(lines, lines[0][:])
	digitsRe := regexp.MustCompile(`[0-9]+`)
	for y, line := range lines {
		matches := digitsRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				log.Fatalln(err)
			}
			adj := ""
			for x := match[0]; x < match[1]; x++ {
				adj = adj + adjacent(lines, x, y)
			}
			//fmt.Printf("%d %s\n", num, adj)
			if adj != "" {
				sum += num
			}
		}
	}
	fmt.Println(sum)
}

func adjacent(lines []string, x, y int) string {
	ret := ""
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			c := lines[j+y][i+x]
			if c == '.' {
				continue
			}
			if c >= '0' && c <= '9' {
				continue
			}
			ret += string(c)
		}
	}
	return ret
}
