package Models

import (
	"time"

	orm "github.com/bigby/project/Database"
)

type TodayExercise struct {
	TEid         int64     `gorm:"primary_key;column:TEid;AUTO_INCREMENT"`
	Eid          int64     `gorm:"column:Eid"`
	Sid          int64     `gorm:"column:Sid"`
	TEdate       time.Time `gorm:"column:TEdate"`
	TEanswerTxt  string    `gorm:"column:TEanswerTxt"`
	TEanswerPath string    `gorm:"column:TEanswerPath"`
}

//增
func (te TodayExercise) Insert() (TEid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("TODAYEXERCISE").Create(&te)
	TEid = te.TEid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (te *TodayExercise) Delete(TEid int64) (Result TodayExercise, err error) {

	if err = orm.Eloquent.Table("TODAYEXERCISE").Select([]string{"TEid"}).First(&te, TEid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("TODAYEXERCISE").Delete(&te).Error; err != nil {
		return
	}
	Result = *te
	return
}

//改
func (te *TodayExercise) UpdateByTEid(TEid int64) (updateTE TodayExercise, err error) {

	if err = orm.Eloquent.Table("TODAYEXERCISE").Select([]string{"TEid"}).First(&updateTE, TEid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("TODAYEXERCISE").Model(&updateTE).Updates(&te).Error; err != nil {
		return
	}
	return
}

//按Sid和Eid改
func (te *TodayExercise) UpdateBySidAndEid(Sid int64, Eid int64, TEdate time.Time) (updateTE TodayExercise, err error) {

	if err = orm.Eloquent.Table("TODAYEXERCISE").Where("Sid = ? and Eid = ? and TEdate = ?", Sid, Eid, TEdate.Format("2006-01-02")).First(&updateTE).Error; err != nil {
		return
	}

	//其他的id和date不动，只修改学生作答
	te.TEid = updateTE.TEid
	te.Eid = updateTE.Eid
	te.Sid = updateTE.Sid
	te.TEdate = updateTE.TEdate
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("TODAYEXERCISE").Model(&updateTE).Updates(&te).Error; err != nil {
		return
	}
	return
}

//查全部
func (te *TodayExercise) QueryAll() (teArray []TodayExercise, err error) {

	if err = orm.Eloquent.Table("TODAYEXERCISE").Find(&teArray).Error; err != nil {
		return
	}
	return
}

//按Sid查和日期查
func (te *TodayExercise) QueryBySidAndDate(Sid int64, date time.Time) (results []TodayExercise, err error) {

	if err = orm.Eloquent.Table("TODAYEXERCISE").Where("Sid = ? and TEdate = ?", Sid, date.Format("2006-01-02")).Find(&results).Error; err != nil {
		return
	}
	return
}
