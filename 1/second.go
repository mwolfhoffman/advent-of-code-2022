package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	current := 0
	sums := []int{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		if scanner.Text() == "" {
			sums = append(sums, current)
			current = 0
			continue
		}

		num, err := strconv.Atoi(scanner.Text())

		if (err) != nil {
			fmt.Println("Could not convert to num")
		}

		current += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(sums[:])

	slice := sums[len(sums)-3 : len(sums)]

	fmt.Println(sum(slice))

}
