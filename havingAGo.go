package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type marks struct {
	class      string  `json:"class"`
	student_id int     `json:"student_id"`
	marks      float32 `json:"mark"`
}

type students struct {
	student_id   int    `json:"student_id"`
	first_name   string `json:"first_name"`
	last_name    string `json:"last_name"`
	age          int    `json:"age"`
	phone_number string `json:"phone_numebr"`
	suburb       string `json:"suburb"`
	city         string `json:"city"`
}
type database struct {
	theMarks    []marks    `json:"marks"`
	theStudents []students `json:"students"`
}

func main() {

	//var dat map[string]interface{}
	var dat map[string]interface{}

	jsonFile, err := ioutil.ReadFile("data.json")
	//fmt.Println(jsonFile)
	check(err)
	//fmt.Print(string(jsonFile))

	//var dat map[string]interface{}

	/*if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}*/
	if err := json.Unmarshal(jsonFile, &dat); err != nil {
		panic(err)
	}
	marksArray := dat["marks"]
	//database := dat["marks"], dat["students"]
	//database{dat["marks"], dat["students"]}
	//fmt.Println(dat)
	fmt.Println(marksArray)

	/*
		res1D := &Response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}

		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))
	*/

}
