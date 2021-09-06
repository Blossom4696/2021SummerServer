package Database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"fmt"
)

var Eloquent *gorm.DB

func InitMysql() {
	var err error
	//不要使用:=，使用=， 因为在上面有全局变量定义过，在这边重新定义，作用域为init()
	Eloquent, err = gorm.Open("mysql", "root:VRbigby4.@tcp(127.0.0.1:3306)/2021Summer?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
