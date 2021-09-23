package Services

import (
	"errors"
	"time"

	"github.com/bigby/project/Models"
)

type LearnSituation struct {
	LSid         int64     `gorm:"primary_key;column:LSid;AUTO_INCREMENT"`
	Sid          int64     `gorm:"column:Sid"`
	LSattendence string    `gorm:"column:LSattendence"`
	LSperform    string    `gorm:"column:LSperform"`
	LShomework   string    `gorm:"column:LShomework"`
	LSexam       string    `gorm:"column:LSexam"`
	LSsubject    string    `gorm:"column:LSsubject"`
	LSgrade      string    `gorm:"column:LSgrade"`
	LSdate       time.Time `gorm:"column:LSdate"`
}

func (ls *LearnSituation) QueryBySid(Sid int64) (result map[string][]Models.LearnSituation, err error) {
	var lsModel Models.LearnSituation

	result, err = lsModel.QueryBySid(Sid)

	return
}

func (ls *LearnSituation) Insert() (LSid int64, err error) {
	var lsModel Models.LearnSituation

	lsModel.Sid = ls.Sid
	lsModel.LSattendence = ls.LSattendence
	lsModel.LSperform = ls.LSperform
	lsModel.LShomework = ls.LShomework
	lsModel.LSexam = ls.LSexam
	lsModel.LSsubject = ls.LSsubject
	lsModel.LSgrade = ls.LSgrade
	lsModel.LSdate = ls.LSdate

	if !(len(lsModel.LSattendence) == 0 && len(lsModel.LSperform) == 0 && len(lsModel.LShomework) == 0 && len(lsModel.LSexam) == 0 && len(lsModel.LSsubject) == 0 && len(lsModel.LSgrade) == 0) {
		LSid, err = lsModel.Insert()
	} else {
		err = errors.New("未填写信息")
	}

	return

}

func (ls *LearnSituation) Update(LSid int64) (result Models.LearnSituation, err error) {
	var lsModel Models.LearnSituation

	lsModel.LSid = ls.LSid
	lsModel.Sid = ls.Sid
	lsModel.LSattendence = ls.LSattendence
	lsModel.LSperform = ls.LSperform
	lsModel.LShomework = ls.LShomework
	lsModel.LSexam = ls.LSexam
	lsModel.LSsubject = ls.LSsubject
	lsModel.LSgrade = ls.LSgrade
	lsModel.LSdate = ls.LSdate

	result, err = lsModel.Update(lsModel.LSid)

	return
}

func (ls *LearnSituation) UpdateFromArray(lsArray []LearnSituation) (result []Models.LearnSituation, err error) {
	var lsModel Models.LearnSituation

	for i := 0; i < len(lsArray); i++ {
		// 首先查看是否需要修改
		resLS, err := lsModel.QueryByLSid(lsArray[i].LSid)
		if err != nil {
			return nil, err
		}

		var resultUpdateLS Models.LearnSituation
		if !(resLS.LSattendence == lsArray[i].LSattendence && resLS.LSperform == lsArray[i].LSperform && resLS.LShomework == lsArray[i].LShomework && resLS.LSexam == lsArray[i].LSexam && resLS.LSsubject == lsArray[i].LSsubject && resLS.LSgrade == lsArray[i].LSgrade && resLS.LSdate == lsArray[i].LSdate) {
			resultUpdateLS, err = lsArray[i].Update(lsArray[i].LSid)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, resultUpdateLS)
	}

	return
}

func (ls *LearnSituation) Delete(LSid int64) (result Models.LearnSituation, err error) {
	var LSModel Models.LearnSituation

	LSModel.LSid = ls.LSid

	result, err = LSModel.Delete(LSid)

	return
}
