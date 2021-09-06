package Controllers

import (
	"net/http"

	"github.com/bigby/project/Services"
	"github.com/gin-gonic/gin"
)

// @Summary 校区列表
// @Description 获取校区列表
// @Tags Campus
// @Accept json
// @Produce  json
// @Success 200 {object} Res{data=[]Models.Campus}
// @Router /app/campus/get_all_campus [get]
func CampusQuery(c *gin.Context) {
	var campusService Services.Campus

	result, err := campusService.QueryAll()

	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "No Data",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Query Success!",
		Data: result,
	})
}
