package Models

import (
	"time"

	orm "github.com/bigby/project/Database"
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

//增
func (ls LearnSituation) Insert() (LSid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("LEARNSITUATION").Create(&ls)
	LSid = ls.LSid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (ls *LearnSituation) Delete(LSid int64) (Result LearnSituation, err error) {

	if err = orm.Eloquent.Table("LEARNSITUATION").Select([]string{"LSid"}).First(&ls, LSid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("LEARNSITUATION").Delete(&ls).Error; err != nil {
		return
	}
	Result = *ls
	return
}

//改
func (ls *LearnSituation) Update(LSid int64) (updateLS LearnSituation, err error) {

	if err = orm.Eloquent.Table("LEARNSITUATION").Select([]string{"LSid"}).First(&updateLS, LSid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("LEARNSITUATION").Model(&updateLS).Updates(&ls).Error; err != nil {
		return
	}
	return
}

//查全部
func (ls *LearnSituation) QueryAll() (lsArray []LearnSituation, err error) {

	if err = orm.Eloquent.Table("LEARNSITUATION").Find(&lsArray).Error; err != nil {
		return
	}
	return
}

//按LSid查
func (ls *LearnSituation) QueryByLSid(LSid int64) (result LearnSituation, err error) {
	if err = orm.Eloquent.Table("LEARNSITUATION").Where("LSid = ?", LSid).First(&result).Error; err != nil {
		return
	}

	return
}

//按Sid查
func (ls *LearnSituation) QueryBySid(Sid int64) (results map[string][]LearnSituation, err error) {
	results = make(map[string][]LearnSituation)

	dailyLS := make([]LearnSituation, 0)
	if err = orm.Eloquent.Table("LEARNSITUATION").Where("Sid = ? and not(LSattendence = '' and LSperform = '' and LShomework = '')", Sid).Find(&dailyLS).Error; err != nil {
		return
	}

	examLS := make([]LearnSituation, 0)
	if err = orm.Eloquent.Table("LEARNSITUATION").Where("Sid = ? and not(LSexam = '' and LSsubject = '' and LSgrade = '')", Sid).Find(&examLS).Error; err != nil {
		return
	}
	results["daily"] = dailyLS
	results["exam"] = examLS
	return
}
