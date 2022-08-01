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
// @Function:CreateTeacherAccess
// @Description:创建教师通行记录
// @Router:/teacherAccess/createTeacherAccess
// @Date:2022/08/01 10:45:48

func CreateTeacherAccess(c *gin.Context) {
	var teacherAccess model.TeacherAccess
	_ = c.ShouldBindJSON(&teacherAccess)

	if err := service.CreateTeacherAccess(teacherAccess); err != nil {
		global.GVA_LOG.Error("教师通行数据创建失败!", zap.Any("err", err))
		response.FailWithMessage("教师通行数据创建失败!", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetTeacherAccessList
// @Description:获取教师通行记录
// @Router:/teacherAccess/getTeacherAccessList
// @Date:2022/08/01 10:45:48

func GetTeacherAccessList(c *gin.Context) {
	var pageInfo request.SearchTeacherAccess
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTeacherAccessList(pageInfo); err != nil {
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
