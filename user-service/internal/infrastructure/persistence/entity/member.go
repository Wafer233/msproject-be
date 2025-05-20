package entity

type Member struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Account         string `gorm:"column:account;type:varchar(20);not null;default:''"`
	Password        string `gorm:"column:password;type:varchar(64);default:''"`
	Name            string `gorm:"column:name;type:varchar(255);default:''"`
	Mobile          string `gorm:"column:mobile;type:varchar(255)"`
	RealName        string `gorm:"column:realname;type:varchar(255)"`
	CreateTime      string `gorm:"column:create_time;type:varchar(30)"`
	Status          int8   `gorm:"column:status;type:tinyint(1);default:0"`
	LastLoginTime   string `gorm:"column:last_login_time;type:varchar(30)"`
	Sex             int8   `gorm:"column:sex;type:tinyint(1);default:0"`
	Avatar          string `gorm:"column:avatar;type:varchar(255);default:''"`
	IDCard          string `gorm:"column:idcard;type:varchar(255)"`
	Province        int    `gorm:"column:province;type:int;default:0"`
	City            int    `gorm:"column:city;type:int;default:0"`
	Area            int    `gorm:"column:area;type:int;default:0"`
	Address         string `gorm:"column:address;type:varchar(255)"`
	Description     string `gorm:"column:description;type:text"`
	Email           string `gorm:"column:email;type:varchar(255)"`
	DingTalkOpenID  string `gorm:"column:dingtalk_openid;type:varchar(50)"`
	DingTalkUnionID string `gorm:"column:dingtalk_unionid;type:varchar(50)"`
	DingTalkUserID  string `gorm:"column:dingtalk_userid;type:varchar(50)"`
}

// 如果需要手动指定表名
func (Member) TableName() string {
	return "ms_member"
}
