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

	// 创建用户账户
	user := &model.SysUser{Username: student.Telephone, NickName: student.Name, Password: "123456", AuthorityId: "03"}
	err, userReturn := service.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "账号注册失败", c)
		return
	}

	// 创建学生信息
	student.SysUserID = userReturn.ID
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

	// 更新用户账户
	user := &model.SysUser{Username: student.Telephone, NickName: student.Name}
	user.ID = student.SysUserID
	err, userReturn := service.UpUserInfoByID(*user)
	if err != nil {
		global.GVA_LOG.Error("账户信息修改失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "账户信息修改失败", c)
		return
	}

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

	var delStudent model.Student
	_ = global.GVA_DB.Where("`id` = ?", student.ID).Find(&delStudent).Error

	err := service.DeleteUser(float64(delStudent.SysUserID))
	if err != nil {
		global.GVA_LOG.Error("账户信息删除失败!", zap.Any("err", err))
		response.FailWithMessage("账户信息删除失败", c)
		return
	}

	if err := service.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
