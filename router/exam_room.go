package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitExamRoomRouter(Router *gin.RouterGroup) {
	ExamRoom := Router.Group("examRoom")
	{
		ExamRoom.GET("getExamRoomList", v1.GetExamRoomList)  // 获取年级列表
		ExamRoom.POST("createExamRoom", v1.CreateExamRoom)   // 新增年级信息
		ExamRoom.PUT("upExamRoom", v1.UpExamRoom)            // 更新年级
		ExamRoom.DELETE("deleteExamRoom", v1.DeleteExamRoom) // 删除年级
	}
}
