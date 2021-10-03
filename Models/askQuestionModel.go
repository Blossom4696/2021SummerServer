package Models

import (
	"time"

	orm "github.com/bigby/project/Database"
)

type AskQuestion struct {
	AQid       int64     `gorm:"primary_key;column:AQid;AUTO_INCREMENT"`
	Sid        int64     `gorm:"column:Sid"`
	Eid        int64     `gorm:"column:Eid"`
	AQtime     time.Time `gorm:"column:AQtime"`
	AQremark   string    `gorm:"column:AQremark"`
	AQisSolved int64     `gorm:"column:AQisSolved"`
}

//增
func (aq AskQuestion) Insert() (AQid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("ASKQUESTION").Create(&aq)
	AQid = aq.AQid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (aq *AskQuestion) Delete(AQid int64) (Result AskQuestion, err error) {

	if err = orm.Eloquent.Table("ASKQUESTION").Select([]string{"AQid"}).First(&aq, AQid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("ASKQUESTION").Delete(&aq).Error; err != nil {
		return
	}
	Result = *aq
	return
}

//改
func (aq *AskQuestion) UpdateByAQid(AQid int64) (updateAQ AskQuestion, err error) {

	if err = orm.Eloquent.Table("ASKQUESTION").Select([]string{"AQid"}).First(&updateAQ, AQid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("ASKQUESTION").Model(&updateAQ).Updates(&aq).Error; err != nil {
		return
	}
	return
}

//查全部
func (aq *AskQuestion) QueryAll() (aqArray []AskQuestion, err error) {

	if err = orm.Eloquent.Table("ASKQUESTION").Find(&aqArray).Error; err != nil {
		return
	}
	return
}

//按AQid查
func (aq *AskQuestion) QueryByAQid(AQid int64) (result AskQuestion, err error) {
	if err = orm.Eloquent.Table("ASKQUESTION").Where("AQid = ?", AQid).First(&result).Error; err != nil {
		return
	}

	return
}

//查找老师监管内学生的全部提问问题
func (aq *AskQuestion) QueryFromTeacher(Tid int64, date string) (result []AskQuestion, err error) {

	sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()

	sub = orm.Eloquent.Table("STUDENT").Select("Sid").Where("Cid in ?", sub).SubQuery()

	if date == "day" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ?", sub, 1).Find(&result).Error
	} else if date == "week" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ?", sub, 7).Find(&result).Error
	} else if date == "month" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ?", sub, 30).Find(&result).Error
	} else if date == "all" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ?", sub).Find(&result).Error
	}

	return
}

//查找老师监管内学生的未解决问题
func (aq *AskQuestion) QueryUnresolvedFromTeacher(Tid int64, date string) (result []AskQuestion, err error) {

	sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()

	sub = orm.Eloquent.Table("STUDENT").Select("Sid").Where("Cid in ?", sub).SubQuery()

	if date == "day" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ? and AQisSolved = 0", sub, 1).Find(&result).Error
	} else if date == "week" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ? and AQisSolved = 0", sub, 7).Find(&result).Error
	} else if date == "month" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and DATEDIFF(now(), AQtime) <= ? and AQisSolved = 0", sub, 30).Find(&result).Error
	} else if date == "all" {
		err = orm.Eloquent.Table("ASKQUESTION").Where("Sid in ? and AQisSolved = 0", sub).Find(&result).Error
	}

	return
}
