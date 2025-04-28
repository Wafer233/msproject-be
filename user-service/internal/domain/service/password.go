package service

import (
	"crypto/md5"
	"encoding/hex"
)

type PasswordService struct {
}

// NewPasswordService 创建密码服务
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (s *PasswordService) EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
