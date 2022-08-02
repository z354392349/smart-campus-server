package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTeacherAttendRouter(Router *gin.RouterGroup) {
	TeacherAttend := Router.Group("teacherAttend")
	{
		TeacherAttend.GET("getTeacherAttendList", v1.GetTeacherAttendList) // 获取教师考勤列表

	}
}
