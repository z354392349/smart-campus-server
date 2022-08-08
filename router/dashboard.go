package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDashboardRouter(Router *gin.RouterGroup) {
	Dashboard := Router.Group("dashboard")
	{

		Dashboard.GET("getDashboardCensusNum", v1.GetDashboardCensusNum) // 获取课程列表
		Dashboard.GET("getTeacherNum", v1.GetTeacherNum)                 // 新增课程信息
		// Dashboard.POST("createCourse", v1.CreateCourse)   // 新增课程信息
		// Dashboard.PUT("upCourse", v1.UpCourse)            // 更新课程
		// Dashboard.DELETE("deleteCourse", v1.DeleteCourse) // 删除课程
	}
}
