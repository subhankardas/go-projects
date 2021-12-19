package dal

import (
	"go-projects/microservices/students-service/core"
	"go-projects/microservices/students-service/model"
	"log"

	"github.com/gocql/gocql"
)

func AddNewStudent(session core.Session, student model.Student) {
	stmt := "INSERT INTO students(id, first_name, last_name, age, marks) VALUES(?, ?, ?, ?, ?)"
	query := session.Get().Query(stmt, student.Id, student.FirstName, student.LastName, student.Age, student.Marks)
	if err := query.Exec(); err != nil {
		log.Printf("Error while adding new student %v\n", student)
		log.Println(err)
		return
	}
	log.Printf("Added new student %v\n", student)
}

func GetAllStudents(session core.Session) []model.Student {
	log.Println("Getting all students")
	var studs []model.Student
	m := map[string]interface{}{}

	iter := session.Get().Query("SELECT * FROM students").Iter()
	for iter.MapScan(m) {
		studs = append(studs, model.Student{
			Id:        m["id"].(gocql.UUID),
			FirstName: m["first_name"].(string),
			LastName:  m["last_name"].(string),
			Age:       m["age"].(int),
			Marks:     m["marks"].(int),
		})
		m = map[string]interface{}{}
	}
	return studs
}

func UpdateStudent(session core.Session, student model.Student) {
	if err := session.Get().Query("UPDATE students SET first_name = ?, last_name = ?, age = ? WHERE id = ?",
		student.FirstName, student.LastName, student.Age, student.Id).Exec(); err != nil {
		log.Printf("Error while updating student %v\n", student)
		log.Println(err)
		return
	}
	log.Printf("Updated existing student %v\n", student)
}

func DeleteStudent(session core.Session, id string) {
	if err := session.Get().Query("DELETE FROM students WHERE id = ?", id).Exec(); err != nil {
		log.Printf("Error while deleting student with ID %v\n", id)
		log.Println(err)
		return
	}
	log.Printf("Deleted existing student with ID %v\n", id)
}
