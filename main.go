package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// make an empty slice
	input := make([]string, 0)

	//create new scanner
	scanner := bufio.NewScanner(os.Stdin)

	//
	for {
		fmt.Print("Numbers: ")

		scanner.Scan()

		text := scanner.Text()

		if len(text) != 0 {

			fmt.Println(text)

			input = append(input, text)
		} else {
			break
		}
	}

	fmt.Println(input)
}
