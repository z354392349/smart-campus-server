package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamItemRouter(Router *gin.RouterGroup) {
	ExamItem := Router.Group("examItem")
	{

		ExamItem.PUT("upExamItemRoomAllot", v1.UpExamItemRoomAllot) // 设置班主任

		// ExamItem.GET("getExamItemList", v1.GetExamItemList) // 获取考场列表
		// ExamItem.POST("createExamItem", v1.CreateExamItem)        // 新增考场信息
		// ExamItem.PUT("upExamItem", v1.UpExamItem)                 // 更新考场
		// ExamItem.DELETE("deleteExamItem", v1.DeleteExamItem)      // 删除考场

		// ExamRoom.DELETE("deleteExamRoom", v1.DeleteExamRoom) // 删除考场

	}
}
