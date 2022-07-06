package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
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
// @Param data body model.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func GetGradeList(c *gin.Context) {
	response.OkWithDetailed(response.PageResult{
		List:     "t123",
		Total:    1223,
		Page:     456,
		PageSize: 789,
	}, "获取成功", c)
}

// @Tags excel
// @Summary  创建年级
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body model.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func CreateGrade(c *gin.Context) {

	var grade model.Grade
	_ = c.ShouldBindJSON(&grade)

	// if err := utils.Verify(grade, utils.ApiVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }

	if err := service.CreateGrade(grade); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}
