package api

import (
	"github.com/gin-gonic/gin"
	"go/model"
)

func GetUsers(c *gin.Context){
	users, err := model.AllUsers(c)
	if err != nil {
		c.JSON(200, ApiRes{
			Code:1000,
			Msg:"发生了错误",
		})
		return
	}
	c.JSON(200, ApiRes{
		Code:200,
		Msg:"数据获取成功",
		Data: gin.H{
			"list": users,
			"page_size":1,
			"page_num":10,
		},

	})
	return
}
