package entity

type Organization struct {
	ID          int64  `gorm:"primaryKey;column:id;autoIncrement;comment:主键"`
	Name        string `gorm:"type:varchar(255);comment:名称"`
	Avatar      string `gorm:"type:varchar(511);comment:头像"`
	Description string `gorm:"type:varchar(500);comment:描述"`
	MemberID    int64  `gorm:"column:member_id;comment:拥有者"`    // 外键字段
	CreateTime  int64  `gorm:"column:create_time;comment:创建时间"` // 保留时间戳类型
	Personal    int32  `gorm:"type:tinyint(1);default:0;comment:是否个人项目"`
	Address     string `gorm:"type:varchar(100);comment:地址"`
	Province    int32  `gorm:"type:int;default:0;comment:省"`
	City        int32  `gorm:"type:int;default:0;comment:市"`
	Area        int32  `gorm:"type:int;default:0;comment:区"`
}

// 表名定义
func (Organization) TableName() string {
	return "ms_organization"
}
