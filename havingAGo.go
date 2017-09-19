/*package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	db, err := sql.Open("postgres", "user=postgres password=password dbname=AlexDB sslmode=disable port=5433")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping() // for testing purposes

	if err != nil {
		fmt.Println("postgres err thing")
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS myTable(msg varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(
		"INSERT INTO myTable(msg) VALUES('hello SIR!')") //insterts a new row with hello world
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.RowsAffected())

	rowCount, err := res.RowsAffected() //informs of number of rows affected
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM myTable")
	var msg string
	for rows.Next() {
		err = rows.Scan(&msg)
		// TO DO: handle err
		log.Printf("Got row: %q", msg)
	}
	rows.Close()

	_, err = db.Exec(
		"DROP TABLE IF EXISTS stuff")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS stuff(column1 varchar(50), column2 int)")
	if err != nil {
		log.Fatal(err)
	}

	res1, err := db.Exec(
		"INSERT INTO stuff(column1, column2) VALUES('Stuff and Things', 5)") //insterts a new row with hello world
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res1.RowsAffected())

	newRows, err := db.Query("SELECT * FROM stuff")
	// TODO: handle err
	var (
		column1 string
		column2 int
	)
	for newRows.Next() {
		err = rows.Scan(&column1, &column2)
		// TO DO: handle err
		log.Printf("Values: %d, %q", column1, column2)
	}
	newRows.Close()

	return

	defer profile.Start(profile.MemProfile, profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//var dat map[string]interface{}
	time.Sleep(1 * time.Second)

	var d Database
	jsonFile, err := ioutil.ReadFile("data.json")
	check(err)
	if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}
	//fmt.Println(d.Marks[0].Class)
	dataOverveiw(d)
}
func dataOverveiw(d Database) {
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

*/