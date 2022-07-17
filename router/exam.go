package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamRouter(Router *gin.RouterGroup) {
	Exam := Router.Group("exam")
	{
		Exam.GET("getExamList", v1.GetExamList)  // 获取考试列表
		Exam.POST("createExam", v1.CreateExam)   // 新增考试
		Exam.PUT("upExam", v1.UpExam)            // 更新考试
		Exam.DELETE("deleteExam", v1.DeleteExam) // 删除考试
	}
}
