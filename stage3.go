package main

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
	defer profile.Start(profile.MemProfile, profile.CPUProfile, profile.ProfilePath(".")).Stop()
	time.Sleep(1 * time.Second)
	t0 := time.Now()

	var d Database
	jsonFile, err := ioutil.ReadFile("long_student_data.json")
	check(err)
	if err := json.Unmarshal(jsonFile, &d); err != nil {
		panic(err)
	}
	t1 := time.Now()
	//fmt.Println(d.Marks)
	t2 := time.Now()
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
		"DROP TABLE IF EXISTS Marks")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS Marks(StudentID int, Class varchar(50), Mark real)")
	if err != nil {
		log.Fatal(err)
	}
	for i :=0; i < len(d.Marks); i++{
		_, err := db.Exec(
			"INSERT INTO Marks(StudentID, Class, Mark) VALUES($1, $2, $3)", d.Marks[i].StudentID, d.Marks[i].Class, d.Marks[i].Mark)  //insterts a new row with hello world
		if err != nil {
			log.Fatal(err)
		}
	}


//Create a new table for Average marks 
	_, err = db.Exec(
		"DROP TABLE IF EXISTS OverView")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS OverView(Suburb varchar(50), Class varchar(50), Mark real)")
	if err != nil {
		log.Fatal(err)
	}

	
	
	_, err = db.Exec(
		"DROP TABLE IF EXISTS Students")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS Students(StudentID int, FirstName varchar(50), LastName varchar(50), Age int, PhoneNumebr varchar(50), Suburb  varchar(50), City varchar(50))")
	if err != nil {
		log.Fatal(err)
	}
	for i :=0; i < len(d.Students); i++{
		_, err := db.Exec(
			"INSERT INTO Students(StudentID, FirstName, LastName, Age, PhoneNumebr, Suburb, City) VALUES($1, $2, $3, $4, $5, $6, $7)", d.Students[i].StudentID, d.Students[i].FirstName, d.Students[i].LastName, d.Students[i].Age, d.Students[i].PhoneNumebr, d.Students[i].Suburb, d.Students[i].City)
		if err != nil {
			log.Fatal(err)
		}
	}
	

	joinMS, err := db.Query("SELECT * FROM Students NATURAL JOIN Marks ORDER BY Class, Suburb, Mark")
		for joinMS.Next(){
			var (
				StudentID int 
				FirstName string
				LastName string
				Age int
				PhoneNumebr string
				Suburb string
				City string
				Class string
				Mark float64
			)
			err = joinMS.Scan(&StudentID, &FirstName, &LastName, &Age, &PhoneNumebr, &Suburb, &City, &Class, &Mark)
			
			
				_, err := db.Exec(
					"INSERT INTO OverView(Class, Suburb, Mark) VALUES($1, $2, $3)", Class, Suburb, Mark) 
				if err != nil {
					log.Fatal(err)
				}
				
			
			

			//fmt.Println(Class)
			//fmt.Println(Suburb)
			//fmt.Println(Mark)
		}
		
	t3 := time.Now()		
/*
	_, err = db.Exec("SELECT DISTINCT Class, Suburb, AVG(Mark) FROM OverView GROUP BY Class, Suburb")
	if err != nil {
		log.Fatal(err)
	}
*/
	t4 := time.Now()
	var stringSlice []string
	groupdOrVw, err := db.Query("SELECT DISTINCT Class, Suburb, AVG(Mark) FROM OverView GROUP BY Class, Suburb ORDER BY AVG(Mark) DESC")
	for groupdOrVw.Next(){
		var (
			Class string
			Suburb string
			Mark float64
		)
		
		err = groupdOrVw.Scan(&Class, &Suburb, &Mark)
		var MarkString string = Class + " " + Suburb + " " + strconv.FormatFloat(Mark, 'f', 1, 64)
		stringSlice = append(stringSlice, MarkString)
		//fmt.Println(MarkString)
		//fmt.Println(Suburb)
		//fmt.Println(Mark)
	}

	t5 :=time.Now()
	groupdOrVw.Close()
	for i :=0; i < len(stringSlice); i++{
		fmt.Println(stringSlice[i])
	}

	fmt.Println("Time Used to load .json", t1.Sub(t0))
	fmt.Println("Time Used to create databases", t3.Sub(t2))
	fmt.Println("Time Used to retrieve/query database ", t5.Sub(t4))
}















 /*
	rows, err := db.Query("SELECT * FROM Students")
	markrows, err := db.Query("SELECT * FROM Marks")
	

	for rows.Next() {
		var (
			StudentID int 
			FirstName string
			LastName string
			Age int
			PhoneNumebr string
			Suburb string
			City string
		)
	err = rows.Scan(&StudentID, &FirstName, &LastName, &Age, &PhoneNumebr, &Suburb, &City)


	var OverviewString string
	//fmt.Println("Values:", StudentID, FirstName, LastName, Age, PhoneNumebr, Suburb, City)
	//var currentID int = StudentID
	OverviewString += FirstName + " " + LastName
	//Sfmt.Println(OverviewString)

			for markrows.Next(){
				var (
					StudentID2 int
					Class string
					Mark float64
				)
				err = markrows.Scan(&StudentID2, &Class, &Mark)
				OverviewString += Class + " " + strconv.FormatFloat(Mark, 'f', 1, 64) + " "
				//fmt.Println(OverviewString)
			}

	}
	rows.Close() 
	markrows.Close()
	var unqueSubs []string
	var unqueClas []string
	//var idsForSubs [][]int
	suburbs, err := db.Query("SELECT DISTINCT Suburb FROM Students")
	for suburbs.Next(){
		var subub string
		err = suburbs.Scan(&subub)
		unqueSubs = append(unqueSubs, subub)
		//fmt.Println(subub)
		
	}
	suburbs.Close()

	classes, err := db.Query("SELECT DISTINCT class FROM Marks")
	for classes.Next(){
		var aclass string
		err = classes.Scan(&aclass)
		unqueClas = append(unqueClas, aclass)
		//fmt.Println(aclass)
		
	}
	suburbs.Close()
	//fmt.Println(unqueClas)
	//fmt.Println(unqueSubs)
	
	
		studentBySuburb, err := db.Query("SELECT StudentID, Suburb FROM Students")
		for studentBySuburb.Next() {
			
			var (
				StudentID int
				Suburb string
			)
			err = studentBySuburb.Scan(&StudentID, &Suburb)
			//fmt.Println(Suburb)
			//fmt.Println(StudentID)
		}
		studentBySuburb.Close()

	

	var (
	StudentID int
	Class string
	Mark float64
	)
	for markrows.Next(){
		err = markrows.Scan(&StudentID, &Class, &Mark)
		OverviewString += Class + strconv.FormatFloat(Mark,'E', -1, 64)
	}*/
	
