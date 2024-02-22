package main

import "github.com/bits-and-blooms/bloom/v3"

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
