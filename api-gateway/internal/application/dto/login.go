package dto

// LoginRequest 登录请求DTO
type LoginRequest struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

// LoginResponse 登录响应DTO
type LoginResponse struct {
	Member           MemberDTO         `json:"member"`
	TokenList        TokenDTO          `json:"tokenList"`
	OrganizationList []OrganizationDTO `json:"organizationList"`
}
