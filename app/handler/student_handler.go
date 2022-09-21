package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-restful-api/app/model"
	"net/http"
	"strconv"
)

var students = []model.Student{
	{Id: 1, FullName: "Student 1"},
	{Id: 2, FullName: "Student 2"},
	{Id: 3, FullName: "Student 3"},
	{Id: 4, FullName: "Student 4"},
	{Id: 5, FullName: "Student 5"},
	{Id: 6, FullName: "Student 6"},
	{Id: 7, FullName: "Student 7"},
	{Id: 8, FullName: "Student 8"},
	{Id: 9, FullName: "Student 9"},
	{Id: 10, FullName: "Student 10"},
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id_str := vars["id"] // read path's value from url
	var student = model.Student{}

	for i := 0; i < len(students); i++ {
		student_id, _ := strconv.Atoi(student_id_str)
		if students[i].Id == student_id {
			student = students[i]
			break
		}
	}
	if student == (model.Student{}) {
		respondError(w, http.StatusNotFound, "Not found")
		return
	}
	respondJson(w, http.StatusOK, student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	student := model.Student{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&student); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var newStudents = make([]model.Student, len(students)+1)
	for i := 0; i < len(students); i++ {
		newStudents[i] = students[i]
	}
	newStudents[len(students)] = student
	students = newStudents
	respondJson(w, http.StatusCreated, student)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	student_id_str := vars["id"] // read path's value from url

	studentReq := model.Student{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&studentReq); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var student = model.Student{}
	var markedIndex = -1
	for i := 0; i < len(students); i++ {
		student_id, _ := strconv.Atoi(student_id_str)
		if students[i].Id == student_id {
			student = students[i]
			markedIndex = i
			break
		}
	}
	if student == (model.Student{}) {
		respondError(w, http.StatusNotFound, "Not found")
		return
	}

	studentReq.Id = student.Id
	students[markedIndex] = studentReq
	respondJson(w, http.StatusCreated, studentReq)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id_str := vars["id"] // read path's value from url
	var student = model.Student{}

	var markedIndex = -1
	for i := 0; i < len(students); i++ {
		student_id, _ := strconv.Atoi(student_id_str)
		if students[i].Id == student_id {
			student = students[i]
			markedIndex = i
			break
		}
	}
	if student == (model.Student{}) {
		respondError(w, http.StatusNotFound, "Not found")
		return
	}

	var newStudents = make([]model.Student, len(students)-1)
	var k = 0
	for i := 0; i < len(students); i++ {
		if markedIndex == i {
			continue
		}
		newStudents[k] = students[i]
		k++
	}
	students = newStudents

	respondJson(w, http.StatusNoContent, nil)
}
