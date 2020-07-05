/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter file name:")
	scanner.Scan()
	fileName := scanner.Text()
	var fileHandler *os.File
	slice := make([]Person, 0, 1)

	for {
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			fmt.Println("Please Enter valid file name:")
			scanner.Scan()
			fileName = scanner.Text()
		} else {
			fileHandler = fi
			break
		}
	}
	fileScanner := bufio.NewScanner(fileHandler)
	var arr []string

	for fileScanner.Scan() {
		arr = strings.Split(fileScanner.Text(), " ")
		slice = append(slice, Person{arr[0], arr[1]})
	}

	fileHandler.Close()

	for _, person := range slice {
		fmt.Printf("%v - %v\n", person.fname, person.lname)
	}
}
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PersonName struct {
	firstName string
	lastName  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	tmpFileName := ""
	fmt.Println("Enter File Name :")
	fmt.Scanln(&tmpFileName)
	file, err := os.Open(tmpFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var personNames []PersonName
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = string(scanner.Text())
		var tfname = strings.Split(line, " ")[0]
		var tlname = strings.Split(line, " ")[1]

		personName := PersonName{
			firstName: tfname,
			lastName:  tlname,
		}

		personNames = append(personNames, personName)

	}

	fmt.Println("Printing Persons Struct")

	fmt.Println(personNames)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Printing Struct in human readable form:")
	fmt.Println("=======================================================")
	for _, personName := range personNames {
		fmt.Println("----------------------------")

		fmt.Println("First Name:", personName.firstName)
		fmt.Println("Last  Name:", personName.lastName)

	}

}
