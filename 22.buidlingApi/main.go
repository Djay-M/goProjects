package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// models
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// helper functions
func validateCourse(course *Course) bool {
	return course.CourseName == ""
}

// dummy data data seeder
func seedData() {
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Node Js",
		CoursePrice: 500,
		Author: &Author{
			FullName: "dev tester",
			Website:  "www.devTester.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Python",
		CoursePrice: 1000,
		Author: &Author{
			FullName: "dev tester Python",
			Website:  "www.devTesterPython.com",
		},
	})
}

func main() {
	fmt.Println("Welcome to go module, for building a API")
	// creatig a router
	router := mux.NewRouter()

	// route handling
	// server status apis
	router.HandleFunc("/", getServerStatus).Methods("GET")
	router.HandleFunc("/status", getServerStatus).Methods("GET")
	router.HandleFunc("/api/status", getServerStatus).Methods("GET")

	// GET apis
	router.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/api/course/{id}", getCourseById).Methods("GET")

	// POST apis
	router.HandleFunc("/api/course", createCourse).Methods("POST")

	// PUT apis
	router.HandleFunc("/api/course/{id}", updateCourseById).Methods("PUT")

	// DELETE apis
	router.HandleFunc("/api/course/{id}", deleteCourseById).Methods("DELETE")

	// seeding dummyData
	seedData()

	// starting a server
	fmt.Println("statring server at port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

// Controllers
func getServerStatus(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("Server is up and running !!!"))
}

func getAllCourses(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses)
}

func getCourseById(writer http.ResponseWriter, reader *http.Request) {
	params := mux.Vars(reader)
	writer.Header().Set("Content-Type", "application/json")

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
	json.NewEncoder(writer).Encode("No course found for the id " + params["id"])
}

func createCourse(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())

	// check for nil reader.body
	if reader.Body == nil {
		json.NewEncoder(writer).Encode("No data found in req.body")
	}

	var course Course
	_ = json.NewDecoder(reader.Body).Decode(&course)

	if validateCourse(&course) {
		json.NewEncoder(writer).Encode("Empty JSON found in req.body")
	}

	course.CourseId = strconv.Itoa((rand.Intn(100)))
	courses = append(courses, course)
	json.NewEncoder(writer).Encode(course)
}

func updateCourseById(writer http.ResponseWriter, reader *http.Request) {
	params := mux.Vars(reader)
	writer.Header().Set("Content-Type", "application/json")

	if reader.Body == nil {
		json.NewEncoder(writer).Encode("req.body is nil, please check")
	}

	var course Course
	json.NewDecoder(reader.Body).Decode(&course)

	if validateCourse(&course) {
		json.NewEncoder(writer).Encode("Empty json foun in req.body, please check")
	}

	for index, curCourse := range courses {
		if curCourse.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			courses = append(courses, course)
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
	json.NewEncoder(writer).Encode("No course found with the given id" + params["id"])
}

func deleteCourseById(writer http.ResponseWriter, reader *http.Request) {
	params := mux.Vars(reader)
	writer.Header().Set("Content-Type", "application/json")

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(writer).Encode(course)
			return
		}
	}

	json.NewEncoder(writer).Encode("No course found by id" + params["id"])
}
