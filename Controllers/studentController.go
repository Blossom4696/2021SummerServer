package Controllers

import (
	"net/http"
	"strconv"
	"time"

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

// @Summary 编辑学生密码
// @Description 编辑一个学生密码信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param studentData formData Services.Student true
// @Success 200 {object} Res{data=Services.Student}
// @Router /app/student/update_password [PUT]
func StudentPasswordUpdateFromStudent(c *gin.Context) {
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

	result, err := studentService.UpdatePasswordFromStudent(studentService.Sid)
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

// @Summary 学生插入预约教师
// @Description 插入预约教师数据
// @Tags Student
// @Accept json
// @Produce  json
// @Param OTData body Models.OrderTeacher true "预约教师数据"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":500001}
// @Router /app/student/insert_order_teacher [post]
func OrderTeacherInsertFromStudent(c *gin.Context) {
	var otService Services.OrderTeacher

	err := c.ShouldBindJSON(&otService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	otService.OTcreateTime = time.Now()

	id, err := otService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Insert Failed!: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Insert Success!",
		Data: id,
	})
}

// @Summary 学生获取学情列表
// @Description 学生获取监管校区内的学生学情列表
// @Tags Student
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=map[string][]Models.LearnSituation}
// @Router /app/student/get_learn_situation_list [get]
func LearnSituationQueryFromStudent(c *gin.Context) {
	var lsService Services.LearnSituation

	var err error

	var Sid int64
	Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := lsService.QueryBySid(Sid)
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

// @Summary 学生提问
// @Description 插入提问问题数据
// @Tags Student
// @Accept json
// @Produce  json
// @Param OTData body Models.AskQuestion true "预约教师数据"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":500001}
// @Router /app/student/insert_ask_question [post]
func AskQuestionInsertFromStudent(c *gin.Context) {
	var aqService Services.AskQuestion

	err := c.ShouldBindJSON(&aqService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	aqService.AQtime = time.Now()

	id, err := aqService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Insert Failed!: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Insert Success!",
		Data: id,
	})
}
