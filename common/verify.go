package common

import (
	"errors"
	"regexp"
)

func Verify(email, mobile, password1, password2 string) error {
	if VerifyEmailFormat(email) {
		return errors.New("邮箱格式不正确")
	}
	if VerifyMobile(mobile) {
		return errors.New("手机号格式不正确")
	}
	if VerifyPassword(password1, password2) {
		return errors.New("两次密码输入不一致")
	}
	return nil
}

// VerifyMobile 验证手机合法性
func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyPassword(password1, password2 string) bool {
	return password1 == password2
}
