package dto

// OrganizationDTO 表示API响应中的组织实体
type OrganizationDTO struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Avatar      interface{} `json:"avatar"` // 使用interface{}以便处理null值
	Description interface{} `json:"description"`
	OwnerCode   string      `json:"owner_code"`
	CreateTime  string      `json:"create_time"`
	Personal    int         `json:"personal"`
	Code        string      `json:"code"`
	Address     interface{} `json:"address"`
	Province    int         `json:"province"`
	City        int         `json:"city"`
	Area        int         `json:"area"`
}

type OrganizationListResponse struct {
	List []OrganizationDTO `json:"list"`
}
