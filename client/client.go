package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/bits-and-blooms/bloom/v3"
)

type Client struct {
	bitCount  uint
	hashCount uint
}

func New(bitCount uint, hashCount uint) *Client {
	return &Client{bitCount, hashCount}
}

func (c *Client) AddName(name string) *bloom.BloomFilter {
	bloomFilter := bloom.New(c.bitCount, c.hashCount)

	rawName := []byte(name)
	bloomFilter.Add(rawName)

	return bloomFilter
}

func QueryName(name string) (string, error) {
	url := "http://localhost:8090/query"
	fmt.Println("sending to url:", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(name)))

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return string(body), nil

}

func SendBloomFilter(bloomFilter *bloom.BloomFilter) (string, error) {
	json, err := bloomFilter.MarshalJSON()

	if err != nil {
		return "", err
	}

	url := "http://localhost:8090/insert"
	fmt.Println("sending to url:", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return string(body), nil

}
