package https_server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "go-chat/api/v1"
)

var GE *gin.Engine

func init() {
	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	GE.Use(cors.New(corsConfig))
	// GE.Use(ssl.TlsHandler(config.GetConfig().MainConfig.Host, config.GetConfig().MainConfig.Port))
	GE.POST("/register", v1.Register)
	GE.POST("/login", v1.Login)
	GE.POST("/user/updateUserInfo", v1.UpdateUserInfo)
	GE.POST("/user/getUserInfoList", v1.GetUserInfoList)
	GE.POST("/user/ableUsers", v1.AbleUsers)
	GE.POST("/user/getUserInfo", v1.GetUserInfo)
	GE.POST("/user/disableUsers", v1.DisableUsers)
	GE.POST("/user/deleteUsers", v1.DeleteUsers)
	GE.POST("/user/setAdmin", v1.SetAdmin)
	GE.POST("/user/sendSmsCode", v1.SendSmsCode)
	GE.POST("/user/smsLogin", v1.SmsLogin)

	GE.POST("/group/createGroup", v1.CreateGroup)
	GE.POST("/group/loadMyGroup", v1.LoadMyGroup)
	GE.POST("/group/checkGroupAddMode", v1.CheckGroupAddMode)
	GE.POST("/group/enterGroupDirectly", v1.EnterGroupDirectly)
	GE.POST("/group/leaveGroup", v1.LeaveGroup)
	GE.POST("/group/dismissGroup", v1.DismissGroup)
	GE.POST("/group/getGroupInfo", v1.GetGroupInfo)
	GE.POST("/group/getGroupInfoList", v1.GetGroupInfoList)
	GE.POST("/group/deleteGroups", v1.DeleteGroups)
	GE.POST("/group/setGroupsStatus", v1.SetGroupsStatus)
	GE.POST("/group/updateGroupInfo", v1.UpdateGroupInfo)
	GE.POST("/group/getGroupMemberList", v1.GetGroupMemberList)
	GE.POST("/group/removeGroupMembers", v1.RemoveGroupMembers)

}
