package Models

import (
	"time"

	orm "github.com/bigby/project/Database"
)

type OrderTeacher struct {
	OTid           int64     `gorm:"primary_key;column:OTid;AUTO_INCREMENT"`
	Sid            int64     `gorm:"column:Sid"`
	OTweek         string    `gorm:"column:OTweek"`
	OTtimeSlot     string    `gorm:"column:OTtimeSlot"`
	OTsubject      string    `gorm:"column:OTsubject"`
	OTcontactName  string    `gorm:"column:OTcontactName"`
	OTcontactPhone string    `gorm:"column:OTcontactPhone"`
	OTaddress      string    `gorm:"type:enum('思明','湖里','集美','海沧','同安','翔安');column:OTaddress;default:NULL"`
	OTcreateTime   time.Time `gorm:"column:OTcreateTime"`
}

//增
func (ot OrderTeacher) Insert() (OTid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("ORDERTEACHER").Create(&ot)
	OTid = ot.OTid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (ot *OrderTeacher) Delete(OTid int64) (Result OrderTeacher, err error) {

	if err = orm.Eloquent.Table("ORDERTEACHER").Select([]string{"OTid"}).First(&ot, OTid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("ORDERTEACHER").Delete(&ot).Error; err != nil {
		return
	}
	Result = *ot
	return
}

//改
func (ot *OrderTeacher) UpdateByOTid(OTid int64) (updateOT OrderTeacher, err error) {

	if err = orm.Eloquent.Table("ORDERTEACHER").Select([]string{"OTid"}).First(&updateOT, OTid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("ORDERTEACHER").Model(&updateOT).Updates(&ot).Error; err != nil {
		return
	}
	return
}

//查全部
func (ot *OrderTeacher) QueryAll() (otArray []OrderTeacher, err error) {

	if err = orm.Eloquent.Table("ORDERTEACHER").Find(&otArray).Error; err != nil {
		return
	}
	return
}

//按OTid查
func (ot *OrderTeacher) QueryByOTid(id int64) (result OrderTeacher, err error) {

	if err = orm.Eloquent.Table("ORDERTEACHER").Where("OTid = ?", id).First(&result).Error; err != nil {
		return
	}
	return
}

//按日期查
func (ot *OrderTeacher) QueryByDate(date time.Time) (results []OrderTeacher, err error) {

	if err = orm.Eloquent.Table("ORDERTEACHER").Where("OTcreateTime < ?", date.Format("2006-01-02 15:04:05")).Find(&results).Error; err != nil {
		return
	}
	return
}
