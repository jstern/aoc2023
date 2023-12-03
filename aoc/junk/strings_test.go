package junk_test

import (
	"testing"

	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func Test_ReverseString(t *testing.T) {
	assert.Equal(t, "evil", junk.ReverseString("live"))
}
