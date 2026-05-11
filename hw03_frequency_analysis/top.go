package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type freq struct {
	s   string
	cnt int
}

func Top10(s string) []string {
	words := strings.Fields(s)

	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	slice := make([]freq, 0, len(m))
	for k, cnt := range m {
		slice = append(slice, freq{k, cnt})
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].cnt != slice[j].cnt {
			return slice[i].cnt > slice[j].cnt
		}
		return slice[i].s < slice[j].s
	})
	res := make([]string, 0, min(len(slice), 10))
	for i := 0; i < min(len(slice), 10); i++ {
		res = append(res, slice[i].s)
	}
	return res
}
