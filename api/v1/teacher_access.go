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
