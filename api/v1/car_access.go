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
// @Function:CreateCarAccess
// @Description:创建通行记录
// @Router:/carAccess/createCarAccess
// @Date:2022/07/29 10:49:26

func CreateCarAccess(c *gin.Context) {
	var carAccess model.CarAccess
	_ = c.ShouldBindJSON(&carAccess)

	if err := service.CreateCarAccess(carAccess); err != nil {
		global.GVA_LOG.Error("车辆通行数据创建失败!", zap.Any("err", err))
		response.FailWithMessage("车辆通行数据创建失败!", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Author: 张佳伟
// @Function:GetCarAccessList
// @Description:获取车辆通行列表
// @Router:/carAccess/getCarAccessList
// @Date:2022/07/29 17:35:43

func GetCarAccessList(c *gin.Context) {
	var pageInfo request.SearchCarAccess
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCarAccessList(pageInfo); err != nil {
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
