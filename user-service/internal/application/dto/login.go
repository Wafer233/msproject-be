package dto

// LoginRequest 登录请求DTO
type LoginRequest struct {
	Account  string `json:"account" `
	Password string `json:"password" `
}

// LoginResponse 登录响应DTO
type LoginResponse struct {
	Member           MemberDTO         `json:"member"`
	TokenList        TokenDTO          `json:"tokenList"`
	OrganizationList []OrganizationDTO `json:"organizationList"`
}

// MemberDTO 用户DTO
type MemberDTO struct {
	Id            int64  `json:"id"`
	Account       string `json:"account"`
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	Status        int    `json:"status"`
	LastLoginTime int64  `json:"last_login_time"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
}

// TokenDTO 令牌DTO
type TokenDTO struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenType      string `json:"tokenType"`
	AccessTokenExp int64  `json:"accessTokenExp"`
}

// OrganizationDTO 组织DTO
type OrganizationDTO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	MemberId    int64  `json:"memberId"`
	CreateTime  int64  `json:"createTime"`
	Personal    int32  `json:"personal"`
	Address     string `json:"address"`
	Province    int32  `json:"province"`
	City        int32  `json:"city"`
	Area        int32  `json:"area"`
}
