package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"takeoutorder/dao"
	"takeoutorder/model"
)

func CusRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	address := ctx.PostForm("address")
	sex := ctx.PostForm("sex")
	phonenum := ctx.PostForm("phonenum")
	money := ctx.PostForm("money")
	fmoney, _ := strconv.ParseFloat(money, 64)
	cus := &model.Customer{
		Name:     name,
		Address:  address,
		Sex:      sex,
		Phonenum: phonenum,
		Money:    fmoney,
	}
	err := dao.DB.Create(&cus).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "create error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "create success",
	})

}
