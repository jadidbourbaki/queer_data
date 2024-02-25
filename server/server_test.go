package main

import (
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/stretchr/testify/assert"
)

func TestAddFromJSON(t *testing.T) {
	bloomFilter := bloom.New(100, 5)
	serverState := New(100, 5)

	jsonBody, err := bloomFilter.MarshalJSON()
	assert.NoError(t, err)

	assert.NoError(t, serverState.AddFromJSON(jsonBody))

	bloomFilter = bloom.New(101, 6)
	jsonBody, err = bloomFilter.MarshalJSON()
	assert.NoError(t, err)

	assert.Error(t, serverState.AddFromJSON(jsonBody))
}

func TestAddFromJSONCorrectness(t *testing.T) {
	bloomFilter := bloom.New(100, 5)
	serverState := New(100, 5)

	testString := "HELLO"
	notTestString := "NOT HELLO"

	bloomFilter.AddString("HELLO")

	assert.False(t, serverState.Query(testString))

	jsonBody, err := bloomFilter.MarshalJSON()
	assert.NoError(t, err)

	assert.NoError(t, serverState.AddFromJSON(jsonBody))

	assert.True(t, serverState.Query(testString))
	assert.False(t, serverState.Query(notTestString))
}
