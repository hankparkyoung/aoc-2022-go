package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	str "strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	inputFile := bufio.NewScanner(file)
	inputList := []string{}

	for inputFile.Scan() {
		inputList = append(inputList, inputFile.Text())
	}

	totalsList := addTotals(inputList)
	largest := findLargest(totalsList)

	fmt.Println("Largest is:")
	fmt.Println(largest)

	largestThree := findLargestThree(totalsList)

	fmt.Println("Largest three are:")
	fmt.Println(largestThree)
	fmt.Println("And their total is:")
	fmt.Println(largestThree[0] + largestThree[1] + largestThree[2])

}

func addTotals(list []string) []int {
	totals := []int{}
	sum := 0

	for _, num := range list {
		if num == "" {
			totals = append(totals, sum)
			sum = 0
		} else {
			toAdd, _ := str.Atoi(num)
			sum += toAdd
		}
	}

	return totals
}

func findLargest(totals []int) int {
	largest := 0

	for _, total := range totals {
		if total > largest {
			largest = total
		}
	}

	return largest
}

func findLargestThree(totals []int) []int {
	first := 0
	second := 0
	third := 0

	for _, total := range totals {
		if total > third {
			if total > second {
				if total > first {
					third, second, first = second, first, total
				} else {
					third, second = second, total
				}
			} else {
				third = total
			}
		}
	}

	return []int{first, second, third}
}
