package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddName(t *testing.T) {
	var bitCount uint = 100
	var hashCount uint = 5
	testValue := "TEST NAME"

	client := New(bitCount, hashCount)
	bloomFilter := client.AddName(testValue)

	assert.Equal(t, bloomFilter.Cap(), bitCount)
	assert.Equal(t, bloomFilter.K(), hashCount)
	assert.True(t, bloomFilter.TestString(testValue))
}
