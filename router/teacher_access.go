package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTeacherAccessRouter(Router *gin.RouterGroup) {
	TeacherAccess := Router.Group("teacherAccess")
	{
		TeacherAccess.GET("getTeacherAccessList", v1.GetTeacherAccessList) // 获取教师列表
		TeacherAccess.POST("createTeacherAccess", v1.CreateTeacherAccess)  // 创建教师通行记录

	}
}
