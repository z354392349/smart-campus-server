package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitStudentResultAnalyseRouter(Router *gin.RouterGroup) {

	StudentResultAnalyse := Router.Group("studentResultAnalyse")
	{
		StudentResultAnalyse.GET("getStudentCourseResult", v1.GetStudentCourseResult)               // 获取学生全部科目成绩
		StudentResultAnalyse.GET("getStudentTotalResultHistory", v1.GetStudentTotalResultHistory)   // 获取学生全部考试总成绩
		StudentResultAnalyse.GET("getStudentCourseResultHistory", v1.GetStudentCourseResultHistory) // 获取学生全部考试单科成绩

	}
}
