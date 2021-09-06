package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 教师获取学生信息列表
// @Description 教师获取监管校区内的学生信息列表
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Student}
// @Router /app/teacher/get_student_list [get]
func StudentListQueryFromTeacher(c *gin.Context) {
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

	result, err := teacherService.QueryStudentList()
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

// @Summary 教师获取学生信息列表（排序后）
// @Description 教师获取监管校区内的学生信息列表（排序后）
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Student}
// @Router /app/teacher/get_student_sorted_list [get]
func StudentListQuerySortedFromTeacher(c *gin.Context) {
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
	sortName := c.Query("sortName")
	sortDir := c.Query("sortDir")

	result, err := teacherService.QueryStudentSortedList(searchString, sortName, sortDir)
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

// @Summary 教师获取一个学生信息
// @Description 教师获取监管校区内的一个学生信息
// @Tags Teacher
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res{data=Models.Student}
// @Router /app/teacher/get_student [get]
func StudentQueryFromTeacher(c *gin.Context) {
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

// @Summary 插入今日习题
// @Description 插入今日习题数据
// @Tags Teacher
// @Accept json
// @Produce  json
// @Param ExerciseData body Models.Exercise true "习题数据"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":500001}
// @Router /app/exercise/insert_exercise [post]
func TodayExerciseInsertFromTeacher(c *gin.Context) {
	var teService Services.TodayExercise

	var err error
	var recv Services.TEInsertRecv
	c.ShouldBindJSON(&recv)
	teService.TEdate, _ = time.Parse("2006/1/2", recv.Date)

	id, err := teService.InsertFromTeacher(recv.Sids, recv.Eids)
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

// @Summary 管理员获取一个教师信息
// @Description 管理员获取一个教师信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param Tid query string true "教师id"
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/teacher/get_teacher [get]
func TeacherQueryFromTeacher(c *gin.Context) {
	var teacherService Services.Teacher

	var err error
	teacherService.Tid, err = strconv.ParseInt(c.Query("Tid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := teacherService.QueryByTid(teacherService.Tid)
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

// @Summary 管理员编辑教师
// @Description 管理员编辑一个教师信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param teacherData formData Services.Teacher true
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/teacher/update_teacher [POST]
func TeacherUpdateFromTeacher(c *gin.Context) {
	var teacherService Services.Teacher

	err := c.ShouldBindJSON(&teacherService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	if teacherService.Ticon == "" {
		teacherService.Ticon = "default_teacher.png"
	}

	result, err := teacherService.UpdateFromTeacher(teacherService.Tid)
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
