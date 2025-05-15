package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

type OrganizationRouter struct {
	GetOrgListHandler *handler.GetOrgListHandler
}

func NewOrganizationRouter(organizationHandler *handler.GetOrgListHandler) *OrganizationRouter {
	return &OrganizationRouter{
		GetOrgListHandler: organizationHandler,
	}
}

func (router *OrganizationRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 受保护的路由（需要认证中间件）
	protected := group.Group("")
	protected.Use(middleware.TokenVerifyMiddleware())

	// 添加受保护的路由 - 确保路径正确
	protected.POST("/organization/_getOrgList", router.GetOrgListHandler.GetOrgList)
}
