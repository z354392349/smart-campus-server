package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDashboardRouter(Router *gin.RouterGroup) {
	Dashboard := Router.Group("dashboard")
	{

		Dashboard.GET("getDashboardCensusNum", v1.GetDashboardCensusNum)   // 获取通行数量，考勤
		Dashboard.GET("getTeacherNum", v1.GetTeacherNum)                   // 获取教师数量区分男女
		Dashboard.GET("getExamPassRate", v1.GetExamPassRate)               // 年级合格率
		Dashboard.GET("getTeacherAttendCensus", v1.GetTeacherAttendCensus) // 获取教师考勤率，准点率
		Dashboard.GET("getStudentNum", v1.GetStudentNum)                   // 获取学生数量
		Dashboard.GET("getStudentAttendCensus", v1.GetStudentAttendCensus) // 获取学生考勤数量
		// Dashboard.PUT("upCourse", v1.UpCourse)            // 更新课程
		// Dashboard.DELETE("deleteCourse", v1.DeleteCourse) // 删除课程
	}
}
