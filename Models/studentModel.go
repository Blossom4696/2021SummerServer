package Models

import (
	orm "github.com/bigby/project/Database"
)

type Student struct {
	Sid       int64  `gorm:"primary_key;column:Sid;AUTO_INCREMENT"`
	Sname     string `gorm:"column:Sname"`
	Snickname string `gorm:"column:Snickname"`
	Spassword string `gorm:"column:Spassword"`
	Sphone    string `gorm:"column:Sphone"`
	Sicon     string `gorm:"column:Sicon"`
	Sgrade    int64  `gorm:"column:Sgrade"`
	Cid       int64  `gorm:"column:Cid"`
}

//增
func (student Student) Insert() (Sid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("STUDENT").Create(&student)
	Sid = student.Sid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (student *Student) Delete(Sid int64) (Result Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Select([]string{"Sid"}).First(&student, Sid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("STUDENT").Delete(&student).Error; err != nil {
		return
	}
	Result = *student
	return
}

//改
func (student *Student) Update(Sid int64) (updateStudent Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Select([]string{"Sid"}).First(&updateStudent, Sid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("STUDENT").Model(&updateStudent).Updates(&student).Error; err != nil {
		return
	}
	return
}

//查全部
func (student *Student) QueryAll() (students []Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Find(&students).Error; err != nil {
		return
	}
	return
}

//查全部(排序)
func (student *Student) QueryAllSorted(sortName string, sortDir string) (result []Student, err error) {
	if sortName == "Cname" {
		query := orm.Eloquent.Table("STUDENT").Order(sortName + " " + sortDir)
		err = query.Joins("JOIN CAMPUS on STUDENT.Cid = CAMPUS.Cid").Select("STUDENT.*, CAMPUS.Cname").Order(sortName + " " + sortDir).Find(&result).Error
	} else {
		err = orm.Eloquent.Table("STUDENT").Order(sortName + " " + sortDir).Find(&result).Error
	}

	return
}

//按一个Sid查
func (student *Student) QueryBySid(Sid int64) (result Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Where("Sid = ?", Sid).First(&result).Error; err != nil {
		return
	}
	return
}

//按多个Sid查
func (student *Student) QueryByMultiSid(Sid []int64) (result []Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Where("Sid = ?", Sid).Find(&result).Error; err != nil {
		return
	}
	return
}

//按Sname查
func (student *Student) QueryByName(Sname string) (result Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Where("Sname = ?", Sname).First(&result).Error; err != nil {
		return
	}
	return
}

//搜索匹配（字符串，搜索域）
func (student *Student) Search(searchString string, searchField string) (result []Student, err error) {

	if err = orm.Eloquent.Table("STUDENT").Where(searchField+" LIKE ?", "%"+searchString+"%").Find(&result).Error; err != nil {
		return
	}
	return
}
