package entity

type Organization struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:name;type:varchar(255)"`
	Avatar      string `gorm:"column:avatar;type:varchar(511)"`
	Description string `gorm:"column:description;type:varchar(500)"`
	MemberID    int64  `gorm:"column:member_id;type:bigint"`
	CreateTime  int64  `gorm:"column:create_time;type:bigint"`
	Personal    int8   `gorm:"column:personal;type:tinyint(1);default:0"`
	Address     string `gorm:"column:address;type:varchar(100)"`
	Province    int    `gorm:"column:province;type:int;default:0"`
	City        int    `gorm:"column:city;type:int;default:0"`
	Area        int    `gorm:"column:area;type:int;default:0"`
}

// 表名定义
func (Organization) TableName() string {
	return "ms_organization"
}
