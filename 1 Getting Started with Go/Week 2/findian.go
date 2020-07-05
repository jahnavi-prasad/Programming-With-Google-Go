package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter your string: ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	line = strings.ToLower(line)

	if strings.HasPrefix(line, "i") {
		if strings.Contains(line, "a") {
			if strings.HasSuffix(line, "n") {
				fmt.Printf("Found!")
			} else {
				fmt.Printf("Not Found!")
			}
		} else {
			fmt.Printf("Not Found!")
		}
	} else {
		fmt.Printf("Not Found!")
	}
}
