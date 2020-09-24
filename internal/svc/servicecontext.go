package svc

import (
	"EBook/internal/config"
	"EBook/internal/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 在此处添加gorm连接 在配置文件中添加连接数据库的drive和source
	DB, err := gorm.Open(mysql.Open(c.DBsource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ebook_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,     // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		logx.Error("connect mysql failed, err:", err)
	}
	// 在返回前在在数据库中建表，即迁移模型到数据库中建表
	DB.AutoMigrate(&model.User{})
	return &ServiceContext{Config: c, DB: DB}
}
