package Models

import (
	orm "github.com/bigby/project/Database"
)

type TeachArea struct {
	TAid int64 `gorm:"primary_key;column:TAid;AUTO_INCREMENT"`
	Tid  int64 `gorm:"column:Tid"`
	Cid  int64 `gorm:"column:Cid"`
}

//按Tid查
func (teachArea *TeachArea) QueryByTid(Tid int64) (result []TeachArea, err error) {

	orm.Eloquent.SubQuery()
	if err = orm.Eloquent.Table("TEACHAREA").Where("Tid = ?", Tid).Find(&result).Error; err != nil {
		return
	}
	return
}

//增
func (teachArea TeachArea) Insert() (TAid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("TEACHAREA").Create(&teachArea)
	TAid = teachArea.TAid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (teachArea *TeachArea) Delete(TAid int64) (Result TeachArea, err error) {

	if err = orm.Eloquent.Table("TEACHAREA").Select([]string{"TAid"}).First(&teachArea, TAid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("TEACHAREA").Delete(&teachArea).Error; err != nil {
		return
	}
	Result = *teachArea
	return
}

//改
func (teachArea *TeachArea) Update(TAid int64) (updateTA TeachArea, err error) {

	if err = orm.Eloquent.Table("TEACHAREA").Select([]string{"TAid"}).First(&updateTA, TAid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("TEACHAREA").Model(&updateTA).Updates(&teachArea).Error; err != nil {
		return
	}
	return
}
