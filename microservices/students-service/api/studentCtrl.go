package api

import (
	"encoding/json"
	"go-projects/microservices/students-service/core"
	"go-projects/microservices/students-service/dal"
	"go-projects/microservices/students-service/model"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
)

func AddNewStudent(w http.ResponseWriter, r *http.Request) {
	session := core.Session{}
	session.New()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var student model.Student
	json.Unmarshal(reqBody, &student)

	student.Id = gocql.TimeUUID()
	dal.AddNewStudent(session, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	session := core.Session{}
	session.New()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var student model.Student
	json.Unmarshal(reqBody, &student)

	dal.UpdateStudent(session, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
