package main

import (
	"github.com/gin-gonic/gin"

	course "github.com/williamluisan/golang_rest_sample/courses"
	student "github.com/williamluisan/golang_rest_sample/students"
	teacher "github.com/williamluisan/golang_rest_sample/teachers"
)

func main() {
	router := gin.Default()

	router.GET("/course", course.GetCourseList)
	router.POST("/course", course.InsertCourse)
	router.PUT("/course/:id", course.UpdateCourse)

	router.GET("/teacher", teacher.GetTeacherList)
	router.POST("/teacher", teacher.InsertTeacher)
	router.PUT("/teacher/:id", teacher.UpdateTeacher)
	
	router.GET("/student", student.GetStudentList)
	router.POST("/student", student.InsertStudent)
	router.PUT("/student/:id", student.UpdateStudent)

	router.Run("localhost:8080")
}