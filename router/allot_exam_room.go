package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAllotExamRoomRouter(Router *gin.RouterGroup) {
	AllotExamRoom := Router.Group("allotExamRoom")
	{

		AllotExamRoom.GET("/getAllotExamRoomList", v1.GetAllotExamRoomList) // 获取考场分配结果列表

	}
}
