package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitStudentRouter(Router *gin.RouterGroup) {
	ExcelRouter := Router.Group("student")
	{
		// ExcelRouter.POST("/importExcel", v1.ImportExcel)          // 导入Excel
		// ExcelRouter.GET("/loadExcel", v1.LoadExcel)               // 加载Excel数据
		// ExcelRouter.POST("/exportExcel", v1.ExportExcel)          // 导出Excel
		// ExcelRouter.GET("/downloadTemplate", v1.DownloadTemplate) // 下载模板文件

		ExcelRouter.GET("/getStudentList", v1.GetClassList1) // 获取学生信息

	}
}
