package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}

type Course struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Credit int    `json:"credit"`
}

type Student struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Major   string   `json:"major"`
	Address Address  `json:"address"`
	Courses []Course `json:"courses"`
}

var students = []Student{
	{
		ID:    1,
		Name:  "Alice Smith",
		Age:   22,
		Major: "Software Engineering",
		Address: Address{
			Street:     "123 Main St",
			City:       "Anytown",
			State:      "Anystate",
			PostalCode: "12345",
		},
		Courses: []Course{
			{Code: "CS101", Name: "Introduction to Programming", Credit: 3},
			{Code: "CS201", Name: "Data Structures", Credit: 4},
		},
	},
	{
		ID:    2,
		Name:  "AAAAAAAAAAAA",
		Age:   22,
		Major: "Software Engineering",
		Address: Address{
			Street:     "123 Main St",
			City:       "Anytown",
			State:      "Anystate",
			PostalCode: "12345",
		},
		Courses: []Course{
			{Code: "CS101", Name: "Introduction to Programming", Credit: 3},
			{Code: "CS201", Name: "Data Structures", Credit: 4},
		},
	},
	{
		ID:    3,
		Name:  "BBBBBBBBBB",
		Age:   22,
		Major: "Software Engineering",
		Address: Address{
			Street:     "123 Main St",
			City:       "Anytown",
			State:      "Anystate",
			PostalCode: "12345",
		},
		Courses: []Course{
			{Code: "CS101", Name: "Introduction to Programming", Credit: 3},
			{Code: "CS201", Name: "Data Structures", Credit: 4},
		},
	},
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(students)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func getStudentByid(w http.ResponseWriter, r *http.Request) {
	string_id := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(string_id)

	if err != nil {
		http.Error(w, "Student not FOund", http.StatusBadRequest)
		return
	}

	for _, student := range students {
		if student.ID == id {
			jsonData, err := json.Marshal(student)
			if err != nil {
				http.Error(w, "error Encoding data", http.StatusInternalServerError)
			}
			w.Write(jsonData)
			return
		}
	}
}

func studentPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error Reading the Body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var addStudent Student
	err = json.Unmarshal(body, &addStudent)
	if err != nil {
		http.Error(w, "Json format is Invalid", http.StatusBadRequest)
	}

	addStudent.ID = len(students) + 1
	students = append(students, addStudent)

	jsonData, _ := json.Marshal(students)
	w.Write(jsonData)

}

func studentPut(w http.ResponseWriter, r *http.Request) {
	stringID := strings.TrimPrefix(r.URL.Path, "/students/update/")
	id, err := strconv.Atoi(stringID)

	var updateStudent Student

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error Reading the Body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &updateStudent)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, student := range students {
		if student.ID == id {
			updateStudent.ID = id
			students[i] = updateStudent

			jsonData, _ := json.Marshal(students)
			w.Write(jsonData)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {

	stringID := strings.TrimPrefix(r.URL.Path, "/students/delete/")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)

			jsonData, _ := json.Marshal(students)
			w.Write(jsonData)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/students/create", studentPost)
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/students/{id}", getStudentByid)
	http.HandleFunc("/students/update/{id}", studentPut)
	http.HandleFunc("/students/delete/{id}", deleteStudent)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
