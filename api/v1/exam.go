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
// @Function:GetExamList
// @Description:获取考试列表
// @Router:/exam/getExamList
// @Date:2022/07/16 20:10:37

func GetExamList(c *gin.Context) {
	var pageInfo request.SearchExamParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExamList(pageInfo); err != nil {
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
// @Function:CreateExam
// @Description:创建考试
// @Router:/exam/createExam
// @Date:2022/07/16 20:06:51

func CreateExam(c *gin.Context) {

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

// @Author: 张佳伟
// @Function:UpExam
// @Description: 更新考试
// @Router:/exam/upExam
// @Date:2022/07/16 20:17:14

func UpExam(c *gin.Context) {

	var exam model.Exam
	_ = c.ShouldBindJSON(&exam)

	if err := service.UpExam(exam); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Author: 张佳伟
// @Function:DeleteExam
// @Description:删除考试
// @Router:/exam/deleteExam
// @Date:2022/07/16 20:17:47

func DeleteExam(c *gin.Context) {
	var exam model.Exam
	_ = c.ShouldBindJSON(&exam)

	if err := service.DeleteExam(exam); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
