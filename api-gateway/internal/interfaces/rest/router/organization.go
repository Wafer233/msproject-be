package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type OrganizationRouter struct {
	GetOrgListHandler     *handler.GetOrgListHandler
	tokenVerifyMiddleware gin.HandlerFunc
}

func NewOrganizationRouter(
	organizationHandler *handler.GetOrgListHandler,
	tokenVerifyMiddleware gin.HandlerFunc,
) *OrganizationRouter {
	return &OrganizationRouter{
		GetOrgListHandler:     organizationHandler,
		tokenVerifyMiddleware: tokenVerifyMiddleware,
	}
}

func (router *OrganizationRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 受保护的路由（需要认证中间件）
	protected := group.Group("")
	protected.Use(router.tokenVerifyMiddleware)

	// 添加受保护的路由 - 确保路径正确
	protected.POST("/organization/_getOrgList", router.GetOrgListHandler.GetOrgList)
}
