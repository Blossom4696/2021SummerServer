package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 管理员搜索学生信息
// @Description 管理员搜索匹配字符串显示学生列表
// @Tags Search
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Student}
// @Router /app/user/common/search_student_from_admin [get]
func SearchStudentFromAdmin(c *gin.Context) {
	var studentService Services.Student

	searchString := c.Query("word")
	result, err := studentService.Search(searchString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 教师搜索学生信息
// @Description 教师搜索匹配字符串显示学生列表
// @Tags Search
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Student}
// @Router /app/user/common/search_student_from_teacher [get]
func SearchStudentFromTeacher(c *gin.Context) {
	var teacherService Services.Teacher
	var err error

	err = json.Unmarshal([]byte(c.Request.Header["Userdata"][0]), &teacherService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	searchString := c.Query("word")

	result, err := teacherService.SearcherStudent(searchString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 搜索教师信息
// @Description 搜索匹配字符串显示教师列表
// @Tags Search
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Teacher}
// @Router /app/user/common/search_teacher [get]
func SearchTeacher(c *gin.Context) {
	var teacherService Services.Teacher

	searchString := c.Query("word")
	result, err := teacherService.Search(searchString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 搜索习题信息
// @Description 搜索匹配字符串显示习题列表
// @Tags Search
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Exercise}
// @Router /app/user/common/search_exercise [get]
func SearchExercise(c *gin.Context) {
	var exerciseService Services.Exercise

	searchString := c.Query("word")
	result, err := exerciseService.Search(searchString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 搜索错题信息
// @Description 搜索匹配字符串显示错题列表
// @Tags Search
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.WrongProblem}
// @Router /app/user/common/search_wrong_problem [get]
func SearchWrongProblem(c *gin.Context) {
	var wpService Services.WrongProblem

	var err error
	wpService.Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	searchString := c.Query("word")
	result, err := wpService.Search(wpService.Sid, searchString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}
