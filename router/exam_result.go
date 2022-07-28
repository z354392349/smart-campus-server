package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamResultRouter(Router *gin.RouterGroup) {
	ExamResult := Router.Group("examResult")
	{
		ExamResult.GET("getExamResultList", v1.GetExamResultList) // 获取考试成绩列表
		ExamResult.PUT("upExamResult", v1.UpExamResult)           // 新增考试
	}
}
