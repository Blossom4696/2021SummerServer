package Controllers

import (
	"net/http"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 学生登录
// @Description 匹配学生密码，若登录成功则返回学生信息
// @Tags Student
// @Accept json
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} Res{data=Models.Student}
// @Router /app/login/student_login [get]
func StudentLogin(c *gin.Context) {
	var studentService Services.Student

	studentService.Sname = c.Query("username")
	studentService.Spassword = c.Query("password")

	res, err := studentService.Login()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "No Data",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 学生注册
// @Description 根据发送的学生信息注册新用户
// @Tags Student
// @Accept json
// @Produce  json
// @Param StudentData body Models.Student true "学生信息"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":100001}
// @Router /app/student/register [post]
func StudentRegister(c *gin.Context) {
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

	res, err := studentService.Register()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Register Failed!: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 教师登录
// @Description 匹配教师密码，若登录成功则返回教师信息
// @Tags Teacher
// @Accept json
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} Res{data=Models.Teacher}
// @Router /app/login/teacher_login [get]
func TeacherLogin(c *gin.Context) {
	var teacherService Services.Teacher

	teacherService.Tname = c.Query("username")
	teacherService.Tpassword = c.Query("password")

	res, err := teacherService.Login()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "No Data",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 管理员登录
// @Description 匹配管理员密码，若登录成功则返回管理员信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} Res{data=Models.Admin}
// @Router /app/login/admin_login [get]
func AdminLogin(c *gin.Context) {
	var adminService Services.Admin

	adminService.Aname = c.Query("username")
	adminService.Apassword = c.Query("password")

	res, err := adminService.Login()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "No Data",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 登录测试
// @Description 登录test
// @Tags Test
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/test/login [get]
func LoginTest(c *gin.Context) {
	var teacherService Services.Teacher

	res, err := teacherService.LoginTest()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "No Data",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Success",
		Data: res,
	})
}
