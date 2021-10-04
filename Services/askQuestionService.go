package Services

import (
	"time"

	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
)

type AskQuestion struct {
	AQid       int64     `gorm:"primary_key;column:AQid;AUTO_INCREMENT"`
	Sid        int64     `gorm:"column:Sid"`
	Eid        int64     `gorm:"column:Eid"`
	AQtime     time.Time `gorm:"column:AQtime"`
	AQremark   string    `gorm:"column:AQremark"`
	AQisSolved int64     `gorm:"column:AQisSolved"`
	Snickname  string
	Etitle     string
}

func (aq *AskQuestion) Insert() (AQid int64, err error) {
	var aqModel Models.AskQuestion

	aqModel.Sid = aq.Sid
	aqModel.Eid = aq.Eid
	aqModel.AQtime = aq.AQtime
	aqModel.AQremark = aq.AQremark
	aqModel.AQisSolved = aq.AQisSolved

	AQid, err = aqModel.Insert()

	return

}

func (aq *AskQuestion) Update() (result Models.AskQuestion, err error) {
	var aqModel Models.AskQuestion

	aqModel.AQid = aq.AQid
	aqModel.Sid = aq.Sid
	aqModel.Eid = aq.Eid
	aqModel.AQtime = aq.AQtime
	aqModel.AQremark = aq.AQremark
	aqModel.AQisSolved = aq.AQisSolved

	result, err = aqModel.UpdateByAQid(aqModel.AQid)

	return

}

func (aq *AskQuestion) QueryByAQid(AQid int64) (result AskQuestion, err error) {
	var aqModel Models.AskQuestion

	tmpAQ, err := aqModel.QueryByAQid(AQid)

	if err != nil {
		return
	}

	err = Utils.CopyFields(&result, tmpAQ)
	if err != nil {
		return
	}

	return
}

func (aq *AskQuestion) QueryFromTeacher(Tid int64, unresolved string, date string) (result []AskQuestion, err error) {
	var aqModel Models.AskQuestion

	var tmpAQ []Models.AskQuestion

	if unresolved == "1" {
		tmpAQ, err = aqModel.QueryFromTeacher(Tid, date)
	} else if unresolved == "0" {
		tmpAQ, err = aqModel.QueryUnresolvedFromTeacher(Tid, date)
	}

	result = make([]AskQuestion, len(tmpAQ))
	// 把 askquestion model 转为 askquestion service
	for i := 0; i < len(tmpAQ); i++ {
		err = Utils.CopyFields(&result[i], tmpAQ[i])
		if err != nil {
			return
		}
	}

	var studentModel Models.Student
	var exerciseModel Models.Exercise
	for i := 0; i < len(result); i++ {
		student, err := studentModel.QueryBySid(result[i].Sid)
		if err != nil {
			return nil, err
		}

		result[i].Snickname = student.Snickname

		if result[i].Eid != -1 {
			exercise, err := exerciseModel.QueryByEid(result[i].Eid)
			if err != nil {
				return nil, err
			}

			result[i].Etitle = exercise.Etitle
		}

	}

	if err != nil {
		return nil, err
	}

	return
}
