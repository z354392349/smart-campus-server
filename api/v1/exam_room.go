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
// @Function: GetExamRoomList
// @Description: 获取考场列表
// @Router:/examRoom/getExamRoomList
// @Date: 2022/7/15 15:17

func GetExamRoomList(c *gin.Context) {

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
// @Function:CreateExamRoom
// @Description:创建考场
// @Router:/examRoom/createExamRoom
// @Date: 2022/7/15 15:11

func CreateExamRoom(c *gin.Context) {

	var examRoom model.ExamRoom
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.CreateExamRoom(examRoom); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败，存在相同的名称", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpExamRoom
// @Description: 更新考场
// @Router: /examRoom/UpExamRoom
// @Date: 2022/7/15 15:23:12

func UpExamRoom(c *gin.Context) {

	var examRoom model.ExamRoom
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.UpExamRoom(examRoom); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function: DeleteExamRoom
// @Description:删除考场
// @Router: /examRoom/deleteExamRoom
// @Date: 2022/7/15 15:23:32

func DeleteExamRoom(c *gin.Context) {
	var examRoom model.ExamRoom
	_ = c.ShouldBindJSON(&examRoom)

	if err := service.DeleteExamRoom(examRoom); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Author: 张佳伟
// @Function: SetExamRoomTeacher
// @Description: 设置监考老师
// @Router: /examRoom/setExamRoomTeacher
// @Date:2022/07/25 10:07:56

func SetExamRoomTeacher(c *gin.Context) {

	var info request.SetExamRoomTeacher
	_ = c.ShouldBindJSON(&info)

	if err := service.SetExamRoomTeacher(info); err != nil {
		global.GVA_LOG.Error("设置监考老师失败!", zap.Any("err", err))
		response.FailWithMessage("设置监考老师失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}

}
