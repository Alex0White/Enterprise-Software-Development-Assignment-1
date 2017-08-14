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
type mark struct {
	class      string
	student_id int
	marks      float32
}

/*
{
        "student_id": 2,
        "first_name": "Denis",
        "last_name": "Goublier",
        "age": 23,
        "phone_numebr": "024 699 0879",
        "suburb": "Mahora",
        "city": "Hastings"
}
*/
func main() {

	jsonFile, err := ioutil.ReadFile("data.json")
	check(err)
	//fmt.Print(string(jsonFile))

	var dat map[string]interface{}

	if err := json.Unmarshal(jsonFile, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	/*
		res1D := &Response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}

		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))
	*/

}
