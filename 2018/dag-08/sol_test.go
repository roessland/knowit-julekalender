package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestM(t *testing.T) {
	dolls := []Doll{
		{50, "red", 0},
		{100, "blue", 0},
		{100, "orange", 0},
		{101, "red", 0},
	}
	assert.Equal(t, 1, M(dolls, 0))
	assert.Equal(t, 2, M(dolls, 1))
	assert.Equal(t, 2, M(dolls, 2))
	assert.Equal(t, 3, M(dolls, 3))
}
