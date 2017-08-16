package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/pkg/profile"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
database{
	[]mark,
	students: [...]
}
*/
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

	var d Database
	jsonFile, err := ioutil.ReadFile("data.json")
	check(err)
	if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}

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
	//fmt.Println(d.Marks[0].Class)

}
