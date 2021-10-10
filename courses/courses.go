package course

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/williamluisan/golang_rest_sample/data"
)

func GetCourseList(c *gin.Context) {
	c.JSON(http.StatusOK, data.Courses[0])
}

func InsertCourse(c *gin.Context) {
	var course data.Course

	// should validation here
	// ...

	if c.ShouldBind(&course) != nil {
		return
	}
	
	data.Courses = append(data.Courses, course)
	c.JSON(http.StatusCreated, "Successfully insert a course")
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	course_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// should validation here
	// ...

	for i, a := range data.Courses {
		if course_id == a.ID {
			data.Courses[i].Name = c.PostForm("name")
			data.Courses[i].Category = c.PostForm("category")
		}
	}

	c.JSON(http.StatusNoContent, nil)
}