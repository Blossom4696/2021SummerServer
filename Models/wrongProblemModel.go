package Models

import (
	"time"

	orm "github.com/bigby/project/Database"
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

//增
func (wrongProblem WrongProblem) Insert() (Wid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("WRONGPROBLEM").Create(&wrongProblem)
	Wid = wrongProblem.Wid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (wrongProblem *WrongProblem) Delete(Wid int64) (Result WrongProblem, err error) {

	if err = orm.Eloquent.Table("WRONGPROBLEM").Select([]string{"Wid"}).First(&wrongProblem, Wid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("WRONGPROBLEM").Delete(&wrongProblem).Error; err != nil {
		return
	}
	Result = *wrongProblem
	return
}

//改
func (wrongProblem *WrongProblem) Update(Wid int64) (updateWrongProblem WrongProblem, err error) {

	if err = orm.Eloquent.Table("WRONGPROBLEM").Select([]string{"Wid"}).First(&updateWrongProblem, Wid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("WRONGPROBLEM").Model(&updateWrongProblem).Updates(&wrongProblem).Error; err != nil {
		return
	}
	return
}

//按照Sid查
func (wrongProblem *WrongProblem) QueryBySid(Sid int64) (wrongProblems []WrongProblem, err error) {

	if err = orm.Eloquent.Table("WRONGPROBLEM").Where("Sid = ?", Sid).Find(&wrongProblems).Error; err != nil {
		return
	}
	return
}

//按Wid查
func (wrongProblem *WrongProblem) QueryByWid(Wid int64) (result WrongProblem, err error) {

	if err = orm.Eloquent.Table("WRONGPROBLEM").Where("Wid = ?", Wid).First(&result).Error; err != nil {
		return
	}
	return
}

// 筛选，按照Sid
func (wrongProblem *WrongProblem) QueryBySortedAndSid(Sid int64) (result []WrongProblem, err error) {
	sql := orm.Eloquent.Table("WRONGPROBLEM")
	sql = sql.Where("Sid = ?", Sid)
	if wrongProblem.Wsubject != "不限" {
		sql = sql.Where("Wsubject = ?", wrongProblem.Wsubject)
	}
	if wrongProblem.Wtype != "不限" {
		sql = sql.Where("Wtype = ?", wrongProblem.Wtype)
	}
	if wrongProblem.Wgrade != "不限" {
		sql = sql.Where("Wgrade = ?", wrongProblem.Wgrade)
	}
	if wrongProblem.Wdifficulty != "不限" {
		sql = sql.Where("Wdifficulty = ?", wrongProblem.Wdifficulty)
	}
	if wrongProblem.Wunit != "不限" {
		sql = sql.Where("Wunit = ?", wrongProblem.Wunit)
	}
	sql.Find(&result)
	return
}

//搜索匹配（字符串，搜索域）
func (wrongProblem *WrongProblem) Search(Sid int64, searchString string, searchField string) (result []WrongProblem, err error) {

	if err = orm.Eloquent.Table("WRONGPROBLEM").Where("Sid = ? and "+searchField+" LIKE ?", Sid, "%"+searchString+"%").Find(&result).Error; err != nil {
		return
	}
	return
}
