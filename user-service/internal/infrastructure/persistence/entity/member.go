package entity

type Member struct {
	ID              int64  `gorm:"primaryKey;column:id;autoIncrement;comment:系统前台用户表"`
	Account         string `gorm:"type:varchar(20);not null;default:'';comment:用户登陆账号"`
	Password        string `gorm:"type:varchar(64);default:'';comment:登陆密码"`
	Name            string `gorm:"type:varchar(255);default:'';comment:用户昵称"`
	Mobile          string `gorm:"type:varchar(255);comment:手机"`
	Realname        string `gorm:"type:varchar(255);comment:真实姓名"`
	CreateTime      int64  `gorm:"column:create_time;comment:创建时间"` // 原字段为 varchar，但你用 int64 表示时间戳，这里按你设定保留 int64
	Status          int    `gorm:"type:tinyint(1);default:0;comment:状态"`
	LastLoginTime   int64  `gorm:"column:last_login_time;comment:上次登录时间"` // 同上，保留 int64 表示时间戳
	Sex             int    `gorm:"type:tinyint(1);default:0;comment:性别"`
	Avatar          string `gorm:"type:varchar(255);default:'';comment:头像"`
	Idcard          string `gorm:"type:varchar(255);comment:身份证"`
	Province        int    `gorm:"type:int;default:0;comment:省"`
	City            int    `gorm:"type:int;default:0;comment:市"`
	Area            int    `gorm:"type:int;default:0;comment:区"`
	Address         string `gorm:"type:varchar(255);comment:所在地址"`
	Description     string `gorm:"type:text;comment:备注"`
	Email           string `gorm:"type:varchar(255);comment:邮箱"`
	DingtalkOpenid  string `gorm:"type:varchar(50);column:dingtalk_openid;comment:钉钉openid"`
	DingtalkUnionid string `gorm:"type:varchar(50);column:dingtalk_unionid;comment:钉钉unionid"`
	DingtalkUserid  string `gorm:"type:varchar(50);column:dingtalk_userid;comment:钉钉用户id"`
}

// 如果需要手动指定表名
func (Member) TableName() string {
	return "ms_member"
}
