package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTeacherRouter(Router *gin.RouterGroup) {
	Teacher := Router.Group("teacher")
	{

		Teacher.POST("createTeacher", v1.CreateTeacher)   // 新增教师
		Teacher.GET("getTeacherList", v1.GetTeacherList)  // 获取教师列表
		Teacher.PUT("upTeacher", v1.UpTeacher)            // 更新教师
		Teacher.DELETE("deleteTeacher", v1.DeleteTeacher) // 删除教师

	}
}
