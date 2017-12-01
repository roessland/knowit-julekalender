package main

import "github.com/stretchr/testify/assert"
import "testing"

func TestNgramify(t *testing.T) {
	assert.Equal(t, "", Ngramify(4, "bob"))
	assert.Equal(t, "", Ngramify(3, "bob"))
	assert.Equal(t, "abbc", Ngramify(2, "abc"))
	assert.Equal(t, "batatmtmaman", Ngramify(3, "batman"))
}
