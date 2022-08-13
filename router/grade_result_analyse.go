package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGradeResultAnalyseRouter(Router *gin.RouterGroup) {

	GradeResultAnalyse := Router.Group("gradeResultAnalyse")
	{

		GradeResultAnalyse.GET("getGradeAverageResult", v1.GetGradeAverageResult)                           // 获取平均学习成绩
		GradeResultAnalyse.GET("getGradePassPercent", v1.GetGradePassPercent)                               // 获取年级考试通过率
		GradeResultAnalyse.GET("getGradeAverageResultHistory", v1.GetGradeAverageResultHistory)             // 获取平均学习成绩 - 历史
		GradeResultAnalyse.GET("getGradeCourseAverageResultHistory", v1.GetGradeCourseAverageResultHistory) // 获取年级科目平均成绩 - 历史

	}
}
