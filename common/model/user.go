package model

type UserModel struct {
	Model
	User
}

func (UserModel) TableName() string {
	return "user"
}

type User struct {
	Model
	UserName string `gorm:"not null;unique;comment:用户名" json:"userName"`    // 用户名
	Avatar   string `gorm:"comment:头像" json:"avatar"`                       // 头像
	Password string `gorm:"default:'123456';comment:密码" json:"passWord"`    // 密码
	Mobile   string `gorm:"not null;unique;comment:手机号" json:"mobile"`      // 手机号
	Creator  uint64 `gorm:"not null;comment:创建者" json:"creator"`            // 创建者
	Editor   uint64 `gorm:"not null;comment:编辑者" json:"editor"`             // 编辑者
	Deleted  uint64 `gorm:"not null;default:0;comment:是否删除" json:"deleted"` // 是否删除
}
