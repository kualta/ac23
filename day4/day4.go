package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := "day4/input.txt"

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		re := regexp.MustCompile(`\|`)
		split := re.Split(text, 2)
		winningNumbers := strings.Fields(split[0])
		cardNumbers := strings.Fields(split[1])

		fmt.Println(winningNumbers)
		fmt.Println(cardNumbers)

		// check whether card numbers in winning numbers

		cardWorth := 0
		for i := 0; i < len(cardNumbers); i++ {
			if slices.Contains(winningNumbers, cardNumbers[i]) {
				if cardWorth == 0 {
					cardWorth = 1
				} else {
					cardWorth *= 2
				}
			}

		}
		sum += cardWorth
	}

	fmt.Println(sum)

}

// recursively calculate value of each card
// func calculateValue(linesValuesMap *map[int]int, winningNumbers []string, cardNumbers []string) int {
// 	cardWorth := 0

// 	for i := 0; i < len(cardNumbers); i++ {
// 		if slices.Contains(winningNumbers, cardNumbers[i]) {
// 			if cardWorth == 0 {
// 				cardWorth = 1
// 			} else {
// 				cardWorth *= 2
// 			}
// 		}
// 	}

// }

type Card struct {
	WinnningNumbers []string
	Numbers         []string
	Amount          int
}

func part2() {
	file, _ := os.Open("day4/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	cards := []Card{}
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`\|`)
		split := re.Split(line, 2)
		winningNumbers := strings.Fields(split[0][7 : len(split[0])-1])
		cardNumbers := strings.Fields(split[1])

		amount := 0
		for i := 0; i < len(cardNumbers); i++ {
			if slices.Contains(winningNumbers, cardNumbers[i]) {
				amount += 1
			}
		}

		cards = append(cards, Card{
			WinnningNumbers: winningNumbers,
			Numbers:         cardNumbers,
			Amount:          amount,
		})
	}

	for i := len(cards) - 1; i >= 0; i-- {
		card := cards[i]
		fmt.Println(card)

		newAmount := card.Amount

		for j := 1; j <= card.Amount; j++ {
			fmt.Println("getting: ", cards[i+j])
			pastAmount := cards[i+j].Amount
			if pastAmount > 0 {
				newAmount += pastAmount
			}
			fmt.Println(newAmount)
		}

		cards[i].Amount = newAmount

	}

	sum := 0

	for i := 0; i < len(cards); i++ {
		sum += cards[i].Amount + 1
	}

	fmt.Println(sum)

}
