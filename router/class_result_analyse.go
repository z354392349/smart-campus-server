package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitClassResultAnalyseRouter(Router *gin.RouterGroup) {

	ClassResultAnalyse := Router.Group("classResultAnalyse")
	{

		ClassResultAnalyse.GET("getClassTotalResult", v1.GetClassTotalResult)                 // 获取学生总成绩
		ClassResultAnalyse.GET("getClassPassPercent", v1.GetClassPassPercent)                 // 获取班级通过率
		ClassResultAnalyse.GET("getClassToTalResultHistory", v1.GetClassToTalResultHistory)   // 获取学生考试总成绩-历史
		ClassResultAnalyse.GET("getClassCourseResultHistory", v1.GetClassCourseResultHistory) // 获取学生考试单科成绩-历史

	}
}
