package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitStudentAttendRouter(Router *gin.RouterGroup) {
	StudentAttend := Router.Group("studentAttend")
	{
		StudentAttend.GET("getStudentAttendList", v1.GetStudentAttendList) // 获取学生考勤列表
	}
}
