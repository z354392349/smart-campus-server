package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCourseRouter(Router *gin.RouterGroup) {
	Grade := Router.Group("course")
	{

		Grade.GET("getCourseList", v1.GetCourseList)  // 获取课程列表
		Grade.POST("createCourse", v1.CreateCourse)   // 新增课程信息
		Grade.PUT("upCourse", v1.UpCourse)            // 更新课程
		Grade.DELETE("deleteCourse", v1.DeleteCourse) // 删除课程
	}
}
