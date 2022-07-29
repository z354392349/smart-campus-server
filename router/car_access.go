package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCarAccessRouter(Router *gin.RouterGroup) {
	CarAccess := Router.Group("carAccess")
	{
		CarAccess.GET("getCarAccessList", v1.GetCarAccessList) // 获取班级列表
		CarAccess.POST("createCarAccess", v1.CreateCarAccess)  // 创建车辆通行记录
		// CarAccess.PUT("upClass", v1.UpClass)                 // 更新班级
		// CarAccess.PUT("setClassMonitor", v1.SetClassMonitor) // 设置班长
		// CarAccess.PUT("setClassTeacher", v1.SetClassTeacher) // 设置班主任
		// CarAccess.DELETE("deleteClass", v1.DeleteClass)      // 删除班级
	}
}
