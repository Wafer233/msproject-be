package login

import (
	"fmt"
	"github.com/Wafer233/msproject-be/project-api/router"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init login router")
	rl := &RouterLogin{}
	router.RegisterRouter(rl)
}

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	rpc.InitUserRpc()
	h := New()
	r.POST("/project/login/getCaptcha", h.GetCaptcha)
}
