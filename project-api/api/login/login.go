package login

import (
	"fmt"
	common "github.com/Wafer233/msproject-be/project-common"
	"github.com/gin-gonic/gin"
)

type HandlerLogin struct {
}

func New() *HandlerLogin {
	return &HandlerLogin{}
}

func (*HandlerLogin) GetCaptcha(c *gin.Context) {
	result := &common.Result{}
	mobile := c.PostForm("mobile")
	ctx := context.Background()
	_, err := rpc.UserClient.GetCaptcha(ctx, &login_service_v1.CaptchaMessage{
		Mobile: mobile,
	})
	if err != nil {
		c.JSON(200, result.Fail(2001, err.Error()))
		return
	}
	c.JSON(200, result.Success(nil))
}
