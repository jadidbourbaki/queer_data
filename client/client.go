package main

import (
	"bytes"
	"fmt"
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

func SendBloomFilter(bloomFilter *bloom.BloomFilter) error {
	json, err := bloomFilter.MarshalJSON()

	if err != nil {
		return err
	}

	url := "http://localhost:8090/queerdata"
	fmt.Println("Sending to URL:", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	if err != nil {
		return err
	}

	req.Header.Set("X-Custom-Header", "BloomFilter")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil

}
