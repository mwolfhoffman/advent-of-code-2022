package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var totalScore int = 0

var drawKey = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var winningKey = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var losingKey = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		round := strings.Split(txt, " ")

		opponent := round[0]
		player := round[1]

		if player == "X" {
			//	Loss
			totalScore += scores[losingKey[opponent]]

		} else if player == "Y" {
			// Draw
			totalScore += 3 + scores[drawKey[opponent]]
		} else {
			//	Win
			totalScore += 6 + scores[winningKey[opponent]]
		}

		fmt.Println(totalScore)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
