package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username  string    `json:"userName" gorm:"comment:用户登录名;unique"`
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称;unique"`
	HeaderImg string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Email     string    `json:"email" gorm:"comment:用户有效;unique"`
	Phone     string    `json:"phone" gomr:"comment:"用户手机号;unique" `
}
