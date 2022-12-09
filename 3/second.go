package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	items := []string{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var bag []string

	for scanner.Scan() {
		bag = append(bag, scanner.Text())

		if len(bag) == 3 {
			fmt.Println(bag)

			for _, c := range bag[0] {
				if strings.ContainsAny(bag[1], string(c)) && strings.ContainsAny(bag[2], string(c)) {
					items = append(items, string(c))
					bag = bag[:0]
					break
				}
			}
		}
	}

	values := "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sum := 0

	for _, char := range items {
		ind := strings.Index(values, char)
		sum += ind
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
