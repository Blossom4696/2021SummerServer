package Controllers

import (
	"net/http"
	"strconv"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 插入错题
// @Description 插入错题数据
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param WrongProblemData body Models.WrongProblem true "习题数据"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":500001}
// @Router /app/wrong_problem/insert_wrong_problem [post]
func WrongProblemInsert(c *gin.Context) {
	var wpService Services.WrongProblem

	err := c.ShouldBindJSON(&wpService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	id, err := wpService.Insert()
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

// @Summary 删除错题
// @Description 删除错题数据
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param Wid query string true "错题id"
// @Success 200 {object} Res {"code":200,"msg":"Delete() success!","data":500001}
// @Router /app/wrong_problem/delete_wrong_problem/:id [delete]
func WrongProblemDelete(c *gin.Context) {
	var wpService Services.WrongProblem

	var err error

	wpService.Wid, err = strconv.ParseInt(c.Query("Wid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := wpService.Delete(wpService.Wid)
	if err != nil || result.Wid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Delete Success!",
		Data: result.Wid,
	})
}

// @Summary 错题列表
// @Description 获取错题列表
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res{data=[]Models.WrongProblem}
// @Router /app/wrongproblem/get_wrongproblem [get]
func WrongProblemQuery(c *gin.Context) {
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

	result, err := wpService.QueryBySid()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
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

// @Summary 错题列表筛选
// @Description 获取筛选后的错题列表
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Param subject query string true "科目"
// @Param type query string true "类型"
// @Param unit query string true "单元"
// @Param difficulty query string true "难度"
// @Param grade query string true "年级"
// @Success 200 {object} Res{data=[]Models.WrongProblem}
// @Router /app/exercise/get_sorted_exercise [get]
func WrongProblemQueryBySorted(c *gin.Context) {
	var wpService Services.WrongProblem

	var err error
	wpService.Sid, err = strconv.ParseInt(c.Query("Sid"), 10, 64)
	wpService.Wsubject = c.Query("subject")
	wpService.Wtype = c.Query("type")
	wpService.Wunit = c.Query("unit")
	wpService.Wdifficulty = c.Query("difficulty")
	wpService.Wgrade = c.Query("grade")

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := wpService.QueryBySort()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
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

// @Summary 查找错题
// @Description 同过Wid查找错题
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param Wid query string true "错题id"
// @Success 200 {object} Res{data=Models.WrongProblem}
// @Router /app/wrong_problem/get_wrong_problem [get]
func WrongProblemQueryByWid(c *gin.Context) {
	var wpService Services.WrongProblem

	var err error
	wpService.Wid, err = strconv.ParseInt(c.Query("Wid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := wpService.QueryByWid()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
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

// @Summary 更改错题
// @Description 更新错题数据
// @Tags WrongProblem
// @Accept json
// @Produce  json
// @Param WPData body Models.WrongProblem true "习题数据"
// @Success 200 {object} Res {"code":200,"msg":"Update() success!","data":500001}
// @Router /app/wrong_problem/update_wrong_problem/:id [put]
func WrongProblemUpdate(c *gin.Context) {
	var wpService Services.WrongProblem

	err := c.ShouldBindJSON(&wpService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := wpService.Update(wpService.Wid)
	if err != nil || result.Wid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Update() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: result.Wid,
	})
}
