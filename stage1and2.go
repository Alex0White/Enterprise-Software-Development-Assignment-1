package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"time"
	_ "github.com/lib/pq"
	"github.com/pkg/profile"
	
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Database struct {
	Marks []struct {
		StudentID int     `json:"student_id"`
		Class     string  `json:"class"`
		Mark      float64 `json:"mark"`
	} `json:"marks"`
	Students []struct {
		StudentID   int    `json:"student_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Age         int    `json:"age"`
		PhoneNumebr string `json:"phone_numebr"`
		Suburb      string `json:"suburb"`
		City        string `json:"city"`
	} `json:"students"`
}

type MarkBySuburb struct {
	mark          float64
	suburbSubject string
}

type By func(p1, p2 *MarkBySuburb) bool//sorting stuff

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(markBySuburbs []MarkBySuburb) {//sorting stuff
	ps := &markSorter{
		markBySuburbs: markBySuburbs,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type markSorter struct {//sorting stuff
	markBySuburbs []MarkBySuburb
	by			func(p1, p2 *MarkBySuburb) bool
}

func (s *markSorter) Len() int {//sorting stuff
	return len(s.markBySuburbs)
}

func (s *markSorter) Swap(i, j int) { //sorting stuff
	s.markBySuburbs[i], s.markBySuburbs[j] = s.markBySuburbs[j], s.markBySuburbs[i]
}

func (s *markSorter) Less(i, j int) bool { //sorting stuff
	return s.by(&s.markBySuburbs[i], &s.markBySuburbs[j])
}

func main() {
	defer profile.Start(profile.MemProfile, profile.CPUProfile, profile.ProfilePath(".")).Stop()
	time.Sleep(1 * time.Second)
	t2 := time.Now()
	var d Database
	jsonFile, err := ioutil.ReadFile("long_student_data.json")
	check(err)
	if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}
	t3 := time.Now()
	//fmt.Println(d.Marks[0].Class)
	t0 := time.Now()
	dataOverveiw(d)
	t1 := time.Now()
	fmt.Println("Time used for dataOverveiw: ", t1.Sub(t0))
	fmt.Println("Time used for reading long_student_data.json: ", t3.Sub(t2))
	//UnqueStringItemizer(d)
	t5 := time.Now()
	stage2 := stage2(d)
	t6 := time.Now()
	
	for i := len(stage2)-1; i >= 0; i-- {
		fmt.Println(stage2[i].suburbSubject + " " + strconv.FormatFloat(stage2[i].mark, 'f', 1, 64))
	}
	fmt.Println("Time Used for stage2: ", t6.Sub(t5))
}

func dataOverveiw(d Database) {

	//fmt.Println(d.Students[0].FirstName)
	for i := 0; i < len(d.Students); i++ { //loop one
		var studentMarks string = ""
		var studentNames string = d.Students[i].FirstName + " " + d.Students[i].LastName + " "
		for chur := 0; chur < len(d.Marks); chur++ { //loop two

			if i == d.Marks[chur].StudentID {
				FlStr := strconv.FormatFloat(d.Marks[chur].Mark, 'f', 2, 64)
				studentMarks += d.Marks[chur].Class + ": " + FlStr + " "
			}
		}
		fmt.Println(studentNames + studentMarks)
	}
}

func UnqueStringItemizer(stringSlice []string, itemString string) bool {  // this takes a string and compares it with a array of strings and returns true if the string does not appear in the array 
	var isUnque = true
	for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == itemString {
			isUnque = false
		}
	}
	return isUnque
}



func stage2(d Database) []MarkBySuburb {
	var studentIds [][]int // this is a array of the studentIds which corrospond to suburbs 
	var suburbs []string
	var classes []string // this contains the names of class subject
	var marksMap = map[string][2]float64{} // this contains a string which decares suburb and a class/subject this is the key. the items are arrays with lengths of 2 
	// the first item of the float array is a sum of a bunch of marks and the second item reprosents the number of marks in the first item  

	for i := 0; i < len(d.Students); i++ { // gets unique suburbs and appends them to suburbs[]
		aSuburb := d.Students[i].Suburb
		if UnqueStringItemizer(suburbs, aSuburb) {
			suburbs = append(suburbs, aSuburb)
		}
	}

	for i := 0; i < len(d.Marks); i++ { // gets unique classes/subjects and appends them to classes[]
		aClass := d.Marks[i].Class
		if UnqueStringItemizer(classes, aClass) {
			classes = append(classes, aClass)
		}
	}

	studentIds = make([][]int, len(suburbs)) // intiallizes the length of the list
	for i := 0; i < len(d.Students); i++ {

		for j := 0; j < len(suburbs); j++ {

			if suburbs[j] == d.Students[i].Suburb {

				studentIds[j] = append(studentIds[j], d.Students[i].StudentID)
				break
			}
		}
	}
	var bigmarks float64 //used to add marks to markMap
	for i := 0; i < len(studentIds); i++ { //suburbs [Tamatea Mayfair Mahora Onekawa Akina Greenmeadows Taradale Ahuriri]
		//new := aMark{i, suburbs[i], "", 0}
		for j := 0; j < len(studentIds[i]); j++ { //a student from Tamatea marks[programming hardware packages etc] looping though i and j should equal 100 loops 
			for m := range d.Marks { // get this student and match him up with marks struct should match with 4 Marks objects which each contain a matching student ID 400 loops
				for c := range classes { //4 loops total of 160000 loops

					if d.Marks[m].StudentID == studentIds[i][j] && classes[c] == d.Marks[m].Class { //this is a mark for a specific class within a specific suburb
						var mapkey string = suburbs[i] + " " + classes[c]
						num, ok := marksMap[mapkey] // this finds items using a map key from the marksMap
						if ok {
							bigmarks = (num[0] + d.Marks[m].Mark)

							num[0] = bigmarks
							num[1] = (num[1] + 1)
							marksMap[mapkey] = num
						} else {
							num[0] = d.Marks[m].Mark
							num[1] = (num[1] + 1)
							marksMap[mapkey] = num

						}
					}
				}
			}
		}
	}
	
	
	var markBySuburbs = []MarkBySuburb{}
	
	var avgMarks []float64 // storeing avgMarks in slice so that can be sorted by number
	for i, j := range marksMap {
		j[0] = (j[0] / j[1])
		a := MarkBySuburb{j[0], i}
		markBySuburbs = append(markBySuburbs, a)
	}

	sort.Float64s(avgMarks)
	mark := func(p1, p2 *MarkBySuburb) bool {
		return p1.mark < p2.mark
	}
	By(mark).Sort(markBySuburbs)
	
	return markBySuburbs
}
