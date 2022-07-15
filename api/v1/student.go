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
// @Function: GetStudentList1
// @Description: 创建学生
// @Router:/student/getStudentList
// @Date: 2022/7/15 11:51:12

func GetStudentList(c *gin.Context) {
	var pageInfo request.SearchStudentParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStudentList(pageInfo); err != nil {
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
// @Function: CreateStudent
// @Description: 创建学生
// @Router:/student/createStudent
// @Date: 2022/7/15 10:40:12

func CreateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)

	if err := service.CreateStudent(student); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpStudent
// @Description: 更新学生
// @Router:/student/upStudent
// @Date: 2022/7/15 10:52

func UpStudent(c *gin.Context) {

	var student model.Student
	_ = c.ShouldBindJSON(&student)

	if err := service.UpStudent(student); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function: DeleteStudent
// @Description:删除学生
// @Router:/student/deleteStudent
// @Date: 2022/7/15 10:52

func DeleteStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)

	if err := service.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
