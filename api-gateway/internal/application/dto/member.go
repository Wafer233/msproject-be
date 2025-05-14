package dto

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
