package Services

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"

	"github.com/bigby/project/Config"
	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
)

type Student struct {
	Sid            int64  `gorm:"primary_key;column:Sid;AUTO_INCREMENT"`
	Sname          string `gorm:"column:Sname"`
	Snickname      string `gorm:"column:Snickname"`
	Spassword      string `gorm:"column:Spassword"`
	SpasswordAgain string
	SoldPassword   string
	Sphone         string `gorm:"column:Sphone"`
	Sicon          string `gorm:"column:Sicon"`
	Sgrade         int64  `gorm:"column:Sgrade"`
	Cid            int64  `gorm:"column:Cid"`
	Cname          string
	SgradeName     string
}

type Res struct {
	Code int64
	Msg  string
	Data interface{}
}

// 学生登录
// 根据校区id返回校区name
func (student Student) Login() (res Res, err error) {
	var studentModel Models.Student

	studentModel.Sname = student.Sname
	studentModel.Spassword = student.Spassword

	var result Student
	tmpStudent, err := studentModel.QueryByName(student.Sname)

	if err != nil {
		return Res{
			Code: 2,
			Msg:  "Name not exist!",
			Data: nil,
		}, err
	}

	// 把student model转为student service
	err = Utils.CopyFields(&result, tmpStudent)

	if err != nil {
		return Res{
			Code: -1,
			Msg:  "Query error!",
			Data: nil,
		}, err
	}

	// 把校区id转换为校区名
	var campusModel Models.Campus
	campus, err := campusModel.QueryByCid(result.Cid)

	if err != nil {
		return Res{
			Code: -1,
			Msg:  "Query error!",
			Data: nil,
		}, err
	}
	result.Cname = campus.Cname

	// 把年级数字转换为年级名
	result.SgradeName = Utils.GradeIntToString(result.Sgrade)

	if result.Spassword == studentModel.Spassword {

		// 创建一条aes加密数据
		signKey, err := Utils.NewKey(32)
		if err != nil {
			return Res{
				Code: -1,
				Msg:  "服务器出错",
				Data: nil,
			}, err
		}

		hashData := Utils.AesEncryptCBC([]byte(strconv.FormatInt(result.Sid, 10)), []byte(signKey))

		rdsVal := make(map[string]interface{}, 2)

		rdsVal["userData"], _ = json.Marshal(result)
		rdsVal["hashData"] = base64.StdEncoding.EncodeToString(hashData)

		rdsKey, err := Models.PutJSON(rdsVal, 3*time.Hour)
		if err != nil {
			return Res{
				Code: -1,
				Msg:  "服务器出错",
				Data: nil,
			}, err
		}

		tokenData := make(map[string]interface{}, 3)
		tokenData["Id"] = result.Sid
		tokenData["UserType"] = Config.StudentType
		tokenData["SignKey"] = signKey
		tokenData["RdsKey"] = rdsKey

		tokenStr, err := Utils.GetToken(tokenData, Config.TokenSecret)

		var resData = make(map[string]interface{}, 2)
		resData["token"] = tokenStr
		resData["userinfo"] = result
		return Res{
			Code: 1,
			Msg:  "Student login success!",
			Data: resData,
		}, err
	} else {
		return Res{
			Code: 2,
			Msg:  "Password not correct!",
			Data: nil,
		}, err
	}
}

// 学生注册
func (student Student) Register() (res Res, err error) {
	if student.Spassword != student.SpasswordAgain {
		return Res{
			Code: 2,
			Msg:  "Password not consistent!",
			Data: nil,
		}, err
	}

	var studentModel Models.Student

	studentModel.Sname = student.Sname
	studentModel.Spassword = student.Spassword
	studentModel.Sgrade = student.Sgrade
	studentModel.Sphone = student.Sphone
	studentModel.Sicon = student.Sicon
	studentModel.Snickname = student.Snickname
	studentModel.Cid = student.Cid

	result, err := studentModel.QueryAll()

	for _, x := range result {
		if x.Sname == studentModel.Sname {
			return Res{
				Code: 2,
				Msg:  "Username existed!",
				Data: nil,
			}, err
		}
	}

	Sid, err := studentModel.Insert()

	return Res{
		Code: 1,
		Msg:  "Register Success!",
		Data: Sid,
	}, err
}

//按Sid查
func (student Student) QueryBySid(Sid int64) (result Student, err error) {
	var studentModel Models.Student

	tmpStudent, err := studentModel.QueryBySid(Sid)

	if err != nil {
		return
	}

	// 把student model转为student service
	err = Utils.CopyFields(&result, tmpStudent)
	if err != nil {
		return
	}

	// 把校区id转换为校区名
	var campusModel Models.Campus
	campus, err := campusModel.QueryByCid(result.Cid)

	result.Cname = campus.Cname

	// 把年级数字转换为年级名
	result.SgradeName = Utils.GradeIntToString(result.Sgrade)

	return
}

// 搜索学生
func (student Student) Search(searchString string) (result []Student, err error) {
	var studentModel Models.Student

	// 匹配昵称
	tmpStudent, err := studentModel.Search(searchString, "Snickname")

	if err != nil {
		return
	}

	result = make([]Student, len(tmpStudent))
	// 把student model转为student service
	for i := 0; i < len(tmpStudent); i++ {
		err = Utils.CopyFields(&result[i], tmpStudent[i])
		if err != nil {
			return
		}
	}

	// 把校区id转换为校区名, 把年级数字转换为年级名
	var campusModel Models.Campus
	for i := 0; i < len(result); i++ {
		campus, err := campusModel.QueryByCid(result[i].Cid)
		if err != nil {
			return nil, err
		}
		result[i].Cname = campus.Cname

		result[i].SgradeName = Utils.GradeIntToString(result[i].Sgrade)
	}

	if err != nil {
		return nil, err
	}

	return
}

// 编辑学生信息
func (student *Student) Update(Sid int64) (result Res, err error) {
	var studentModel Models.Student

	studentModel.Sid = Sid
	studentModel.Snickname = student.Snickname
	studentModel.Sphone = student.Sphone
	studentModel.Sicon = student.Sicon
	studentModel.Cid = student.Cid
	studentModel.Sgrade = student.Sgrade

	resStudent, err := studentModel.Update(studentModel.Sid)
	if err != nil {
		return
	}

	return Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: resStudent,
	}, err
}

// 删除学生信息
func (student *Student) Delete(Sid int64) (result Models.Student, err error) {
	var studentModel Models.Student

	studentModel.Sid = student.Sid

	result, err = studentModel.Delete(Sid)

	return
}

// 修改学生密码
func (student *Student) UpdatePasswordFromStudent(Sid int64) (result Res, err error) {
	var studentModel, tmpStudent Models.Student

	studentModel.Sid = Sid

	tmpStudent, err = studentModel.QueryBySid(Sid)
	if student.SoldPassword != tmpStudent.Spassword {
		return Res{
			Code: 2,
			Msg:  "Password not correct!",
			Data: nil,
		}, err
	}

	if student.Spassword == "" {
		return Res{
			Code: 2,
			Msg:  "Password is empty!",
			Data: nil,
		}, err
	}

	if student.Spassword != student.SpasswordAgain {
		return Res{
			Code: 2,
			Msg:  "Password not consistent!",
			Data: nil,
		}, err
	}

	if student.SoldPassword == student.Spassword {
		return Res{
			Code: 2,
			Msg:  "The two passwords are correct!",
			Data: nil,
		}, err
	}

	studentModel.Spassword = student.Spassword

	resStudent, err := studentModel.Update(studentModel.Sid)

	return Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: resStudent,
	}, err
}
