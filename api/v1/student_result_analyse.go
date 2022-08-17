package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Author: 张佳伟
// @Function:GetStudentCourseResult
// @Description:获取学生每一个科目的成绩
// @Router:/studentResultAnalyse/getStudentCourseResult
// @Date:2022/08/16 17:44:32

func GetStudentCourseResult(c *gin.Context) {
	var info request.StudentResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetStudentCourseResult(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function: GetStudentTotalResultHistory
// @Description: 获取学生历史考试每一次的总成绩
// @Router:/studentResultAnalyse/getStudentTotalResultHistory
// @Date:2022/08/17 16:22:07

func GetStudentTotalResultHistory(c *gin.Context) {
	var info request.StudentResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetStudentTotalResultHistory(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}
