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

func GetDashboardCensusNum(c *gin.Context) {
	var pageInfo request.SearchCourseParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, carAccess, peopleAccess, teacherCensus, studentCensus := service.GetDashboardCensusNum(); err != nil {
		fmt.Println(err)
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
