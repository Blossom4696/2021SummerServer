package Middlewares

import (
	"context"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/bigby/project/Config"
	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
	"github.com/gin-gonic/gin"
)

// Gin的中间件
func JWT(userType Config.UserType) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Request.Header.Get("TTToken")
		if token == "" {
			// 如果未取到，说明用户未登录
			resErr(c, "请登录")
			return
		}

		// 解析token
		claims, err := Utils.ParseToken(token, Config.TokenSecret)

		// 如果token解析失败
		if err != nil {
			resErr(c, "token检测失败，请重新登录")
			return
		}

		// 如果token时间超过有效期
		if time.Now().Unix() > claims.ExpiresAt {
			resErr(c, "token过期")
			return
		}

		//从token中获取数据
		tokenRdsKey := claims.RdsKey
		tokenSignKey := claims.SignKey
		tokenUserId := claims.UserID
		tokenUserType := claims.UserType

		// 校验权限，学生-学生，教师-学生、教师，管理员-学生、教师、管理员
		switch userType {
		case Config.StudentType:
			if !(tokenUserType == Config.StudentType || tokenUserType == Config.TeacherType || tokenUserType == Config.AdminType) {
				resErr(c, "权限不够")
				return
			}
		case Config.TeacherType:
			if tokenUserType == Config.StudentType {
				resErr(c, "权限不够")
				return
			}
		case Config.AdminType:
			if tokenUserType == Config.TeacherType || tokenUserType == Config.StudentType {
				resErr(c, "权限不够")
				return
			}
		default:
			resErr(c, "权限出错")
			return
		}

		// 从redis数据库中拿到aes加密的hashData、用户数据
		var ctx = context.Background()
		rdsData, err := Models.GetJSON(ctx, string(tokenRdsKey))
		// 如果redis中存储过期
		if err != nil {
			resErr(c, err.Error())
			return
		}

		// 取到hashData
		hashData := rdsData["hashData"]
		hashDataBytes, err := base64.StdEncoding.DecodeString(hashData) // 转为byte[]

		if err != nil {
			resErr(c, "服务器发生错误")
			return
		}

		// 验证aes加密
		if decrypted, err := Utils.AesDecryptCBC(hashDataBytes, []byte(tokenSignKey)); err != nil {
			resErr(c, "服务器发生错误")
			return
		} else {
			uid, err := strconv.ParseInt(string(decrypted), 10, 64)
			if err != nil {
				resErr(c, "服务器发生错误")
				return
			}
			// 如果用户id不同，验证失败
			if uid != tokenUserId {
				resErr(c, "验证失败")
				return
			}
		}

		//登录验证完成，把用户数据提取出来
		userData := rdsData["userData"]
		c.Request.Header.Add("Userdata", userData)
	}
}

func resErr(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": -1,
		"msg":  msg,
		"data": nil,
	})
	//
	c.Abort()
}
