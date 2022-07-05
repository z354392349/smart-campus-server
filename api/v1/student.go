package v1

import (
	"gin-vue-admin/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body model.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func GetStudentList(c *gin.Context) {
	response.OkWithDetailed(response.PageResult{
		List:     "t123",
		Total:    1223,
		Page:     456,
		PageSize: 789,
	}, "获取成功", c)
}
