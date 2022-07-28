package v1

import (
	"fmt"
	"gin-vue-admin/global"
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

// @Author: 张佳伟
// @Function:UpExamResult
// @Description:更新学生成绩
// @Router:/examResult/upExamResult/
// @Date:2022/07/28 17:37:05

func UpExamResult(c *gin.Context) {

	var exam request.SetExamResultParams
	_ = c.ShouldBindJSON(&exam)

	if err := service.UpExamResult(exam); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
