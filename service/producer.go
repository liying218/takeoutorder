package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"takeoutorder/dao"
	"takeoutorder/model"
)

// 注册
func ProRegister(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	phonenum := c.PostForm("phonenum")
	address := c.PostForm("address")
	money := c.PostForm("money")
	fmoney, _ := strconv.ParseFloat(money, 64)
	pro := &model.Producer{
		Name:     name,
		Phonenum: phonenum,
		Address:  address,
		Money:    fmoney,
	}
	err := dao.DB.Create(&pro).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create Producer Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Create Producer Success",
	})
}
