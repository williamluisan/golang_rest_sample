package teacher

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	
	"github.com/williamluisan/golang_rest_sample/data"
)

func GetTeacherList(c *gin.Context) { 
	c.JSON(http.StatusOK, data.Teachers)
}

func InsertTeacher(c *gin.Context) {
	var teacher data.Teacher

	// should validation here
	// ...

	if c.ShouldBind(&teacher) != nil {
		return
	}

	data.Teachers = append(data.Teachers, teacher)
	c.JSON(http.StatusCreated, "Successfully insert a teacher")
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	teacher_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// should validation here
	// ...

	for i, a := range data.Teachers {
		if teacher_id == a.ID {
			data.Teachers[i].Name = c.PostForm("name")
			data.Teachers[i].Gender = c.PostForm("category")
			data.Teachers[i].Speciality = c.PostForm("speciality")
		}
	}

	c.JSON(http.StatusNoContent, nil)
}