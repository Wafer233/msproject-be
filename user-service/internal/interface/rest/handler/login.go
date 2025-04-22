package handler

import (
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
)

// LoginHandler 处理登录相关请求
type LoginHandler struct {
	as service.AuthService
}

func NewLoginHandler(as service.AuthService) *LoginHandler {
	return &LoginHandler{
		as: nil,
	}
}
