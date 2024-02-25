package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"takeoutorder/dao"
	"takeoutorder/model"
)

func MenuRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	icount, _ := strconv.Atoi(ctx.PostForm("count"))
	ipid, _ := strconv.Atoi(ctx.PostForm("pid"))
	fprice, _ := strconv.ParseFloat(ctx.PostForm("price"), 64)
	menu := &model.Menu{
		Name:  name,
		Price: fprice,
		Count: int64(icount),
		Pid:   int64(ipid),
	}
	err := dao.DB.Create(&menu).Error
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
