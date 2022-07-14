package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitClassRouter(Router *gin.RouterGroup) {
	Class := Router.Group("class")
	{
		Class.GET("getClassList", v1.GetClassList)  // 获取班级列表
		Class.POST("createClass", v1.CreateClass)   // 新增班级信息
		Class.PUT("upClass", v1.UpClass)            // 更新班级
		Class.DELETE("deleteClass", v1.DeleteClass) // 删除班级
	}
}
