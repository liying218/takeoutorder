package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"takeoutorder/dao"
	"takeoutorder/model"
)

func RepRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	sex := ctx.PostForm("sex")
	phonenum := ctx.PostForm("phonenum")
	count := ctx.PostForm("count")
	icount, _ := strconv.Atoi(count)
	rep := &model.Reporter{
		Name:     name,
		Sex:      sex,
		Phonenum: phonenum,
		Count:    int64(icount),
	}
	err := dao.DB.Create(&rep).Error
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
