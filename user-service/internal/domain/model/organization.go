package model

// Organization 组织实体
type Organization struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
	MemberId    int64 // 所属用户
	CreateTime  int64
	Personal    int32 // 是否个人项目 1: 是, 0: 否
	Address     string
	Province    int32
	City        int32
	Area        int32
}
