package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGradeRouter(Router *gin.RouterGroup) {
	Grade := Router.Group("grade")
	{
		// ExcelRouter.POST("/importExcel", v1.ImportExcel)          // 导入Excel
		// ExcelRouter.GET("/loadExcel", v1.LoadExcel)               // 加载Excel数据
		// ExcelRouter.POST("/exportExcel", v1.ExportExcel)          // 导出Excel
		// ExcelRouter.GET("/downloadTemplate", v1.DownloadTemplate) // 下载模板文件

		Grade.GET("getGradeList", v1.GetGradeList) // 获取年级列表
		Grade.POST("createGrade", v1.CreateGrade)  // 新增年级信息
		Grade.PUT("customer", v1.UpGrade)          // 更新客户
		// Grade.GET("getGradeList", v1.GetGradeList) // 获取年级列表
	}
}
