package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)                 // 用户注册账号
		UserRouter.POST("changePassword", v1.ChangePassword)     // 用户修改密码
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
	}
}
