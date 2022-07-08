package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags excel
// @Summary 获取年级列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Success 200
// @Router /excel/exportExcel [post]

func GetGradeList(c *gin.Context) {

	var pageInfo request.SearchGradeParams
	_ = c.ShouldBindQuery(&pageInfo)
	fmt.Println("--------------------")
	fmt.Printf("%+v", pageInfo)
	if err, list, total := service.GetCreateList(pageInfo); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags excel
// @Summary  创建年级
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Success 200
// @Router /excel/exportExcel [post]

func CreateGrade(c *gin.Context) {

	var grade model.Grade
	_ = c.ShouldBindJSON(&grade)

	if err := service.CreateGrade(grade); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Tags excel
// @Summary  更新年级
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Success 200
// @Router /excel/exportExcel [post]

func UpGrade(c *gin.Context) {

	var grade model.Grade
	_ = c.ShouldBindJSON(&grade)

	if err := service.CreateGrade(grade); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}
