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
// @Function:GetGradeAverageResult
// @Description: 获取平均学习成绩
// @Router:/gradeResultAnalyse/getGradeAverageResult
// @Date:2022/08/08 21:02:59

func GetGradeAverageResult(c *gin.Context) {
	var info request.GradeResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetGradeAverageResult(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetGradePassPercent
// @Description:获取年级考试通过率
// @Router:/gradeResultAnalyse/GetGradePassPercent
// @Date:2022/08/12 15:40:40

func GetGradePassPercent(c *gin.Context) {
	var info request.GradeResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, percent := service.GetGradePassPercent(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(percent, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetGradeAverageResultHistory
// @Description: 获取年级
// @Router:/gradeResultAnalyse/GetGradeAverageResultHistory
// @Date:2022/08/13 17:13:08

func GetGradeAverageResultHistory(c *gin.Context) {
	var info request.GradeResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, percent := service.GetGradeAverageResultHistory(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(percent, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetGradeCourseAverageResultHistory
// @Description: 获取年级科目平均成绩
// @Router:/gradeResultAnalyse/GetGradeCourseAverageResultHistory
// @Date:2022/08/13 17:13:08

func GetGradeCourseAverageResultHistory(c *gin.Context) {
	var info request.GradeResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, percent := service.GetGradeCourseAverageResultHistory(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(percent, "获取成功", c)
	}
}
