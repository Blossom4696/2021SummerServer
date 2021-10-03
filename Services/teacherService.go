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

type Teacher struct {
	Tid            int64  `gorm:"primary_key;column:Tid;AUTO_INCREMENT"`
	Tname          string `gorm:"column:Tname"`
	Tnickname      string `gorm:"column:Tnickname"`
	Tpassword      string `gorm:"column:Tpassword"`
	TpasswordAgain string
	ToldPassword   string
	Tphone         string `gorm:"column:Tphone"`
	Ticon          string `gorm:"column:Ticon"`
	Cnames         []string
}

type WPPrintData struct {
	Wids []int64 `json:"Wids"`
}

// 教师登录
func (teacher Teacher) Login() (res Res, err error) {
	var teacherModel Models.Teacher

	teacherModel.Tname = teacher.Tname
	teacherModel.Tpassword = teacher.Tpassword

	result, err := teacherModel.QueryByName(teacher.Tname)

	if err != nil {
		return Res{
			Code: 2,
			Msg:  "Name not exist!",
			Data: nil,
		}, err
	}

	if result.Tpassword == teacherModel.Tpassword {

		// 创建一条aes加密数据
		signKey, err := Utils.NewKey(32)
		if err != nil {
			return Res{
				Code: -1,
				Msg:  "服务器出错",
				Data: nil,
			}, err
		}

		hashData := Utils.AesEncryptCBC([]byte(strconv.FormatInt(result.Tid, 10)), []byte(signKey))

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
		tokenData["Id"] = result.Tid
		tokenData["UserType"] = Config.TeacherType
		tokenData["SignKey"] = signKey
		tokenData["RdsKey"] = rdsKey

		tokenStr, err := Utils.GetToken(tokenData, Config.TokenSecret)

		var resData = make(map[string]interface{}, 2)
		resData["token"] = tokenStr
		resData["userinfo"] = result
		return Res{
			Code: 1,
			Msg:  "Teacher login success!",
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

// 教师注册
func (teacher Teacher) Register() (res Res, err error) {
	if teacher.Tpassword != teacher.TpasswordAgain {
		return Res{
			Code: 2,
			Msg:  "Password not consistent!",
			Data: nil,
		}, err
	}

	var teacherModel Models.Teacher

	teacherModel.Tname = teacher.Tname
	teacherModel.Tpassword = teacher.Tpassword
	teacherModel.Tphone = teacher.Tphone
	teacherModel.Ticon = teacher.Ticon
	teacherModel.Tnickname = teacher.Tnickname

	result, err := teacherModel.QueryAll()

	for _, x := range result {
		if x.Tname == teacherModel.Tname {
			return Res{
				Code: 2,
				Msg:  "Username existed!",
				Data: nil,
			}, err
		}
	}

	Tid, err := teacherModel.Insert()

	return Res{
		Code: 1,
		Msg:  "Register Success!",
		Data: Tid,
	}, err
}

// 教师根据校区查询学生列表
func (teacher Teacher) QueryStudentList() (result []Student, err error) {
	var teacherModel Models.Teacher

	tmpStudent, err := teacherModel.QueryStudentByTid(teacher.Tid)

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

// 教师根据校区查询学生列表（首先搜索，再排序）->参数（搜索字符串、排序名、排序方式）
func (teacher Teacher) QueryStudentSortedList(searchString string, sortName string, sortDir string) (result []Student, err error) {
	var teacherModel Models.Teacher

	fieldMap := make(map[string]string)
	fieldMap["姓名"] = "Snickname"
	fieldMap["年级"] = "Sgrade"
	fieldMap["校区"] = "Cname"

	tmpStudent, err := teacherModel.QueryStudentByTidSorted(teacher.Tid, searchString, "STUDENT.Snickname", fieldMap[sortName], sortDir)

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

// 教师登录测试
func (teacher Teacher) LoginTest() (res Res, err error) {
	return
}

//按Tid查
func (teacher Teacher) QueryByTid(Tid int64) (result Teacher, err error) {
	var teacherModel Models.Teacher

	tmpTeacher, err := teacherModel.QueryByTid(Tid)

	if err != nil {
		return
	}

	// 把teacher model转为teacher service
	err = Utils.CopyFields(&result, tmpTeacher)

	if err != nil {
		return
	}

	// 添加每个教师监管校区
	var teachAreaModel Models.TeachArea
	var campusModel Models.Campus
	cnameList := make([]string, 0)
	tas, err := teachAreaModel.QueryByTid(result.Tid)

	if err != nil {
		return
	}
	for j := 0; j < len(tas); j++ {
		var campus Models.Campus
		campus, err = campusModel.QueryByCid(tas[j].Cid)

		if err != nil {
			return
		}
		cnameList = append(cnameList, campus.Cname)
	}

	result.Cnames = cnameList

	return
}

//搜索学生
func (teacher Teacher) SearcherStudent(searchString string) (result []Student, err error) {
	var teacherModel Models.Teacher

	tmpStudent, err := teacherModel.SearchStudentByTid(teacher.Tid, searchString, "Snickname")

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

//搜索老师
func (teacher Teacher) Search(searchString string) (result []Teacher, err error) {
	var teacherModel Models.Teacher

	// 匹配昵称
	tmpTeacher, err := teacherModel.Search(searchString, "Tnickname")
	if err != nil {
		return
	}

	result = make([]Teacher, len(tmpTeacher))
	// 把teacher model转为teacher service
	for i := 0; i < len(tmpTeacher); i++ {
		err = Utils.CopyFields(&result[i], tmpTeacher[i])
		if err != nil {
			return
		}
	}

	return
}

// 编辑教师信息(From Admin)
func (teacher *Teacher) Update(Tid int64) (result Res, err error) {
	var teacherModel Models.Teacher

	teacherModel.Tid = Tid
	teacherModel.Tnickname = teacher.Tnickname
	teacherModel.Tphone = teacher.Tphone
	teacherModel.Ticon = teacher.Ticon

	teacherModel.Update(teacherModel.Tid)

	// 修改校区管理权限
	var taModel Models.TeachArea
	taModel.Tid = teacher.Tid
	tas, err := taModel.QueryByTid(teacher.Tid)
	if err != nil {
		return
	}

	// 返回数据
	resData := make(map[string][]Models.TeachArea)
	resData["insert"] = make([]Models.TeachArea, 0)
	// 增加的
	for i := 0; i < len(teacher.Cnames); i++ {

		taModel.Cid, err = strconv.ParseInt(teacher.Cnames[i], 10, 64)
		if err != nil {
			return
		}

		hasItem := false
		for j := 0; j < len(tas); j++ {
			if tas[j].Cid == taModel.Cid {
				hasItem = true
				break
			}
		}

		if !hasItem {
			TAid, _ := taModel.Insert()
			resData["insert"] = append(resData["insert"], Models.TeachArea{
				TAid: TAid,
				Tid:  taModel.Tid,
				Cid:  taModel.Cid,
			})
		}
	}

	resData["delete"] = make([]Models.TeachArea, 0)
	// 删除的
	for i := 0; i < len(tas); i++ {
		var taModel Models.TeachArea
		hasItem := false
		for j := 0; j < len(teacher.Cnames); j++ {
			var err error
			tmpCid, err := strconv.ParseInt(teacher.Cnames[j], 10, 64)
			if err != nil {
				return Res{
					Code: -1,
					Msg:  err.Error(),
					Data: nil,
				}, err
			}

			if tmpCid == tas[i].Cid {
				hasItem = true
				break
			}
		}

		if !hasItem {
			taModel.Delete(tas[i].TAid)
			resData["delete"] = append(resData["delete"], tas[i])
		}

	}

	return Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: resData,
	}, err
}

// 编辑教师信息
func (teacher *Teacher) UpdateFromTeacher(Tid int64) (result Res, err error) {
	var teacherModel Models.Teacher

	teacherModel.Tid = Tid
	teacherModel.Tnickname = teacher.Tnickname
	teacherModel.Tphone = teacher.Tphone
	teacherModel.Ticon = teacher.Ticon

	resTeacher, err := teacherModel.Update(teacherModel.Tid)

	return Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: resTeacher,
	}, err
}

func (teacher *Teacher) Delete(Tid int64) (result Models.Teacher, err error) {
	var teacherModel Models.Teacher

	teacherModel.Tid = teacher.Tid

	result, err = teacherModel.Delete(Tid)

	return
}

// 修改教师密码
func (teacher *Teacher) UpdatePasswordFromTeacher(Tid int64) (result Res, err error) {
	var teacherModel, tmpTeacher Models.Teacher

	teacherModel.Tid = Tid

	tmpTeacher, err = teacherModel.QueryByTid(Tid)
	if teacher.ToldPassword != tmpTeacher.Tpassword {
		return Res{
			Code: 2,
			Msg:  "Password not correct!",
			Data: nil,
		}, err
	}

	if teacher.Tpassword == "" {
		return Res{
			Code: 2,
			Msg:  "Password is empty!",
			Data: nil,
		}, err
	}

	if teacher.Tpassword != teacher.TpasswordAgain {
		return Res{
			Code: 2,
			Msg:  "Password not consistent!",
			Data: nil,
		}, err
	}

	if teacher.ToldPassword == teacher.Tpassword {
		return Res{
			Code: 2,
			Msg:  "The two passwords are correct!",
			Data: nil,
		}, err
	}

	teacherModel.Tpassword = teacher.Tpassword

	resTeacher, err := teacherModel.Update(teacherModel.Tid)

	return Res{
		Code: 1,
		Msg:  "Update Success!",
		Data: resTeacher,
	}, err
}
