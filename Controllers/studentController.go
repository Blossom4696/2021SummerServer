package Controllers

import (
	"net/http"
	"strconv"

	"github.com/bigby/project/Services"
	"github.com/bigby/project/Utils"
	"github.com/gin-gonic/gin"
)

// @Summary 获取自己信息
// @Description 获取自己的用户信息
// @Tags Student
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res{data=Models.Student}
// @Router /app/student/get_student_self [get]
func StudentQueryFromStudent(c *gin.Context) {
	var studentService Services.Student

	var err error

	studentService.Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := studentService.QueryBySid(studentService.Sid)
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

// @Summary 编辑学生
// @Description 编辑一个学生信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param studentData formData Services.Student true
// @Success 200 {object} Res{data=Services.Student}
// @Router /app/student/update_student [POST]
func StudentUpdateFromStudent(c *gin.Context) {
	var studentService Services.Student

	err := c.ShouldBindJSON(&studentService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	studentService.Sgrade = Utils.GradeStringToInt(studentService.SgradeName)
	if studentService.Sicon == "" {
		studentService.Sicon = "default_boy.png"
	}

	result, err := studentService.Update(studentService.Sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
