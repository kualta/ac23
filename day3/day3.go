package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Gear struct {
	Counter int
	Value   int
}

func findNumber(matrix *[][]string, i int, j int) int {
	symbols := []string{"#", "+", "*", "$", "%", "-", "&", "@", "=", "!", "^", "?", "/", "(", ")", "{", "}", "."}
	number := ""

	for {
		if i < 0 || j < 0 || i >= len((*matrix)) || j >= len((*matrix)[0]) {
			break
		}
		element := (*matrix)[i][j]

		if element >= "0" && element <= "9" {
			j--
		}

		if slices.Contains(symbols, element) || element == "." || j < 0 {
			j++
			break
		}
	}

	for {
		if i < 0 || j < 0 || i >= len((*matrix)) || j >= len((*matrix)[0]) {
			break
		}
		element := (*matrix)[i][j]
		fmt.Println(element)

		if element >= "0" && element <= "9" {
			number += element
			(*matrix)[i][j] = " "

			j++
		}

		if slices.Contains(symbols, element) {
			break
		}
	}

	num, _ := strconv.Atoi(number)
	return num
}

func readNumber(matrix *[][]string, sum *int, gears *map[string]Gear, i int, j int) {
	symbols := []string{"#", "+", "*", "$", "%", "-", "&", "@", "=", "!", "^", "?", "/", "(", ")", "{", "}"}
	number := ""
	included := false

	for {
		if i < 0 || j < 0 || i >= len((*matrix)) || j >= len((*matrix)[0]) {
			fmt.Println("out of bounds")
			break
		}

		element := (*matrix)[i][j]

		if slices.Contains(symbols, element) || element == "." {
			fmt.Println("found symbol")
			break
		}

		dx := []int{-1, 1, 0, 0, -1, -1, 1, 1}
		dy := []int{0, 0, -1, 1, -1, 1, -1, 1}

		for k := 0; k < 8; k++ {
			ni := i + dx[k]
			nj := j + dy[k]
			if ni >= 0 && ni < len(*matrix) && nj >= 0 && nj < len((*matrix)[0]) {
				if slices.Contains(symbols, (*matrix)[ni][nj]) {
					included = true
					break
				}
			}
		}

		number += element
		(*matrix)[i][j] = " "

		j++
	}

	if included {
		fmt.Println("included")
		num, _ := strconv.Atoi(number)
		*sum += num
	} else {
		fmt.Println(" ugly: " + number)
	}

	fmt.Println(number)
}

func part1() {
	input := "day3/input.txt"

	content, _ := os.ReadFile(input)

	fileContent := string(content)
	rows := strings.Split(fileContent, "\n")

	matrix := make([][]string, len(rows))
	for i, row := range rows {
		symbols := strings.Split(row, "")
		matrix[i] = symbols
	}

	fmt.Println(matrix)

	sum := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			element := matrix[i][j]
			if element >= "0" && element <= "9" {
				readNumber(&matrix, &sum, nil, i, j)
			}
		}

	}
	fmt.Println(sum)
}

func part2() {
	input := "day3/input.txt"

	content, _ := os.ReadFile(input)

	fileContent := string(content)
	rows := strings.Split(fileContent, "\n")

	matrix := make([][]string, len(rows))
	for i, row := range rows {
		symbols := strings.Split(row, "")
		matrix[i] = symbols
	}

	gearMap := make(map[string]Gear)

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			element := matrix[i][j]
			if element == "*" {
				dx := []int{-1, 1, 0, 0, -1, -1, 1, 1}
				dy := []int{0, 0, -1, 1, -1, 1, -1, 1}

				for k := 0; k < 8; k++ {
					ni := i + dx[k]
					nj := j + dy[k]
					if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len((matrix)[0]) {
						if (matrix)[ni][nj] >= "0" && (matrix)[ni][nj] <= "9" {
							num := findNumber(&matrix, ni, nj)

							key := fmt.Sprintf("%d-%d", i, j)
							if gear, ok := gearMap[key]; !ok {
								gearMap[key] = Gear{
									Counter: 1,
									Value:   num,
								}
							} else {
								gearMap[key] = Gear{
									Counter: gear.Counter + 1,
									Value:   gear.Value * num,
								}
							}
						}
					}
				}

			}
		}
	}

	gearSum := 0
	for _, gear := range gearMap {
		if gear.Counter == 2 {
			gearSum += gear.Value
		}
	}

	fmt.Println(gearMap)
	fmt.Println("Sum of gears with count 2:", gearSum)
}
