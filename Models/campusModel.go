package Models

import (
	orm "github.com/bigby/project/Database"
)

type Campus struct {
	Cid   int64  `gorm:"primary_key;column:Cid;AUTO_INCREMENT"`
	Cname string `gorm:"column:Cname"`
}

//按Cid查
func (campus *Campus) QueryByCid(Cid int64) (result Campus, err error) {

	if err = orm.Eloquent.Table("CAMPUS").Where("Cid = ?", Cid).First(&result).Error; err != nil {
		return
	}
	return
}

// 查全部
func (campus *Campus) QueryAll() (results []Campus, err error) {

	if err = orm.Eloquent.Table("CAMPUS").Find(&results).Error; err != nil {
		return
	}
	return
}

//增
func (campus Campus) Insert() (Cid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("CAMPUS").Create(&campus)
	Cid = campus.Cid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (campus *Campus) Delete(Cid int64) (Result Campus, err error) {

	if err = orm.Eloquent.Table("CAMPUS").Select([]string{"Cid"}).First(&campus, Cid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("CAMPUS").Delete(&campus).Error; err != nil {
		return
	}
	Result = *campus
	return
}

//改
func (campus *Campus) Update(Cid int64) (updateCampus Campus, err error) {

	if err = orm.Eloquent.Table("CAMPUS").Select([]string{"Cid"}).First(&updateCampus, Cid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("CAMPUS").Model(&updateCampus).Updates(&campus).Error; err != nil {
		return
	}
	return
}
