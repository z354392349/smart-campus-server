package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Author: 张佳伟
// @Function:GetDashboardCensusNum
// @Description: 获取通行数量，考勤
// @Router:/dashboard/getDashboardCensusNum
// @Date:2022/08/08 21:02:59

func GetDashboardCensusNum(c *gin.Context) {

	if err, carAccess, peopleAccess, teacherCensus, studentCensus := service.GetDashboardCensusNum(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.GetDashboardCensusNum{
			CarAccess:     carAccess,
			PeopleAccess:  peopleAccess,
			TeacherCensus: teacherCensus,
			StudentCensus: studentCensus,
		}, "获取成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetTeacherNum
// @Description:获取教师数量区分男女
// @Router:/dashboard/getTeacherNum
// @Date:2022/08/08 21:20:46

func GetTeacherNum(c *gin.Context) {

	if err, Teacherlist := service.GetTeacherNum(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(Teacherlist, "获取成功", c)
	}
}