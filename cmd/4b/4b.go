package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	Count   int
	Matches int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cards := []*Card{}
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
		cards = append(cards, &Card{
			Count:   1,
			Matches: matches,
		})
	}
	for i, card := range cards {
		for j := 1; j <= card.Matches; j++ {
			cards[i+j].Count += card.Count
		}
	}
	sum := 0
	for _, card := range cards {
		sum += card.Count
	}
	fmt.Println(sum)
}
