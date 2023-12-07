package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := "day2/input.txt"
	file, err := os.Open(input)
	scanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	maxValues := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		idPattern := regexp.MustCompile(`Game (\d+)`)
		idMatch := idPattern.FindStringSubmatch(text)
		text = idPattern.ReplaceAllString(text, "")

		id, err := strconv.Atoi(idMatch[1])
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
			continue
		}

		possible := true
		setRegex := regexp.MustCompile(`[^;]+`)
		sets := setRegex.FindAllString(text, -1)
		for _, set := range sets {

			set = strings.TrimSpace(set)
			elements := strings.Split(set, ",")

			for i, element := range elements {
				counts := make(map[string]int)

				elements[i] = strings.TrimSpace(element)

				colorRegex := regexp.MustCompile(`([a-zA-Z]+)`)
				numberRegex := regexp.MustCompile(`(\d+)`)

				color := colorRegex.FindString(elements[i])
				number, _ := strconv.Atoi(numberRegex.FindString(elements[i]))

				counts[color] += number
				if counts[color] > maxValues[color] {
					possible = false
					break
				}

				fmt.Println(color, number, counts[color])
			}

		}
		if possible {
			sum += id
			fmt.Println(id, sum)
		}

	}
	fmt.Println(sum)

}

func part2() {
	input := "day2/input.txt"
	file, err := os.Open(input)
	scanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		idPattern := regexp.MustCompile(`Game (\d+)`)
		idMatch := idPattern.FindStringSubmatch(text)
		text = idPattern.ReplaceAllString(text, "")

		id, err := strconv.Atoi(idMatch[1])
		if err != nil {
			fmt.Printf("Error converting number: %v\n", err)
			continue
		}

		maxValues := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		setRegex := regexp.MustCompile(`[^;]+`)
		sets := setRegex.FindAllString(text, -1)
		for _, set := range sets {

			set = strings.TrimSpace(set)
			elements := strings.Split(set, ",")

			for i, element := range elements {
				counts := make(map[string]int)

				elements[i] = strings.TrimSpace(element)

				colorRegex := regexp.MustCompile(`([a-zA-Z]+)`)
				numberRegex := regexp.MustCompile(`(\d+)`)

				color := colorRegex.FindString(elements[i])
				number, _ := strconv.Atoi(numberRegex.FindString(elements[i]))

				counts[color] += number
				if counts[color] > maxValues[color] {
					maxValues[color] = counts[color]
				}

				fmt.Println(color, number, counts[color])
			}
		}
		power := maxValues["red"] * maxValues["green"] * maxValues["blue"]
		sum += power
		fmt.Println(maxValues)
		fmt.Println(id, sum)

	}
	fmt.Println(sum)

}
