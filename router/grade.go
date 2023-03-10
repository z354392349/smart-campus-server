package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGradeRouter(Router *gin.RouterGroup) {
	Grade := Router.Group("grade")
	{
		Grade.GET("getGradeList", v1.GetGradeList)  // 获取年级列表
		Grade.POST("createGrade", v1.CreateGrade)   // 新增年级信息
		Grade.PUT("upGrade", v1.UpGrade)            // 更新年级
		Grade.DELETE("deleteGrade", v1.DeleteGrade) // 删除年级
	}
}
