package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

var students = []Student{
	{ID: 1, Name: "Hilmi", Class: "12-B", Teacher: "Kemal"},
	{ID: 2, Name: "Burak", Class: "8-C", Teacher: "Mehmet"},
}

func getStudentsAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, students)
}

func getStudent(intId int) (*Student, error) {
	for i, s := range students {
		if s.ID == intId {
			return &students[i], nil
		}
	}

	return nil, errors.New("student not found")
}

func getStudentsByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	student, err := getStudent(intId)
	if err == nil {
		ctx.IndentedJSON(http.StatusOK, student)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found!"})
	}
}

func createStudent(ctx *gin.Context) {
	var studentByUser Student
	err := ctx.BindJSON(&studentByUser)

	if err == nil && studentByUser.ID != 0 && studentByUser.Name != "" && studentByUser.Class != "" && studentByUser.Teacher != "" {
		students = append(students, studentByUser)
		ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Student created successfully.", "student_id": studentByUser.ID})
		return
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Student could not created!"})
		return
	}
}

func main() {
	router := gin.Default()

	router.GET("/students", getStudentsAll)
	router.GET("/students/:id", getStudentsByID)

	router.POST("/students", createStudent)

	router.Run("localhost:8080")
}
