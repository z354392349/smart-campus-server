package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitStudentAccessRouter(Router *gin.RouterGroup) {
	StudentAccess := Router.Group("studentAccess")
	{
		StudentAccess.GET("getStudentAccessList", v1.GetStudentAccessList)        // 获取学生列表
		StudentAccess.POST("createStudentAccess", v1.CreateStudentAccess)         // 创建学生通行记录
		StudentAccess.POST("createStudentAccessList", v1.CreateStudentAccessList) // 创建学生通行记录

	}
}
