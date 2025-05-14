package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type OrganizationRouter struct {
	oh             *handler.OrganizationHandler
	authMiddleware gin.HandlerFunc
}

func NewOrganizationRouter(oh *handler.OrganizationHandler, authMiddleware gin.HandlerFunc) *OrganizationRouter {
	return &OrganizationRouter{
		oh:             oh,
		authMiddleware: authMiddleware,
	}
}

func (r *OrganizationRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 受保护的路由（需要认证中间件）
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// 添加受保护的路由
	protected.POST("/organization/_getOrgList", r.oh.GetOrgList)
}
