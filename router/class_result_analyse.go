package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitStudentResultAnalyseRouter(Router *gin.RouterGroup) {

	StudentResultAnalyse := Router.Group("studentResultAnalyse")
	{

		StudentResultAnalyse.GET("getStudentTotalResult", v1.GetStudentTotalResult) // 获取学生总成绩

	}
}
