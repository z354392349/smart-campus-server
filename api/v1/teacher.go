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
// @Function: GetTeacherList
// @Description: 获取教师列表
// @Router:/teacher/GetTeacherList
// @Date: 2022/7/11 11:54:25

func GetTeacherList(c *gin.Context) {
	var pageInfo request.SearchTeacherParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTeacherList(pageInfo); err != nil {
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
// @Function: CreateTeacher
// @Description: 创建教师
// @Router: /teacher/createTeacher
// @Date:2022/07/10 19:46:37

func CreateTeacher(c *gin.Context) {

	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)
	if err := service.CreateTeacher(teacher); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpTeacher
// @Description: 更新教师
// @Router: /teacher/upTeacher
// @Date: 2022/7/11 16:36:31

func UpTeacher(c *gin.Context) {

	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)

	if err := service.UpTeacher(teacher); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function:
// @Description:
// @Router: /grade/deleteGrade
// @Date:2022/07/09 10:51:45

func DeleteTeacher(c *gin.Context) {
	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)

	if err := service.DeleteTeacher(teacher); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
