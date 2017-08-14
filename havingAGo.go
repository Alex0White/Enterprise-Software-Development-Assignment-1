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
	class string
	student_id int
	marks float32
}

type students struct {
	student_id int
	first_name string
	last_name string
	age int
	phone_numbr string
	suburb string
	city string
}
type database struct {
	markInfo []marks
	studentInfo []students
}

func main() {

	var dat map[string]interface{}
	var d database


	jsonFile, err := ioutil.ReadFile("data.json")
	check(err)
	//fmt.Print(string(jsonFile))

	//var dat map[string]interface{}

	/*if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}*/
	if err := json.Unmarshal(jsonFile, &dat); err != nil {
		panic(err)
	}


	fmt.Println(d)
	fmt.Println(dat)

	/*
		res1D := &Response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}

		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))
	*/

}
