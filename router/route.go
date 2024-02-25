package router

import (
	"github.com/gin-gonic/gin"
	"takeoutorder/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 商家
	r.POST("/createpro", service.ProRegister)

	// 客户
	r.POST("/createcus", service.CusRegister)

	//外卖员
	r.POST("/createrep", service.RepRegister)

	//订单
	r.POST("/createorder", service.OrderRegister)

	//菜单
	r.POST("/createmenu", service.MenuRegister)
	return r

}
