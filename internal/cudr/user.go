package cudr

import (
	"EBook/internal/model"
	"EBook/pkg/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

func Register(u model.User, DB *gorm.DB) error {
	//在数据库中插入数据
	// 先查询用户名是否注册
	//user := &model.User{}
	//DB.Where("username=?", u.Username).First(user)
	//if user.Username == u.Username {
	//	return  errors.New("用户名已存在")
	//}
	u.Password = utils.SHA256(u.Password)
	u.UUID = uuid.NewV4()
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func Login(u *model.User, DB *gorm.DB) (*model.User, error) {
	user := &model.User{}
	u.Password = utils.SHA256(u.Password)
	if err := DB.Where("username=? AND password=?", u.Username, u.Password).First(user).Error; err != nil {
		return nil, errors.New("用户名或密码出错")
	}
	return user, nil
}

func Delete(u *model.User, DB *gorm.DB) error {
	user := &model.User{}
	u.Password = utils.SHA256(u.Password)
	if err := DB.Where("username=? AND password=?", u.Username, u.Password).Delete(user).Error; err != nil {
		logx.Error("删除失败")
		return err
	}
	return nil
}

func GetUserInfo(u *model.User, DB *gorm.DB) (*model.User, error) {
	user := &model.User{}
	if err := DB.Where("username=?", u.Username).
		Select("username", "nick_name", "email").
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func UpdateSecert(u *model.User, newp string, DB *gorm.DB) error {
	user := &model.User{Username: u.Username, Password: utils.SHA256(u.Password)}
	if err := DB.Model(user).Where("username = ? AND password=?", u.Username, utils.SHA256(u.Password)).Updates(map[string]interface{}{"password": utils.SHA256(newp)}).Error; err != nil {
		return err
	}
	return nil
}