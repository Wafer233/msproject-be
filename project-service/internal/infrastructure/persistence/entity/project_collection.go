package entity

type ProjectCollection struct {
	Id          int64 `gorm:"column:id;primaryKey;autoIncrement"`
	ProjectCode int64 `gorm:"column:project_code"`
	MemberCode  int64 `gorm:"column:member_code"`
	CreateTime  int64 `gorm:"column:create_time"`
}

func (*ProjectCollection) TableName() string {
	return "ms_project_collection"
}
