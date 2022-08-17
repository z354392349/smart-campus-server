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
// @Function:GetClassTotalResult
// @Description: 获取班级下每个学生总成绩
// @Router:/classResultAnalyse/getClassTotalResult
// @Date:2022/08/14 18:05:44

func GetClassTotalResult(c *gin.Context) {
	var info request.ClassResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetClassTotalResult(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetClassPassPercent
// @Description: 获取指定班级的通过率
// @Router:/classResultAnalyse/getClassPassPercent
// @Date:2022/08/15 09:16:41

func GetClassPassPercent(c *gin.Context) {
	var info request.ClassResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetClassPassPercent(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:getClassToTalResultHistory
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/classResultAnalyse/getClassToTalResultHistory
// @Date:2022/08/15 09:52:31

func GetClassToTalResultHistory(c *gin.Context) {
	var info request.ClassResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetClassToTalResultHistory(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetClassCourseResultHistory1
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/classResultAnalyse/getClassCourseResultHistory1
// @Date:2022/08/15 09:52:31

func GetClassCourseResultHistory(c *gin.Context) {
	var info request.ClassResultAnalyse
	_ = c.ShouldBindQuery(&info)
	if err, carAccess := service.GetClassCourseResultHistory(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(carAccess, "获取成功", c)
	}
}
