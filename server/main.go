package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bits-and-blooms/bloom/v3"
)

func printHeaders(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received queer data")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	bloomFilter := bloom.BloomFilter{}
	bloomFilter.UnmarshalJSON(body)

	if bloomFilter.TestString("TEST NAME") {
		fmt.Println("Check 1 passed")
	}

	if !bloomFilter.TestString("NOT TEST NAME") {
		fmt.Println("Check 2 passed")
	}

	if bloomFilter.Cap() == uint(100) {
		fmt.Println("Check 3 passed")
	}

	if bloomFilter.K() == uint(5) {
		fmt.Println("Check 4 passed")
	}

}

func main() {
	http.HandleFunc("/queerdata", printHeaders)

	http.ListenAndServe("localhost:8090", nil)
}
