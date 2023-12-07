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
	input := "day1/input.txt"
	file, err := os.Open(input)
	scanner := bufio.NewScanner(file)
	sum := 0

	if err != nil {
		fmt.Println(err)
		return
	}

	for scanner.Scan() {

		text := scanner.Text()
		spelledRegExp := regexp.MustCompile(`(oneight|twone|threeight|fiveight|sevenine|eightwo|eighthree|nineight|one|two|three|four|five|six|seven|eight|nine)`)
		cleanText := spelledRegExp.ReplaceAllStringFunc(text, func(match string) string {
			return replaceSpelledNumber(match)
		})

		numbersRegExp := regexp.MustCompile(`[0-9]`)
		numbers := strings.Join(numbersRegExp.FindAllString(cleanText, -1), "")

		if len(numbers) == 1 {
			numbers = numbers + numbers
		}

		if len(numbers) > 2 {
			numbers = string(numbers[0]) + string(numbers[len(numbers)-1])
		}

		num, err := strconv.Atoi(numbers)

    fmt.Println(text, cleanText, num, sum)

		if err != nil {
			fmt.Println(err)
			return
		}

		sum += num
	}

	fmt.Println(sum)

	defer file.Close()
}

func replaceSpelledNumber(match string) string {
	spelledToNumeric := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
    "oneight": "18",
    "twone": "21",
    "threeight": "38",
    "fiveight": "58",
    "sevenine": "79",
    "eightwo": "82",
    "eighthree": "83",
    "nineight": "98",
	}

	return spelledToNumeric[match]
}
