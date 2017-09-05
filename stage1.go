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
func averageMarkTwo(d Database) {
	type Suburbs []struct {
		suburb      string
		student_ids []int
	}
	var suburbs Suburbs
	for i := 0; i < len(d.Students); i++ {
		aSuburb := d.Students[i].Suburb
		fmt.Println(aSuburb)
		for j := 0; j < len(suburbs); j++ {
			fmt.Println("it lives")
			if aSuburb == suburbs[j].suburb {
				break
			} else if aSuburb != suburbs[j].suburb && j == len(suburbs) {

				suburbs = append(suburbs)

			}
		}
	}
	fmt.Println(suburbs[0].suburb)
}
