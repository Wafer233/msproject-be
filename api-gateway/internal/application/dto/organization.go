package dto

type OrganizationDTO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	OwnerCode   string `json:"owner_code"`
	CreateTime  string `json:"create_time"`
	Personal    int    `json:"personal"`
	Code        string `json:"code"`
	Address     string `json:"address"`
	Province    int    `json:"province"`
	City        int    `json:"city"`
	Area        int    `json:"area"`
}

type OrganizationListResponse struct {
	List []OrganizationDTO `json:"list"`
}
