package main

import "sort"

type Pb struct {
	score int
	name  string
}
type Byid []Pb

func (s Byid) Len() int           { return len(s) }
func (s Byid) Less(i, j int) bool { return s[i].score < s[j].score }
func (s Byid) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func main() {
	S := Byid{{1, "123"}, {123, "123"}, {3, "45"}}
	sort.Sort(S)
}
