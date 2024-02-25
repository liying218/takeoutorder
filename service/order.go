package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"takeoutorder/dao"
	"takeoutorder/model"
)

func OrderRegister(ctx *gin.Context) {
	ipid, _ := strconv.Atoi(ctx.PostForm("pid"))
	icid, _ := strconv.Atoi(ctx.PostForm("cid"))
	irid, _ := strconv.Atoi(ctx.PostForm("rid"))
	foods := ctx.PostForm("foods")
	fprice, _ := strconv.ParseFloat(ctx.PostForm("price"), 64)
	state := ctx.PostForm("state")
	order := &model.Order{
		Pid:   int64(ipid),
		Cid:   int64(icid),
		Rid:   int64(irid),
		Foods: foods,
		Price: fprice,
		State: state,
	}
	err := dao.DB.Create(&order).Error
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
