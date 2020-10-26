package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	stringSlice := SliceStream{data: []string{"a", "b", "c", "1", "D"}}
	stringSlice.Map(strings.ToUpper).Filter(isLetter)
	fmt.Println(stringSlice.data)
}

type any interface{}

type Stream interface {
	Map(m Mapper) Stream
	Filter(p Predicate) Stream
	Reduce(a Accumulator) any
}

type SliceStream struct {
	data []string
}

func isLetter(input string) bool {
	r := []rune(input)
	return unicode.IsLetter(r[0])
}

func (s *SliceStream) Map(m Mapper) *SliceStream {
	for i, e := range s.data {
		s.data[i] = m(e)
	}
	return s
}

func (s *SliceStream) Filter(p Predicate) *SliceStream {
	for i := range s.data {
		if !p(s.data[i]) {
			s.data = remove(s.data, i)
			i = i - 1
		}
	}
	return s
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func (s *SliceStream) Reduce(a Accumulator) string {
	return a(concat(s), "")
}

func concat(s *SliceStream) string {
	return strings.Join(s.data[:], ",")
}

type Accumulator func(x, y string) string
type Mapper func(x string) string
type Predicate func(x string) bool
