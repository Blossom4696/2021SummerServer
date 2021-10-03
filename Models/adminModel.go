package Models

import (
	orm "github.com/bigby/project/Database"
)

type Admin struct {
	Aid       int64  `gorm:"primary_key;column:Aid;AUTO_INCREMENT"`
	Aname     string `gorm:"column:Aname"`
	Apassword string `gorm:"column:Apassword"`
}

//增
func (admin Admin) Insert() (Aid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("ADMIN").Create(&admin)
	Aid = admin.Aid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (admin *Admin) Delete(Aid int64) (Result Admin, err error) {

	if err = orm.Eloquent.Table("ADMIN").Select([]string{"Aid"}).First(&admin, Aid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("ADMIN").Delete(&admin).Error; err != nil {
		return
	}
	Result = *admin
	return
}

//改
func (admin *Admin) Update(Aid int64) (updateAdmin Admin, err error) {

	if err = orm.Eloquent.Table("ADMIN").Select([]string{"Aid"}).First(&updateAdmin, Aid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("ADMIN").Model(&updateAdmin).Updates(&admin).Error; err != nil {
		return
	}
	return
}

//查全部
func (admin *Admin) QueryAll() (admins []Admin, err error) {

	if err = orm.Eloquent.Table("ADMIN").Find(&admins).Error; err != nil {
		return
	}
	return
}

//按Tid查
func (admin *Admin) QueryByAid(Aid int64) (result Admin, err error) {

	if err = orm.Eloquent.Table("ADMIN").Where("Aid = ?", Aid).First(&result).Error; err != nil {
		return
	}
	return
}

//按Tname查
func (admin *Admin) QueryByName(Aname string) (result Admin, err error) {

	if err = orm.Eloquent.Table("ADMIN").Where("Aname = ?", Aname).First(&result).Error; err != nil {
		return
	}
	return
}
