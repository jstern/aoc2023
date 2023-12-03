package junk

import "github.com/samber/lo"

// here's a nice unnecessary detour

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := make(map[T]struct{})
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}
	s[v] = struct{}{}
	return true
}

func (s Set[T]) AddAll(v ...T) {
	for _, vv := range v {
		s[vv] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) bool {
	if s.Contains(v) {
		delete(s, v)
		return true
	}
	return false
}

func (s Set[T]) Contains(v T) bool {
	_, found := s[v]
	return found
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	return lo.Assign(s, s2)
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	return lo.PickBy(s, func(k T, v struct{}) bool {
		return s2.Contains(k)
	})
}

func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	return lo.PickBy(s, func(k T, _ struct{}) bool {
		return !s2.Contains(k)
	})
}
