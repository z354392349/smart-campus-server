package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamResultResultRouter(Router *gin.RouterGroup) {
	ExamResult := Router.Group("examResult")
	{
		ExamResult.GET("getExamResultList", v1.GetExamResultList) // 获取考试成绩列表
		// ExamResult.POST("createExamResult", v1.CreateExamResult)   // 新增考试
		// ExamResult.PUT("upExamResult", v1.UpExamResult)            // 更新考试成绩
		// ExamResult.DELETE("deleteExamResult", v1.DeleteExamResult) // 删除考试
	}
}
