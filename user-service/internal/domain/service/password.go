package service

import (
	"crypto/md5"
	"encoding/hex"
)

type PasswordService struct{}

// NewPasswordService 创建密码服务
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// EncryptPassword 加密密码
func (s *PasswordService) EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyPassword 验证密码
func (s *PasswordService) VerifyPassword(inputPassword, storedPassword string) bool {
	return s.EncryptPassword(inputPassword) == storedPassword
}
