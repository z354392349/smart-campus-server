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
// @Function: GetClassList
// @Description: 获取班级列表
// @Router: /class/getClassList
// @Date:2022/07/10 13:47:29

func GetClassList(c *gin.Context) {

	var pageInfo request.SearchClassParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetClassList(pageInfo); err != nil {
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
// @Function: CreateClass
// @Description: 创建年级
// @Router: /class/createClass
// @Date: 2022/7/12 10:43:23

func CreateClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)

	if err := service.CreateClass(class); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Author: 张佳伟
// @Function: UpGrade
// @Description: 修改年级
// @Router: /grade/upGrade
// @Date:2022/07/09 10:25:34
func UpGrade1(c *gin.Context) {

	var grade model.Grade
	_ = c.ShouldBindJSON(&grade)

	if err := service.UpCreate(grade); err != nil {
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

func DeleteGrade1(c *gin.Context) {
	var grade model.Grade
	_ = c.ShouldBindJSON(&grade)

	if err := service.DeleteGrade(grade); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
