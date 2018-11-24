package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// IGts zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

// UnionWith sets s to the union of s and t.

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)

			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}

	return len
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var r IntSet
	r = *s
	return &r
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func main() {
	var x, y IntSet
	var z, copied *IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(x.Len()) //4
	fmt.Println(y.Len()) //2

	y.Add(42)
	x.Add(12)
	y.Add(34)
	x.Add(23)
	x.Add(91)
	y.Add(67)

	fmt.Println(x.Len()) //6
	fmt.Println(y.Len()) //4

	fmt.Println(x.String()) // "{1 9 12 23 42 91 144}"
	fmt.Println(y.String()) // "{9 34 42 67}"

	x.Remove(9)
	fmt.Println(x.String()) // "{1 12 23 42 91 144}"
	x.Remove(997899)
	fmt.Println(x.String()) // "{1 12 23 42 91 144}"

	z = x.Copy()
	copied = &x
	fmt.Println(z.String())      // "{1 12 23 42 91 144}"
	fmt.Println(copied.String()) // "{1 12 23 42 91 144}"
	x.Add(919)
	fmt.Println(x.String())      // "{1 12 23 42 91 144 919}"
	fmt.Println(copied.String()) // "{1 12 23 42 91 144 919}"
	fmt.Println(z.String())      // "{1 12 23 42 91 144}"

	x.Clear()
	fmt.Println(x.String())      // "{}"
	fmt.Println(copied.String()) // "{}"
	fmt.Println(z.String())      // "{1 12 23 42 91 144}"

	z.AddAll(1, 2, 78, 34, 67)
	fmt.Println(z.String()) // "{1 2 12 23 34 42 67 78 91 144}"

}
