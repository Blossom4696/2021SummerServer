package Services

import (
	"time"

	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
)

type WrongProblem struct {
	Wid           int64     `gorm:"primary_key;column:Wid;AUTO_INCREMENT"`
	Wtitle        string    `gorm:"column:Wtitle"`
	Wdescription  string    `gorm:"column:Wdescription"`
	WproblemTxt   string    `gorm:"column:WproblemTxt"`
	WproblemPath  string    `gorm:"column:WproblemPath"`
	WanswerTxt    string    `gorm:"column:WanswerTxt"`
	WanswerPath   string    `gorm:"column:WanswerPath"`
	Wgrade        string    `gorm:"column:Wgrade"`
	Wsubject      string    `gorm:"column:Wsubject"`
	Wtype         string    `gorm:"column:Wtype"`
	Wdifficulty   string    `gorm:"type:enum('简单','中等','困难','竞赛');column:Wdifficulty;default:简单"`
	Wunit         string    `gorm:"column:Wunit"`
	Sid           int64     `gorm:"column:Sid"`
	WmyAnswerTxt  string    `gorm:"column:WmyAnswerTxt"`
	WmyAnswerPath string    `gorm:"column:WmyAnswerPath"`
	WcreateTime   time.Time `gorm:"column:WcreateTime"`
}

func (wp *WrongProblem) Insert() (Wid int64, err error) {
	var wpModel Models.WrongProblem

	// wpModel.Wid = exercise.Wid
	wpModel.Wtitle = wp.Wtitle
	wpModel.Wdescription = wp.Wdescription
	wpModel.WproblemTxt = wp.WproblemTxt
	wpModel.WproblemPath = wp.WproblemPath
	wpModel.WanswerTxt = wp.WanswerTxt
	wpModel.WanswerPath = wp.WanswerPath
	wpModel.WmyAnswerTxt = wp.WmyAnswerTxt
	wpModel.WmyAnswerPath = wp.WmyAnswerPath
	wpModel.Wgrade = wp.Wgrade
	wpModel.Wsubject = wp.Wsubject
	wpModel.Wtype = wp.Wtype
	wpModel.Wdifficulty = wp.Wdifficulty
	wpModel.Wunit = wp.Wunit
	wpModel.Sid = wp.Sid
	wpModel.WcreateTime = time.Now()

	Wid, err = wpModel.Insert()

	return

}

func (wp *WrongProblem) Delete(Wid int64) (result Models.WrongProblem, err error) {
	var wpModel Models.WrongProblem

	wpModel.Wid = wp.Wid

	result, err = wpModel.Delete(wpModel.Wid)

	return
}

func (wp *WrongProblem) QueryBySid() (result []Models.WrongProblem, err error) {
	var wpModel Models.WrongProblem

	result, err = wpModel.QueryBySid(wp.Sid)

	return
}

func (wp *WrongProblem) QueryByWid() (result Models.WrongProblem, err error) {
	var wpModel Models.WrongProblem

	result, err = wpModel.QueryByWid(wp.Wid)

	return
}

func (wp *WrongProblem) QueryBySort() (result []Models.WrongProblem, err error) {
	var wpModel Models.WrongProblem

	wpModel.Wsubject = wp.Wsubject
	wpModel.Wgrade = wp.Wgrade
	wpModel.Wdifficulty = wp.Wdifficulty
	wpModel.Wtype = wp.Wtype
	wpModel.Wunit = wp.Wunit

	result, err = wpModel.QueryBySortedAndSid(wp.Sid)

	return
}

func (wp *WrongProblem) Search(Sid int64, searchString string) (result []WrongProblem, err error) {
	var wpModel Models.WrongProblem

	// 匹配title
	tmpTitle, err := wpModel.Search(Sid, searchString, "Wtitle")
	if err != nil {
		return
	}
	// 匹配subject
	tmpSubject, err := wpModel.Search(Sid, searchString, "Wsubject")
	if err != nil {
		return
	}
	// 匹配type
	tmpType, err := wpModel.Search(Sid, searchString, "Wtype")
	if err != nil {
		return
	}
	// 匹配difficulty
	tmpDifficulty, err := wpModel.Search(Sid, searchString, "Wdifficulty")
	if err != nil {
		return
	}
	// 匹配unit
	tmpUnit, err := wpModel.Search(Sid, searchString, "Wunit")
	if err != nil {
		return
	}
	// 匹配grade
	tmpGrade, err := wpModel.Search(Sid, searchString, "Wgrade")
	if err != nil {
		return
	}

	tmpWP := append(tmpTitle, tmpSubject...)
	tmpWP = append(tmpWP, tmpType...)
	tmpWP = append(tmpWP, tmpDifficulty...)
	tmpWP = append(tmpWP, tmpUnit...)
	tmpWP = append(tmpWP, tmpGrade...)
	tmpWP = DeleteDuplicateValueWP(tmpWP)

	result = make([]WrongProblem, len(tmpWP))
	// 把teacher model转为teacher service
	for i := 0; i < len(tmpWP); i++ {
		err = Utils.CopyFields(&result[i], tmpWP[i])
		if err != nil {
			return
		}
	}

	return
}

// 去重切片
func DeleteDuplicateValueWP(s []Models.WrongProblem) (ret []Models.WrongProblem) {
	tmpM := make(map[Models.WrongProblem]int) // key的类型要和切片中的数据类型一致
	for _, v := range s {
		tmpM[v] = 1
	}
	// 先清空s
	s = []Models.WrongProblem{}
	for i := range tmpM {
		s = append(s, i)
	}
	return s
}
