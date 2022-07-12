package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitClassRouter(Router *gin.RouterGroup) {
	Class := Router.Group("class")
	{
		Class.GET("getClassList", v1.GetClassList) // 获取班级列表
		Class.POST("createClass", v1.CreateClass)  // 获取班级列表
		//Class.GET("createClass", v1.GetGradeList)
		// Grade.GET("getGradeList", v1.GetGradeList)  // 获取年级列表
		// Grade.POST("createGrade", v1.CreateGrade)   // 新增年级信息
		// Grade.PUT("upGrade", v1.UpGrade)            // 更新年级
		// Grade.DELETE("deleteGrade", v1.DeleteGrade) // 删除年级
	}
}
