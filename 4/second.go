package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()

		sections := strings.Split(txt, ",")

		first := sections[0]
		second := sections[1]

		firstParams := strings.Split(first, "-")

		secondParms := strings.Split(second, "-")

		a1, _ := strconv.Atoi(firstParams[0])
		a2, _ := strconv.Atoi(firstParams[1])

		b1, _ := strconv.Atoi(secondParms[0])
		b2, _ := strconv.Atoi(secondParms[1])

		if (b1 >= a1 && b1 <= a2) || (a1 >= b1 && a1 <= b2) || (b2 >= a1 && b2 <= a2) || (a2 >= b1 && a2 <= b2) {
			sum++
			continue
		}

	}
	fmt.Println(sum)
}
