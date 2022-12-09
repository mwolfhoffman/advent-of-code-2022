package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	repeatedItems := []string{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()
		firstStr := txt[0:(len(txt) / 2)]
		secondStr := txt[(len(txt) / 2):]
		fmt.Println(firstStr)
		println(secondStr)

		for _, c := range firstStr {
			if strings.ContainsAny(secondStr, string(c)) {
				slice := append(repeatedItems, string(c))
				repeatedItems = slice
				break
			}
		}

		values := "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		sum := 0

		for _, char := range repeatedItems {
			ind := strings.Index(values, char)
			sum += ind
		}
		fmt.Println(sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
