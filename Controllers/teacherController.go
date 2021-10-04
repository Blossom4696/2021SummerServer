package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

// @Summary 获取一个教师信息
// @Description 获取一个教师信息
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

// @Summary 编辑教师
// @Description 编辑一个教师信息
// @Tags Admin
// @Accept json
// @Produce  json
// @Param teacherData formData Services.Teacher true "教师数据"
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

// @Summary 编辑教师
// @Description 编辑一个教师信息
// @Tags Teacher
// @Accept json
// @Produce  json
// @Param teacherData formData Services.Teacher true "教师数据"
// @Success 200 {object} Res{data=Services.Teacher}
// @Router /app/teacher/update_password [PUT]
func TeacherPasswordUpdateFromTeacher(c *gin.Context) {
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

	result, err := teacherService.UpdatePasswordFromTeacher(teacherService.Tid)
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

// @Summary 教师获取学生学情列表
// @Description 教师获取监管校区内的学生学情列表
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=map[string][]Models.LearnSituation}
// @Router /app/teacher/get_learn_situation_list [get]
func LearnSituationQueryFromTeacher(c *gin.Context) {
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

// @Summary 教师添加学生学情
// @Description 教师添加监管校区内的学生学情
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/teacher/insert_learn_situation [post]
func LearnSituationInsertFromTeacher(c *gin.Context) {
	var lsService Services.LearnSituation

	err := c.ShouldBindJSON(&lsService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := lsService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: 2,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 教师修改学生学情
// @Description 教师修改监管校区内的学生学情
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/teacher/update_learn_situation [PUT]
func LearnSituationUpdateFromTeacher(c *gin.Context) {
	var lsService Services.LearnSituation

	var lsArray []Services.LearnSituation

	err := c.ShouldBindJSON(&lsArray)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := lsService.UpdateFromArray(lsArray)
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: 2,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 教师删除学生学情
// @Description 教师删除监管校区内的学生学情
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/teacher/delete_learn_situation/:id [DELETE]
func LearnSituationDeleteFromTeacher(c *gin.Context) {
	var lsService Services.LearnSituation

	var err error

	lsService.LSid, err = strconv.ParseInt(c.Query("LSid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := lsService.Delete(lsService.LSid)
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: 2,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}

// @Summary 教师获取学生提问列表
// @Description 教师获取监管校区内的学生提问列表
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.AskQuestion}
// @Router /app/teacher/get_ask_question_list [get]
func AskQuestionListQueryFromTeacher(c *gin.Context) {
	var aqService Services.AskQuestion
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

	unresolved := c.Query("unresolved")
	date := c.Query("date")

	result, err := aqService.QueryFromTeacher(teacherService.Tid, unresolved, date)
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

// @Summary 教师获取学生提问信息
// @Description 教师获取监管校区内的学生提问信息
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=Models.AskQuestion}
// @Router /app/teacher/get_ask_question [get]
func AskQuestionQueryFromTeacher(c *gin.Context) {
	var aqService Services.AskQuestion

	var err error

	aqService.AQid, err = strconv.ParseInt(c.Query("AQid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := aqService.QueryByAQid(aqService.AQid)
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

// @Summary 教师更改学生提问信息
// @Description 教师更改监管校区内的学生提问信息
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=Models.AskQuestion}
// @Router /app/teacher/update_ask_question [post]
func AskQuestionUpdateFromTeacher(c *gin.Context) {
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

	result, err := aqService.Update()
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

// @Summary 教师整理错题信息
// @Description 教师整理错题信息并转成pdf
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/teacher/generate_wrong_problem_pdf [get]
func GenerateWPPdfFromTeacher(c *gin.Context) {
	var teacherService Services.Teacher
	var pdfService Services.Pdf

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

	recv := c.Query("Wids")

	recvStr := strings.Split(recv[1:len(recv)-1], ",")

	Wids := make([]int64, 0)

	for _, v := range recvStr {
		wid, _ := strconv.ParseInt(v, 10, 64)
		Wids = append(Wids, wid)
	}

	result, err := pdfService.WrongProblemToPDF(Wids)
	if err != nil {
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 教师整理习题信息
// @Description 教师整理习题信息并转成pdf
// @Tags Teacher
// @Accept json
// @Produce  json
// @Success 200 {object} Res
// @Router /app/teacher/generate_exercise_pdf [get]
func GenerateExercisePdfFromTeacher(c *gin.Context) {
	var teacherService Services.Teacher
	var pdfService Services.Pdf

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

	recv := c.Query("Eids")

	recvStr := strings.Split(recv[1:len(recv)-1], ",")

	Eids := make([]int64, 0)

	for _, v := range recvStr {
		eid, _ := strconv.ParseInt(v, 10, 64)
		Eids = append(Eids, eid)
	}

	var recvBool bool
	if c.Query("isPrintAnswer") == "1" {
		recvBool = true
	} else {
		recvBool = false
	}

	result, err := pdfService.ExerciseToPDF(Eids, recvBool)
	if err != nil {
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
