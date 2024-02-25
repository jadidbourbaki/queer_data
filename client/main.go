package main

import (
	"fmt"
	"os"
)

const bitCount uint = 1000
const hashCount uint = 5

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("usage:", args[0], "[insert | query] <name>")
		return
	}

	command := args[1]
	name := args[2]

	if command != "insert" && command != "query" {
		fmt.Println("invalid command")
		return
	}

	if command == "insert" {
		client := New(bitCount, hashCount)

		fmt.Println("adding name:", name)
		bloomFilter := client.AddName(name)

		response, err := SendBloomFilter(bloomFilter)

		if err != nil {
			panic(err)
		}

		fmt.Println("response:")
		fmt.Println(response)
		return
	}

	// query
	response, err := QueryName(name)

	if err != nil {
		panic(err)
	}

	fmt.Println("response:")
	fmt.Println(response)

}
