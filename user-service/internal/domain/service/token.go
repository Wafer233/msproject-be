package service

// TokenService 令牌服务接口(领域服务)
type TokenService interface {
	GenerateTokens(userId string) (accessToken string, refreshToken string, accessExp int64)
	ValidateToken(tokenString string) (userId string, err error)
}
