package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamItemRouter(Router *gin.RouterGroup) {
	ExamItem := Router.Group("examItem")
	{

		ExamItem.PUT("allotExamItemRoom", v1.AllotExamItemRoom)             // 分配考场
		ExamItem.PUT("cancelAllotExamItemRoom", v1.CancelAllotExamItemRoom) // 撤销分配的考场

		// ExamItem.GET("getExamItemList", v1.GetExamItemList) // 获取考场列表
		// ExamItem.POST("createExamItem", v1.CreateExamItem)  // 新增考场信息
		// ExamItem.DELETE("deleteExamItem", v1.DeleteExamItem)// 删除考场
		// ExamRoom.DELETE("deleteExamRoom", v1.DeleteExamRoom) // 删除考场

	}
}
