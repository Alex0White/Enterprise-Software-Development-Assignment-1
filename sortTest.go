/*package main

import (
	"fmt"
	"sort"
)

type MarkBySuburb struct {
	mark          float64
	suburbSubject string
}

type By func(p1, p2 *MarkBySuburb) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(markBySuburbs []MarkBySuburb) {
	ps := &markSorter{
		markBySuburbs: markBySuburbs,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}



type markSorter struct {
	markBySuburbs []MarkBySuburb
	by			func(p1, p2 *MarkBySuburb) bool
}

func (s *markSorter) Len() int {
	return len(s.markBySuburbs)
}

func (s *markSorter) Swap(i, j int) {
	s.markBySuburbs[i], s.markBySuburbs[j] = s.markBySuburbs[j], s.markBySuburbs[i]
}

func (s *markSorter) Less(i, j int) bool {
	return s.by(&s.markBySuburbs[i], &s.markBySuburbs[j])
}

var markBySuburbs = []MarkBySuburb{
	{0.01, "sup"},
	{0.09, "yo"},
	{1.0, "zebra"},
	{0.0001, "small"},
}
/*










var avgMarks []float64 // storeing avgMarks in slice so that can be sorted by number
for i, j := range marksMap {
	j[0] = (j[0] / j[1])
	a := MarkBySuburb{j[0], i}
	markSorter.markBySuburbs = append(markSorter.markBySuburbs, a)
	fmt.Println(i)
	fmt.Println(j[0])
}

sort.Float64s(avgMarks)

for i := range avgMarks {

	for w, j := range marksMap {
		if avgMarks[i] == j[0] {
			fmt.Println(w)
			fmt.Println(j[0])
			break
		}
	}
}
fmt.Println(markSorter)

type By func(p1, p2 *MarkBySuburb) bool




func main(){

	mark := func(p1, p2 *MarkBySuburb) bool {
		return p1.mark < p2.mark
	}

	By(mark).Sort(markBySuburbs)
	fmt.Println("By marks:", markBySuburbs)


}
*/