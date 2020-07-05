package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a name:")
	name, _ := reader.ReadString('\n')
	mname := string(name)
	m["name"] = mname[0 : len(mname)-1]
	fmt.Printf("Input an address:")
	address, _ := reader.ReadString('\n')
	maddress := string(address)
	m["address"] = maddress[0 : len(maddress)-1]

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	os.Stdout.Write(b)
}
