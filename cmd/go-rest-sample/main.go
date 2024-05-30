package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kekaswork/go-rest-sample/internal/database"
	"github.com/kekaswork/go-rest-sample/internal/service/mark"
	"github.com/kekaswork/go-rest-sample/internal/service/student"
	"github.com/kekaswork/go-rest-sample/internal/service/subject"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db := database.NewService()
	defer db.Close()

	router := gin.Default()

	// Students
	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudent)
	router.PUT("/students/:id", updateStudent)
	router.POST("/students", createStudent)
	router.DELETE("/students/:id", deleteStudent)

	// Subjects
	router.GET("/subjects", getSubjects)
	router.GET("/subjects/:id", getSubject)
	router.PUT("/subjects/:id", updateSubject)
	router.POST("/subjects", createSubject)
	router.DELETE("/subjects/:id", deleteSubject)

	// Marks
	router.GET("/marks", getMarks)
	router.GET("/marks/:id", getMark)
	router.PUT("/marks/:id", updateMark)
	router.POST("/marks", createMark)
	router.DELETE("/marks/:id", deleteMark)

	// router.GET("/report", generateReport)

	router.Run(":3000")
}

func getStudents(c *gin.Context) {
	studentsService := student.NewService()
	students, err := studentsService.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": students}
	c.JSON(http.StatusOK, response)
}

func getStudent(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	studentsService := student.NewService()
	student, err := studentsService.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": student}
	c.JSON(http.StatusOK, response)
}

func createStudent(c *gin.Context) {
	var req student.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := student.NewService()
	student, err := service.Add(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": student}
	c.JSON(http.StatusOK, response)
}

func updateStudent(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req student.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := student.NewService()
	student, err := service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": student}
	c.JSON(http.StatusOK, response)
}

func deleteStudent(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	studentsService := student.NewService()
	student, err := studentsService.Remove(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": student}
	c.JSON(http.StatusOK, response)
}

func getSubjects(c *gin.Context) {
	service := subject.NewService()
	subjects, err := service.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": subjects}
	c.JSON(http.StatusOK, response)
}

func getSubject(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	service := subject.NewService()
	subject, err := service.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": subject}
	c.JSON(http.StatusOK, response)
}

func createSubject(c *gin.Context) {
	var req subject.CreateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := subject.NewService()
	subject, err := service.Add(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": subject}
	c.JSON(http.StatusOK, response)
}

func updateSubject(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	var req subject.CreateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := subject.NewService()
	subject, err := service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{"code": 200, "data": subject}
	c.JSON(http.StatusOK, response)
}

func deleteSubject(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	service := subject.NewService()
	subject, err := service.Remove(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{"code": 200, "data": subject}
	c.JSON(http.StatusOK, response)
}

func getMarks(c *gin.Context) {
	service := mark.NewService()
	marks, err := service.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": marks}
	c.JSON(http.StatusOK, response)
}

func getMark(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mark ID"})
		return
	}

	service := mark.NewService()
	mark, err := service.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": mark}
	c.JSON(http.StatusOK, response)
}

func createMark(c *gin.Context) {
	var req mark.CreateMarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := mark.NewService()
	mark, err := service.Add(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response := gin.H{"code": 200, "data": mark}
	c.JSON(http.StatusOK, response)
}

func deleteMark(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mark ID"})
		return
	}

	service := mark.NewService()
	mark, err := service.Remove(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{"code": 200, "data": mark}
	c.JSON(http.StatusOK, response)
}

func updateMark(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mark ID"})
		return
	}

	var req mark.CreateMarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := mark.NewService()
	mark, err := service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{"code": 200, "data": mark}
	c.JSON(http.StatusOK, response)
}
