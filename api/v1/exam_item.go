package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetExamRoomList1(c *gin.Context) {

	var pageInfo request.SearchExamRoomParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExamRoomList(pageInfo); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:UpExamItemRoomAllot
// @Description:考试项，考场分配
// @Router:	/examItem/upExamItemRoomAllot
// @Date:2022/07/25 16:47:35

func UpExamItemRoomAllot(c *gin.Context) {

	var examRoom request.SetExamRoomItemAllot
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.UpExamItemRoomAllot(examRoom); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func DeleteExamRoom1(c *gin.Context) {
	var examRoom model.ExamRoom
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.DeleteExamRoom(examRoom); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func SetExamRoomTeacher1(c *gin.Context) {

	var info request.SetExamRoomTeacher
	_ = c.ShouldBindJSON(&info)

	if err := service.SetExamRoomTeacher(info); err != nil {
		global.GVA_LOG.Error("设置监考老师失败!", zap.Any("err", err))
		response.FailWithMessage("设置监考老师失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}

}
