package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	defer profile.Start(profile.MemProfile, profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//var dat map[string]interface{}
	time.Sleep(1 * time.Second)
	t2 := time.Now()
	var d Database
	jsonFile, err := ioutil.ReadFile("data.json")
	check(err)
	if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}
	t3 := time.Now()
	//fmt.Println(d.Marks[0].Class)
	t0 := time.Now()
	dataOverveiw(d)
	t1 := time.Now()
	fmt.Println("Used time: ", t1.Sub(t0))
	fmt.Println("Used time: ", t2.Sub(t3))
	//UnqueStringItemizer(d)
	averageMarkTwo(d)

}

func dataOverveiw(d Database) {

	fmt.Println(d.Students[0].FirstName)
	for i := 0; i < len(d.Students); i++ {
		var studentMarks string = ""
		var studentNames string = d.Students[i].FirstName + " " + d.Students[i].LastName + " "
		for chur := 0; chur < len(d.Marks); chur++ {

			if i == d.Marks[chur].StudentID {
				FlStr := strconv.FormatFloat(d.Marks[chur].Mark, 'f', 2, 64)
				studentMarks += d.Marks[chur].Class + ": " + FlStr + " "
			}
		}
		fmt.Println(studentNames + studentMarks)
	}
}

/*
func averageMark(d Database) {
	type suburb []struct{ //use a map
		 singleSuburb string

	}


	var suburbStruct suburb
	var aSuburb string





	for i := 0; i < len(d.Students); i++ {
		aSuburb := d.Students[i].Suburb
		 if len(suburbStruct) == 0 {
			 suburbStruct[].singleSuburb = append(suburbStruct.singleSuburb ,aSuburb)
		 }
		for i := 0; i < len(suburbStruct); i++ {
			if aSuburb == suburbStruct[i]{
				break
			}else if aSuburb != suburbStuct[i] && i == len(suburbStruct){
				suburb.singleSuburb := append(suburb.singleSuburb ,aSuburb)
			}
		}
	}


}
*/
<<<<<<< HEAD
func isInSuburbs(s []string, aS string) bool {
	var answer bool
	answer = true
	for i := 0; i < len(s); i++ {
		if s[i] == aS {
			answer = true
		} else if s[i] != aS {
			answer = false
		}

	}
	return answer
}
func averageMarkTwo(d Database) {
	/*
		type Suburbs []struct {
			suburb      string
			student_ids []int
		}
	*/

	//student_ids := [][]int{}

	// These are the first two rows.
	suburbs := []string{}

	//row1 := []int{1, 2, 3}
	//row2 := []int{4, 5, 6}

	// Append each row to the two-dimensional slice.
	//values = append(values, row1)
	//values = append(values, row2)
	//fmt.Println(values[0])
	//stringThing := []string{"lol"}
	//values = append(values, stringThing)
	// if [this array contains suburb] return true
	//if len(suburbs) == 0 {
	//	suburbs = append(suburbs, d.Students[0].Suburb)
	//}

	for i := 0; i < len(d.Students); i++ {
		aSuburb := d.Students[i].Suburb
		fmt.Println(aSuburb)
		isInSuburbs(suburbs, aSuburb)
		for j := 0; j < len(suburbs); j++ {
			//fmt.Println("it lives")

			if aSuburb != suburbs[j] {

				suburbs = append(suburbs, aSuburb)
=======
func UnqueStringItemizer(stringSlice []string, itemString string) bool {
	var isUnque = true
	for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == itemString {
			isUnque = false
		}
	}
	return isUnque
}

func averageMarkTwo(d Database) {
	var studentIds [][]int // this is a array of the studentIds which corrospond to suburbs
	var suburbs []string
	var classes []string
	var marksMap = map[string]float64{}
	type aMark struct {
		markID  int
		suburb  string
		class   string
		markAvg int
	}
	type overveiw struct {
		suburbAvgMark []struct {
			suburb  string
			Classes []struct {
				classType string
				marks     []int
				marksAvg  int
			}
		}
	}

	for i := 0; i < len(d.Students); i++ {
		aSuburb := d.Students[i].Suburb
		if UnqueStringItemizer(suburbs, aSuburb) {
			suburbs = append(suburbs, aSuburb)
		}
	}
	fmt.Println(suburbs)
	for i := 0; i < len(d.Marks); i++ {
		aClass := d.Marks[i].Class
		if UnqueStringItemizer(classes, aClass) {
			classes = append(classes, aClass)
		}
	}
	/*
		fmt.Println(classes)
		fmt.Println(len(suburbs))
		fmt.Println("d.Students lenght", len(d.Students))*/
	studentIds = make([][]int, len(suburbs))
	for i := 0; i < len(d.Students); i++ {

		for j := 0; j < len(suburbs); j++ {

			if suburbs[j] == d.Students[i].Suburb {

				studentIds[j] = append(studentIds[j], d.Students[i].StudentID)
				break
			}
		}
	}
	//fmt.Println(studentIds)
	for i := 0; i < len(studentIds); i++ { //suburbs [Tamatea Mayfair Mahora Onekawa Akina Greenmeadows Taradale Ahuriri]
		//new := aMark{i, suburbs[i], "", 0}
		for j := 0; j < len(studentIds[i]); j++ { //a student from Tamatea marks[programming hardware packages etc]
			for m := range d.Marks { // get this student and match him up with marks struct should match with 4 Marks objects which each contain a matching student ID
				for c := range classes {

					if d.Marks[m].StudentID == studentIds[i][j] && classes[c] == d.Marks[m].Class { //this is a mark for a specific class within a specific suburb
						var mapkey string = suburbs[i] + " " + classes[c]
						num, ok := marksMap[mapkey]
						if ok {
							marksMap[mapkey] = (num + d.Marks[m].Mark) / 2
						} else {
							marksMap[mapkey] = d.Marks[m].Mark
						}
					}
				}
>>>>>>> 242d5e4af92cec55e4f393a41498050751d396d7
			}
		}

	}
<<<<<<< HEAD
	//fmt.Println(suburbs[0])

=======
	fmt.Println(marksMap)
	fmt.Println(len(marksMap))
>>>>>>> 242d5e4af92cec55e4f393a41498050751d396d7
}
