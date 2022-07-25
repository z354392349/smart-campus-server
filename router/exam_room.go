package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamRoomRouter(Router *gin.RouterGroup) {
	ExamRoom := Router.Group("examRoom")
	{
		ExamRoom.GET("getExamRoomList", v1.GetExamRoomList)       // 获取考场列表
		ExamRoom.POST("createExamRoom", v1.CreateExamRoom)        // 新增考场信息
		ExamRoom.PUT("upExamRoom", v1.UpExamRoom)                 // 更新考场
		ExamRoom.PUT("setExamRoomTeacher", v1.SetExamRoomTeacher) // 设置班主任
		ExamRoom.DELETE("deleteExamRoom", v1.DeleteExamRoom)      // 删除考场
	}
}
