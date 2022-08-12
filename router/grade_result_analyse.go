package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGradeResultAnalyseRouter(Router *gin.RouterGroup) {

	GradeResultAnalyse := Router.Group("gradeResultAnalyse")
	{

		GradeResultAnalyse.GET("getGradeAverageResult", v1.GetGradeAverageResult) // 获取平均学习成绩
		GradeResultAnalyse.GET("getGradePassPercent", v1.GetGradePassPercent)     // 获取年级考试通过率

	}
}
