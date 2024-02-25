package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

type ServerState struct {
	bloomFilter *bloom.BloomFilter
	insertCount int
	queryCount  int
	bitCount    uint
	hashCount   uint
}

func New(bitCount uint, hashCount uint) ServerState {
	rtn := ServerState{
		bloomFilter: bloom.New(bitCount, hashCount),
		bitCount:    bitCount,
		hashCount:   hashCount,
	}

	return rtn
}

func (s *ServerState) AddFromJSON(jsonBody []byte) error {
	bloomFilter := bloom.BloomFilter{}
	err := bloomFilter.UnmarshalJSON(jsonBody)

	if err != nil {
		return err
	}

	err = s.bloomFilter.Merge(&bloomFilter)
	if err != nil {
		return err
	}

	s.insertCount++

	return nil
}

func (s *ServerState) Query(name string) bool {
	s.queryCount++

	return s.bloomFilter.TestString(name)
}

func (s *ServerState) String() string {
	return fmt.Sprintln("[server state] insertCount:", s.insertCount, "queryCount:", s.queryCount)
}
