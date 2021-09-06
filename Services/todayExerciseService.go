package Services

import (
	"strconv"
	"time"

	"github.com/bigby/project/Models"
)

type TodayExercise struct {
	TEid         int64     `gorm:"primary_key;column:TEid;AUTO_INCREMENT"`
	Eid          int64     `gorm:"column:Eid"`
	Sid          int64     `gorm:"column:Sid"`
	TEdate       time.Time `gorm:"column:TEdate"`
	TEanswerTxt  string    `gorm:"column:TEanswerTxt"`
	TEanswerPath string    `gorm:"column:TEanswerPath"`
}

type TEInsertRecv struct {
	Eids []int64 `json:"Eids"`
	Sids []int64 `json:"Sids"`
	Date string  `json:"Date"`
}

func (te *TodayExercise) Query() (result []map[string]interface{}, err error) {
	var teModel Models.TodayExercise

	teArray, err := teModel.QueryBySidAndDate(te.Sid, te.TEdate)

	var exerciseModel Models.Exercise

	var exercise Models.Exercise

	for i := 0; i < len(teArray); i++ {
		resultMap := make(map[string]interface{})
		exercise, err = exerciseModel.QueryByEId(teArray[i].Eid)

		resultMap["Exercise"] = exercise

		// 系统自动判题
		if len(exercise.EanswerTxt) == 0 {
			resultMap["Judge"] = "无标准答案"
		} else if exercise.EanswerTxt == teArray[i].TEanswerTxt {
			resultMap["Judge"] = "正确"
		} else {
			resultMap["Judge"] = "错误"
		}

		resultMap["Answer"] = teArray[i].TEanswerTxt
		result = append(result, resultMap)
	}

	return
}

func (te *TodayExercise) Judge(answerMap map[string]string) (result []map[string]string, err error) {

	var exerciseModel Models.Exercise
	result = make([]map[string]string, 0)

	var teModel Models.TodayExercise

	for id, answer := range answerMap {
		Eid, _ := strconv.ParseInt(id, 10, 64)
		exercise, _ := exerciseModel.QueryByEId(Eid)

		resultMap := make(map[string]string)
		resultMap["answerTxt"] = exercise.EanswerTxt
		resultMap["answerPath"] = exercise.EanswerTxt

		// 系统自动判题
		if len(exercise.EanswerTxt) == 0 {
			resultMap["judge"] = "无标准答案"
		} else if exercise.EanswerTxt == answer {
			resultMap["judge"] = "正确"
		} else {
			resultMap["judge"] = "错误"
		}
		result = append(result, resultMap)

		// 把学生答案上传数据库
		if answer == "" {
			answer = "学生未作答"
		}
		teModel.TEanswerTxt = answer //先上传文字版
		_, err := teModel.UpdateBySidAndEid(te.Sid, Eid, te.TEdate)

		if err != nil {
			return result, err
		}
	}

	return
}

func (te *TodayExercise) InsertFromTeacher(Sids []int64, Eids []int64) (result []int64, err error) {
	var teModel Models.TodayExercise

	teModel.TEdate = te.TEdate
	teModel.TEanswerTxt = ""
	teModel.TEanswerPath = ""

	var TEid int64
	for i := 0; i < len(Sids); i++ {
		teModel.Sid = Sids[i]

		for j := 0; j < len(Eids); j++ {
			teModel.Eid = Eids[j]

			TEid, err = teModel.Insert()
			if err != nil {
				continue
			}
			result = append(result, TEid)
		}
	}
	return
}
