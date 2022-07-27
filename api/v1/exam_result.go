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
// @Function:GetExamResultList
// @Description:获取学生成绩列表
// @Router:/examResult/GetExamResultList
// @Date:2022/07/27 17:58:39

func GetExamResultList(c *gin.Context) {
	var pageInfo request.SearchExamResultParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExamResultList(pageInfo); err != nil {
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

func CreateExam1(c *gin.Context) {

	var exam model.Exam
	_ = c.ShouldBindJSON(&exam)
	if err := service.CreateExam(exam); err != nil {
		fmt.Println(err)
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

func UpExam1(c *gin.Context) {

	var exam model.Exam
	_ = c.ShouldBindJSON(&exam)

	if err := service.UpExam(exam); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func DeleteExam1(c *gin.Context) {
	var exam model.Exam
	_ = c.ShouldBindJSON(&exam)

	if err := service.DeleteExam(exam); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
