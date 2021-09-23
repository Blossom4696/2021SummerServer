package Controllers

import (
	"net/http"
	"strconv"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data"`
}

// @Summary 习题列表
// @Description 获取习题列表
// @Tags Exercise
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Exercise}
// @Router /app/exercise/get_all_exercise [get]
func ExerciseQuery(c *gin.Context) {
	var exerciseService Services.Exercise

	result, err := exerciseService.Query()

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

// @Summary 插入习题
// @Description 插入习题数据
// @Tags Exercise
// @Accept json
// @Produce  json
// @Param ExerciseData body Models.Exercise true "习题数据"
// @Success 200 {object} Res {"code":200,"msg":"Insert() success!","data":500001}
// @Router /app/exercise/insert_exercise [post]
func ExerciseInsert(c *gin.Context) {
	var exerciseService Services.Exercise

	err := c.ShouldBindJSON(&exerciseService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	id, err := exerciseService.Insert()
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

// @Summary 更改习题
// @Description 更新习题数据
// @Tags Exercise
// @Accept json
// @Produce  json
// @Param EnameTxt body Models.Exercise true "习题数据"
// @Success 200 {object} Res {"code":200,"msg":"Update() success!","data":500001}
// @Router /app/exercise/update_exercise/:id [put]
func ExerciseUpdate(c *gin.Context) {
	var exerciseService Services.Exercise

	err := c.ShouldBindJSON(&exerciseService)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := exerciseService.Update(exerciseService.Eid)
	if err != nil || result.Eid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Update() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: result.Eid,
	})
}

// @Summary 删除习题
// @Description 删除习题数据
// @Tags Exercise
// @Accept json
// @Produce  json
// @Param Eid query string true "习题id"
// @Success 200 {object} Res {"code":200,"msg":"Delete() success!","data":500001}
// @Router /app/exercise/delete_exercise/:id [delete]
func ExerciseDelete(c *gin.Context) {
	var exerciseService Services.Exercise

	var err error

	exerciseService.Eid, err = strconv.ParseInt(c.Query("Eid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := exerciseService.Delete(exerciseService.Eid)
	if err != nil || result.Eid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Delete Success!",
		Data: result.Eid,
	})
}

// @Summary 查找习题
// @Description 通过Eid查找习题
// @Tags Exercise
// @Accept json
// @Produce  json
// @Param Eid query string true "习题id"
// @Success 200 {object} Res {"code":200,"msg":"Query() success!","data":500001}
// @Router /app/exercise/get_exercise [get]
func ExerciseQueryByEid(c *gin.Context) {
	var exerciseService Services.Exercise

	var err error
	exerciseService.Eid, err = strconv.ParseInt(c.Query("Eid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := exerciseService.QueryByEid()
	if err != nil || result.Eid == 0 {
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

// @Summary 获取习题属性值
// @Description 获取习题列表年级、科目、类型、难度、单元的全部可能的值
// @Tags Exercise
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Services.ExerciseField}
// @Router /app/exercise/get_all_exercise_field [get]
func ExerciseGetField(c *gin.Context) {
	var exerciseService Services.Exercise

	result, err := exerciseService.QueryAllField()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Query() error!",
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Field Success!",
		Data: result,
	})
}

// @Summary 习题列表筛选
// @Description 获取筛选后的习题列表
// @Tags Exercise
// @Accept json
// @Produce  json
// @Param subject query string true "科目"
// @Param type query string true "类型"
// @Param unit query string true "单元"
// @Param difficulty query string true "难度"
// @Param grade query string true "年级"
// @Success 200 {object} Res{data=[]Models.Exercise}
// @Router /app/exercise/get_sorted_exercise [get]
func ExerciseQueryBySorted(c *gin.Context) {
	var exerciseService Services.Exercise

	var err error
	exerciseService.Esubject = c.Query("subject")
	exerciseService.Etype = c.Query("type")
	exerciseService.Eunit = c.Query("unit")
	exerciseService.Edifficulty = c.Query("difficulty")
	exerciseService.Egrade = c.Query("grade")

	if err != nil {
		c.JSON(http.StatusBadRequest, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	result, err := exerciseService.QueryBySort()

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
