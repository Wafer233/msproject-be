package model

// Member 用户实体
type Member struct {
	Id            int64
	Account       string
	Password      string
	Name          string
	Mobile        string
	Email         string
	Status        int // 1: 正常, 0: 禁用
	CreateTime    int64
	LastLoginTime int64
	Avatar        string
	Description   string
	Address       string
	Province      int
	City          int
	Area          int
}

// IsValid 检查用户状态是否有效
func (m *Member) IsValid() bool {
	return m.Status == 1
}
