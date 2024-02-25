package main

import (
	"fmt"
	"io"
	"net/http"
)

const bitCount uint = 1000
const hashCount uint = 5

var serverState ServerState

func insert(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inserting")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "failure")
		return
	}

	err = serverState.AddFromJSON(body)
	if err != nil {
		io.WriteString(w, "failure")
		return
	}

	fmt.Println(serverState.String())
	io.WriteString(w, "success")
}

func query(w http.ResponseWriter, req *http.Request) {
	fmt.Println("querying")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	if serverState.Query(string(body)) {
		io.WriteString(w, "true")
		return
	}

	io.WriteString(w, "false")
}

func main() {
	url := "localhost:8090"

	serverState = New(bitCount, hashCount)

	http.HandleFunc("/insert", insert)
	http.HandleFunc("/query", query)

	fmt.Println("starting at url", url)
	http.ListenAndServe(url, nil)
}
