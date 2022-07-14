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
// @Function: CreateCourse
// @Description: 获取课程列表
// @Router: /course/createCourse
// @Date:2022/07/14 22:02:34

func GetCourseList(c *gin.Context) {

	var pageInfo request.SearchCourseParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCourseList(pageInfo); err != nil {
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
// @Function:  CreateSubject
// @Description: 创建课程
// @Router: /course/createCourse
// @Date:2022/07/14 21:45:58

func CreateCourse(c *gin.Context) {

	var course model.Course
	_ = c.ShouldBindJSON(&course)

	if err := service.CreateCourse(course); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpCourse
// @Description: 更新课程
// @Router: /course/upCourse
// @Date:2022/07/14 22:45:30

func UpCourse(c *gin.Context) {

	var course model.Course
	_ = c.ShouldBindJSON(&course)

	if err := service.UpCourse(course); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function:DeleteCourse
// @Description:删除课程
// @Router:/course/deleteCourse
// @Date:2022/07/14 22:06:56

func DeleteCourse(c *gin.Context) {
	var course model.Course
	_ = c.ShouldBindJSON(&course)

	if err := service.DeleteCourse(course); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
