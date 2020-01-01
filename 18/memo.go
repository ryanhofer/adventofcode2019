package main

import (
	"sort"

	"github.com/ryanhofer/adventofcode2019/geom"
)

type cacheKey struct {
	pos  geom.Vec
	keys string
}

var cache map[cacheKey]int

func Cached(pos geom.Vec, keys string) (int, bool) {
	k := toCacheKey(pos, keys)
	score, ok := cache[k]
	return score, ok
}

func Remember(pos geom.Vec, keys string, score int) {
	k := toCacheKey(pos, keys)
	cache[k] = score
}

func toCacheKey(pos geom.Vec, keys string) cacheKey {
	s := []rune(keys)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return cacheKey{
		pos:  pos,
		keys: string(s),
	}
}
