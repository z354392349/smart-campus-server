package v1

import (
	"fmt"
	"gin-vue-admin/global"
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
// @Function:allotExamItemRoom
// @Description:考试项，考场分配
// @Router:	/examItem/allotExamItemRoom
// @Date:2022/07/25 16:47:35

func AllotExamItemRoom(c *gin.Context) {

	var examRoom request.AllotExamRoomItem
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.AllotExamItemRoom(examRoom); err != nil {
		global.GVA_LOG.Error("分配失败!", zap.Any("err", err))
		response.FailWithMessage("分配失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function:CancelAllotExamItemRoom
// @Description:撤销已分配的考场
// @Router: /examItem/cancelAllotExamItemRoom
// @Date:2022/07/26 16:03:21

func CancelAllotExamItemRoom(c *gin.Context) {
	var info request.CancelAllotExamRoomItem
	_ = c.ShouldBindJSON(&info)

	if err := service.CancelAllotExamItemRoom(info); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
