package Config

var (
	TokenSecret     = "2021Summer" //tokenKey
	TokenExpireTime = 3 * 3600     //token有效期——3小时
)

// 用户类型，用于jwt判断权限
type UserType int

const (
	AdminType UserType = iota
	TeacherType
	StudentType
)
