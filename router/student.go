package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitStudentRouter(Router *gin.RouterGroup) {
	Student := Router.Group("student")
	{

		// todo: 缺一个 批量修改

		Student.GET("/getStudentList", v1.GetStudentList)                     // 获取学生列表
		Student.POST("/createStudent", v1.CreateStudent)                      // 创建学生
		Student.PUT("/upStudent", v1.UpStudent)                               // 更新学生
		Student.DELETE("/deleteStudent", v1.DeleteStudent)                    // 删除学生
		Student.PUT("/setStudentsGradeAndClass", v1.SetStudentsGradeAndClass) // 设置学生年级
	}
}
