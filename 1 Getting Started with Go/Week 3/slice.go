package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intList := make([]int, 0, 3)

	var scanValue string
	for {
		fmt.Print("Enter an integer", ": ")
		_, err := fmt.Scan(&scanValue)
		exitCode := strings.ToLower(scanValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if exitCode == "x" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		parseIntScanValue, err := strconv.Atoi(scanValue)
		intList = append(intList, parseIntScanValue)
		sort.Ints(intList)
		fmt.Println(intList)
	}
}
