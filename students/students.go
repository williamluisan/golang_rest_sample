package student

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/williamluisan/golang_rest_sample/data"
)

func GetStudentList(c *gin.Context) {
	c.JSON(http.StatusOK, data.Students)
}

func InsertStudent(c *gin.Context) {
	var student data.Student

	// should validation here
	// ...

	if c.ShouldBind(&student) != nil {
		return
	}

	data.Students = append(data.Students, student)
	c.JSON(http.StatusCreated, "Successfully insert a student")
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	student_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// should validation here
	// ...

	for i, a := range data.Students {
		if student_id == a.ID {
			data.Students[i].Name = c.PostForm("name")
			age, err := strconv.Atoi(c.PostForm("age"))
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			data.Students[i].Age = age
			data.Students[i].Gender = c.PostForm("gender")
		}
	}

	c.JSON(http.StatusNoContent, nil)
}