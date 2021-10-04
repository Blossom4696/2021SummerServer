package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 今日习题列表
// @Description 获取今日习题列表
// @Tags TodayExercise
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Success 200 {object} Res{data=[]Models.TodayExercise}
// @Router /app/today/get_today_exercise [get]
func TodayExerciseQuery(c *gin.Context) {
	var teService Services.TodayExercise

	teService.Sid, _ = strconv.ParseInt(c.Query("Sid"), 10, 64)
	teService.TEdate, _ = time.Parse("2006/1/2", time.Now().Format("2006/1/2"))

	result, err := teService.Query()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusOK, Res{
			Code: 2,
			Msg:  "No exercise!",
			Data: result,
		})
	} else {
		c.JSON(http.StatusOK, Res{
			Code: 1,
			Msg:  "Query Success!",
			Data: result,
		})
	}

}

// @Summary 判题(学生)
// @Description 学生提交今日习题，系统进行保存并自动判题
// @Tags TodayExercise
// @Accept json
// @Produce  json
// @Param Sid query string true "学生id"
// @Param Answer query string true "题目id和学生答案"
// @Success 200 {object} Res{data=[]Models.TodayExercise}
// @Router /app/today/judge_today_exercise [get]
func JudgeTodayExercise(c *gin.Context) {
	var teService Services.TodayExercise

	var answer map[string]string
	teService.Sid, _ = strconv.ParseInt(c.Query("Sid"), 10, 64)
	json.Unmarshal([]byte(c.Query("SendAnswer")), &answer)
	teService.TEdate, _ = time.Parse("2006/1/2", time.Now().Format("2006/1/2"))
	result, err := teService.Judge(answer)

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
		Msg:  "Judge Success!",
		Data: result,
	})
}
