package service

type TokenService interface {
	GenerateTokens(userId string) (accessToken string, refreshToken string, accessExp int64)
	ValidateToken(tokenString string) (userId string, err error)
}
