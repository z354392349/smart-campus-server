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

// @Author: 张佳伟
// @Function: GetTeacherAttendList
// @Description: 获取教师通勤列表
// @Router:/teacherAttend/getTeacherAttendList
// @Date:2022/08/02 14:38:36

func GetTeacherAttendList(c *gin.Context) {
	var pageInfo request.SearchTeacherAttend
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTeacherAttendList(pageInfo); err != nil {
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
