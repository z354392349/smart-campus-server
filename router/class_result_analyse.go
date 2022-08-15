package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitStudentResultAnalyseRouter(Router *gin.RouterGroup) {

	StudentResultAnalyse := Router.Group("studentResultAnalyse")
	{

		StudentResultAnalyse.GET("getStudentTotalResult", v1.GetStudentTotalResult)                 // 获取学生总成绩
		StudentResultAnalyse.GET("getClassPassPercent", v1.GetClassPassPercent)                     // 获取班级通过率
		StudentResultAnalyse.GET("getStudentToTalResultHistory", v1.GetStudentToTalResultHistory)   // 获取学生考试总成绩-历史
		StudentResultAnalyse.GET("getStudentCourseResultHistory", v1.GetStudentCourseResultHistory) // 获取学生考试单科成绩-历史

	}
}
