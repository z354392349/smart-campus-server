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

func CreateStudentAccess(c *gin.Context) {
	var studentAccess model.StudentAccess
	_ = c.ShouldBindJSON(&studentAccess)

	if err := service.CreateStudentAccess(studentAccess); err != nil {
		global.GVA_LOG.Error("学生通行数据创建失败!", zap.Any("err", err))
		response.FailWithMessage("学生通行数据创建失败!", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func CreateStudentAccessList(c *gin.Context) {
	var studentAccess []model.StudentAccess
	_ = c.ShouldBindJSON(&studentAccess)

	if err := service.CreateStudentAccessList(studentAccess); err != nil {
		global.GVA_LOG.Error("学生通行数据创建失败!", zap.Any("err", err))
		response.FailWithMessage("学生通行数据创建失败!", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func GetStudentAccessList(c *gin.Context) {
	var pageInfo request.SearchStudentAccess
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStudentAccessList(pageInfo); err != nil {
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
