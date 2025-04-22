package dto

import (
	"errors"
	"regexp"
)

// RegisterRequest 注册请求DTO
type RegisterRequest struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

// RegisterResponse 注册响应DTO
type RegisterResponse struct {
	// 注册成功后可能返回一些数据，如用户ID等
}

// Verify 验证注册参数
func (r RegisterRequest) Verify() error {
	if !verifyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !verifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if r.Password != r.Password2 {
		return errors.New("两次密码输入不一致")
	}
	return nil
}

// verifyEmailFormat 验证邮箱格式
func verifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// verifyMobile 验证手机号
func verifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
