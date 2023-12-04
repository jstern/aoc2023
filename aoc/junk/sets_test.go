package junk_test

import (
	"testing"

	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func Test_NewSet(t *testing.T) {
	s := junk.NewSet[string]("hi", "bye")
	assert.Len(t, s, 2)
	assert.True(t, s.Contains("hi"))
	assert.True(t, s.Contains("bye"))
}

func Test_SetContains(t *testing.T) {
	s := junk.NewSet[rune]()
	s.Add('k')
	assert.True(t, s.Contains('k'))
	assert.False(t, s.Contains('j'))
}

func Test_SetAdd(t *testing.T) {
	s := junk.NewSet[int]()
	assert.True(t, s.Add(1))
	assert.False(t, s.Add(1))
	assert.True(t, s.Add(2))
}

func Test_SetAddAll(t *testing.T) {
	s := junk.NewSet[float64]()
	s.AddAll(1.2, 3.4)
	assert.Len(t, s, 2)
	assert.True(t, s.Contains(1.2))
	assert.True(t, s.Contains(3.4))
}

func Test_SetRemove(t *testing.T) {
	s := junk.NewSet[bool]()
	s.Add(false)
	assert.True(t, s.Contains(false))
	assert.True(t, s.Remove(false))
	assert.False(t, s.Contains(false))
}

func Test_SetUnion(t *testing.T) {
	s1 := junk.NewSet('a', 'b', 'c')
	s2 := junk.NewSet('c', 'd', 'e')
	s1s2 := s1.Union(s2)
	s2s1 := s2.Union(s1)
	assert.Len(t, s1s2, 5)
	assert.Len(t, s2s1, 5)
	for v := range s1 {
		assert.True(t, s1s2.Contains(v))
		assert.True(t, s2s1.Contains(v))
	}
}

func Test_SetIntersection(t *testing.T) {
	s1 := junk.NewSet('a', 'b', 'c')
	s2 := junk.NewSet('c', 'd', 'e')
	s1s2 := s1.Intersection(s2)
	s2s1 := s2.Intersection(s1)
	assert.Len(t, s1s2, 1)
	assert.Len(t, s2s1, 1)
	assert.True(t, s1s2.Contains('c'))
	assert.True(t, s2s1.Contains('c'))
}

func Test_SetDifference(t *testing.T) {
	s1 := junk.NewSet('a', 'b', 'c')
	s2 := junk.NewSet('c', 'd', 'e')
	s1s2 := s1.Difference(s2)
	s2s1 := s2.Difference(s1)
	assert.Len(t, s1s2, 2)
	assert.Len(t, s2s1, 2)

	assert.True(t, s1s2.Contains('a'))
	assert.True(t, s1s2.Contains('b'))
	assert.False(t, s1s2.Contains('c'))

	assert.True(t, s2s1.Contains('d'))
	assert.True(t, s2s1.Contains('e'))
	assert.False(t, s2s1.Contains('c'))
}

func Test_SetValues(t *testing.T) {
	s := junk.NewSet('a', 'b', 'c')
	assert.ElementsMatch(t, s.Values(), []rune{'a', 'b', 'c'})
}
