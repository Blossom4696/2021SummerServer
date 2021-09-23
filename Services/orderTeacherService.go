package Services

import (
	"time"

	"github.com/bigby/project/Models"
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

func (ot *OrderTeacher) QueryListByDate() (result []Models.OrderTeacher, err error) {
	var otModel Models.OrderTeacher

	result, err = otModel.QueryByDate(time.Now())

	return
}

func (ot *OrderTeacher) QueryByOTid(id int64) (result Models.OrderTeacher, err error) {
	var otModel Models.OrderTeacher

	result, err = otModel.QueryByOTid(id)

	return
}


func (ot *OrderTeacher) Insert() (OTid int64, err error) {
	var otModel Models.OrderTeacher

	otModel.Sid = ot.Sid
	otModel.OTweek = ot.OTweek
	otModel.OTtimeSlot = ot.OTtimeSlot
	otModel.OTsubject = ot.OTsubject
	otModel.OTcontactName = ot.OTcontactName
	otModel.OTcontactPhone = ot.OTcontactPhone
	otModel.OTaddress = ot.OTaddress
	otModel.OTcreateTime = ot.OTcreateTime

	OTid, err = otModel.Insert()

	return

}