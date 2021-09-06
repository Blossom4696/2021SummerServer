package Services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"

	"github.com/bigby/project/Config"
	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
)

type Admin struct {
	Aid       int64  `gorm:"primary_key;column:Aid;AUTO_INCREMENT"`
	Aname     string `gorm:"column:Aname"`
	Apassword string `gorm:"column:Apassword"`
}

// 管理员登录
func (admin Admin) Login() (res Res, err error) {
	var adminModel Models.Admin

	adminModel.Aname = admin.Aname
	adminModel.Apassword = admin.Apassword

	result, err := adminModel.QueryByName(adminModel.Aname)

	if err != nil {
		return Res{
			Code: -1,
			Msg:  "Query error!",
			Data: nil,
		}, err
	}

	if result.Apassword == adminModel.Apassword {

		// 创建一条aes加密数据
		signKey, err := Utils.NewKey(32)
		if err != nil {
			return Res{
				Code: -1,
				Msg:  "服务器出错",
				Data: nil,
			}, err
		}

		hashData := Utils.AesEncryptCBC([]byte(strconv.FormatInt(result.Aid, 10)), []byte(signKey))

		rdsVal := make(map[string]interface{}, 2)

		rdsVal["userData"], _ = json.Marshal(result)
		rdsVal["hashData"] = base64.StdEncoding.EncodeToString(hashData)

		var ctx = context.Background()
		rdsKey, err := Models.PutJSON(ctx, rdsVal, 3*time.Hour)
		if err != nil {
			return Res{
				Code: -1,
				Msg:  "服务器出错",
				Data: nil,
			}, err
		}

		tokenData := make(map[string]interface{}, 3)
		tokenData["Id"] = result.Aid
		tokenData["UserType"] = Config.AdminType
		tokenData["SignKey"] = signKey
		tokenData["RdsKey"] = rdsKey

		tokenStr, err := Utils.GetToken(tokenData, Config.TokenSecret)

		var resData = make(map[string]interface{}, 2)
		resData["token"] = tokenStr
		resData["userinfo"] = result
		return Res{
			Code: 1,
			Msg:  "Admin login success!",
			Data: resData,
		}, err
	} else {
		return Res{
			Code: 2,
			Msg:  "Admin login failed!",
			Data: nil,
		}, err
	}
}

// 获取学生列表
func (admin Admin) QueryStudentList() (result []Student, err error) {
	var studentModel Models.Student

	tmpStudent, err := studentModel.QueryAll()

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

// 获取学生列表（排序）
func (admin Admin) QueryStudentSortedList(sortName string, sortDir string) (result []Student, err error) {
	var studentModel Models.Student

	fieldMap := make(map[string]string)
	fieldMap["姓名"] = "Snickname"
	fieldMap["年级"] = "Sgrade"
	fieldMap["校区"] = "Cname"

	tmpStudent, err := studentModel.QueryAllSorted(fieldMap[sortName], sortDir)
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

// 获取教师列表
func (admin Admin) QueryTeacherList() (result []Teacher, err error) {
	var teacherModel Models.Teacher

	tmpTeacher, err := teacherModel.QueryAll()

	result = make([]Teacher, len(tmpTeacher))
	// 把teacher model转为teacher service
	for i := 0; i < len(tmpTeacher); i++ {
		err = Utils.CopyFields(&result[i], tmpTeacher[i])
		if err != nil {
			return
		}
	}

	if err != nil {
		return nil, err
	}

	// 添加每个教师监管校区
	var teachAreaModel Models.TeachArea
	var campusModel Models.Campus
	for i := 0; i < len(result); i++ {
		cnameList := make([]string, 0)
		tas, err := teachAreaModel.QueryByTid(result[i].Tid)

		if err != nil {
			return nil, err
		}
		for j := 0; j < len(tas); j++ {
			campus, err := campusModel.QueryByCid(tas[j].Cid)

			if err != nil {
				return nil, err
			}
			cnameList = append(cnameList, campus.Cname)
		}

		result[i].Cnames = cnameList
	}

	return
}

// 获取教师列表（排序）
func (admin Admin) QueryTeacherSortedList(sortName string, sortDir string) (result []Teacher, err error) {
	var teacherModel Models.Teacher

	fieldMap := make(map[string]string)
	fieldMap["姓名"] = "Tnickname"
	fieldMap["校区"] = "Cname"

	tmpTeacher, err := teacherModel.QueryAllSorted(fieldMap[sortName], sortDir)

	result = make([]Teacher, len(tmpTeacher))
	// 把teacher model转为teacher service
	for i := 0; i < len(tmpTeacher); i++ {
		err = Utils.CopyFields(&result[i], tmpTeacher[i])
		if err != nil {
			return
		}
	}

	if err != nil {
		return nil, err
	}

	// 添加每个教师监管校区
	var teachAreaModel Models.TeachArea
	var campusModel Models.Campus
	for i := 0; i < len(result); i++ {
		cnameList := make([]string, 0)
		tas, err := teachAreaModel.QueryByTid(result[i].Tid)

		if err != nil {
			return nil, err
		}
		for j := 0; j < len(tas); j++ {
			campus, err := campusModel.QueryByCid(tas[j].Cid)

			if err != nil {
				return nil, err
			}
			cnameList = append(cnameList, campus.Cname)
		}

		result[i].Cnames = cnameList
	}

	return
}
