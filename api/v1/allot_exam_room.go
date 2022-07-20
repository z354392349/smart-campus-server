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

// @Author: 张佳伟
// @Function: 获取考场分配列表
// @Description: GetAllotExamRoomList
// @Router:/allotExamRoom/getAllotExamRoomList
// @Date: 2022/7/20 15:09

func GetAllotExamRoomList(c *gin.Context) {
	var pageInfo request.SearchAllotExamRoomParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAllotExamRoomList(pageInfo); err != nil {
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
// @Function:CreateAllotExamRoom
// @Description:创建考场分配列表
// @Router:/allotExamRoom/createAllotExamRoom
// @Date: 2022/7/20 15:11

func CreateAllotExamRoom(c *gin.Context) {

	var exam request.AllotExamRoom
	_ = c.ShouldBindJSON(&exam)

	var exams []model.AllotExamRoom

	if err := service.CreateAllotExamRoom(exams); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpAllotExamRoom
// @Description: 批量更新考场分配信息
// @Router: /allotExamRoom/upAllotExamRoom
// @Date: 2022/7/20 15:11

func UpAllotExamRoom(c *gin.Context) {
	var exam request.AllotExamRoom
	_ = c.ShouldBindJSON(&exam)
	var exams []model.AllotExamRoom

	if err := service.UpAllotExamRoom(exams); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
