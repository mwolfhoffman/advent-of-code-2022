package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	highest := 0
	current := 0

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > highest {
				highest = current
			}
			current = 0
			continue
		}

		num, err := strconv.Atoi(scanner.Text())

		if (err) != nil {
			fmt.Println("Could not convert to num")
		}

		current += num

		fmt.Println(highest)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// 72602
