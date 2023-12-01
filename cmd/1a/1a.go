package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var filtered string
		for _, v := range line {
			if v >= '0' && v <= '9' {
				filtered = filtered + string(v)
			}
		}
		if len(filtered) == 1 {
			filtered = filtered + filtered
		}
		num := filtered[:1] + filtered[len(filtered)-1:]
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
		}
		sum += val
	}
	fmt.Println(sum)
}
