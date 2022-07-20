package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAllotExamRoomRouter(Router *gin.RouterGroup) {
	AllotExamRoom := Router.Group("allotExamRoom")
	{

		AllotExamRoom.GET("/getAllotExamRoomList", v1.GetAllotExamRoomList) // 获取考场列表
		AllotExamRoom.POST("/createAllotExamRoom", v1.CreateAllotExamRoom)  // 创建考场
		AllotExamRoom.PUT("/upAllotExamRoom", v1.UpAllotExamRoom)           // 更新考场

		//AllotExamRoom.DELETE("/deleteAllotExamRoom", v1.DeleteAllotExamRoom) // 暂不考虑删除考场
	}
}
