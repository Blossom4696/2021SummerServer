package Models

import (
	orm "github.com/bigby/project/Database"
)

type Teacher struct {
	Tid       int64  `gorm:"primary_key;column:Tid;AUTO_INCREMENT"`
	Tname     string `gorm:"column:Tname"`
	Tnickname string `gorm:"column:Tnickname"`
	Tpassword string `gorm:"column:Tpassword"`
	Tphone    string `gorm:"column:Tphone"`
	Ticon     string `gorm:"column:Ticon"`
}

//增
func (teacher Teacher) Insert() (Tid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("TEACHER").Create(&teacher)
	Tid = teacher.Tid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (teacher *Teacher) Delete(Tid int64) (Result Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Select([]string{"Tid"}).First(&teacher, Tid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("TEACHER").Delete(&teacher).Error; err != nil {
		return
	}
	Result = *teacher
	return
}

//改
func (teacher *Teacher) Update(Tid int64) (updateTeacher Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Select([]string{"Tid"}).First(&updateTeacher, Tid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("TEACHER").Model(&updateTeacher).Updates(&teacher).Error; err != nil {
		return
	}
	return
}

//查全部
func (teacher *Teacher) QueryAll() (teachers []Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Find(&teachers).Error; err != nil {
		return
	}
	return
}

//查全部(排序)
func (teacher *Teacher) QueryAllSorted(sortName string, sortDir string) (result []Teacher, err error) {
	if sortName == "Cname" {
		sql := orm.Eloquent.Table("TEACHER")
		if sortDir != "不限" {
			sub := orm.Eloquent.Table("CAMPUS").Select("Cid").Where("Cname = ?", sortDir).SubQuery()
			sub = orm.Eloquent.Table("TEACHAREA").Select("Tid").Where("Cid in ?", sub).SubQuery()
			sql = sql.Where("Tid in ?", sub)
		}
		sql.Find(&result)
	} else {
		err = orm.Eloquent.Table("TEACHER").Order(sortName + " " + sortDir).Find(&result).Error
	}

	return
}

//按Tid查
func (teacher *Teacher) QueryByTid(Tid int64) (result Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Where("Tid = ?", Tid).First(&result).Error; err != nil {
		return
	}
	return
}

//按Tname查
func (teacher *Teacher) QueryByName(Tname string) (result Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Where("Tname = ?", Tname).First(&result).Error; err != nil {
		return
	}
	return
}

//查监管的全部学生
//按Tid查校区，找出校区的全部学生
func (teacher *Teacher) QueryStudentByTid(Tid int64) (result []Student, err error) {
	sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()

	err = orm.Eloquent.Table("STUDENT").Where("Cid in ?", sub).Find(&result).Error

	return
}

//查监管的全部学生(排序)
//按Tid查校区，找出校区的全部学生(排序)（教师id，字符串，搜索域，排序名，排序方式）
func (teacher *Teacher) QueryStudentByTidSorted(Tid int64, searchString string, searchField string, sortName string, sortDir string) (result []Student, err error) {
	if sortName == "Cname" {
		sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()
		query := orm.Eloquent.Table("STUDENT").Where("STUDENT.Cid in ? and "+searchField+" LIKE ?", sub, "%"+searchString+"%").Order(sortName + " " + sortDir)
		err = query.Joins("JOIN CAMPUS on STUDENT.Cid = CAMPUS.Cid").Select("STUDENT.*, CAMPUS.Cname").Order(sortName + " " + sortDir).Find(&result).Error
	} else {
		sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()
		err = orm.Eloquent.Table("STUDENT").Where("Cid in ? and "+searchField+" LIKE ?", sub, "%"+searchString+"%").Order(sortName + " " + sortDir).Find(&result).Error
	}

	return
}

//搜索监管的全部学生
//按Tid查校区，搜索校区的全部学生（教师id，字符串，搜索域）
func (teacher *Teacher) SearchStudentByTid(Tid int64, searchString string, searchField string) (result []Student, err error) {
	sub := orm.Eloquent.Table("TEACHAREA").Select("Cid").Where("Tid = ?", Tid).SubQuery()

	err = orm.Eloquent.Table("STUDENT").Where("Cid in ? and "+searchField+" LIKE ?", sub, "%"+searchString+"%").Find(&result).Error

	return
}

//搜索匹配（字符串，搜索域）
func (teacher *Teacher) Search(searchString string, searchField string) (result []Teacher, err error) {

	if err = orm.Eloquent.Table("TEACHER").Where(searchField+" LIKE ?", "%"+searchString+"%").Find(&result).Error; err != nil {
		return
	}
	return
}
