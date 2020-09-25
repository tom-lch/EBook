package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Ebook struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" gorm:"comment:书籍UUID"`
	Bookname    string    `json:"bookname"`
	Author      string    `json:"author"`
	Uploader    string    `json:"uploader"`
	Publisher   string    `json:"publisher"`
	PublishDate string    `json:"publishDate"`
	Version     string    `json:"version"`
	Cover       string    `json:"cover"`
	BookAllname string    `json:"bookAllname"`
	StoreAddr   string    `json:"storeAddr"`
}

type UserUpBook struct {
	gorm.Model
	Username string `json:"username"` // 关联到Ebook中的Uploader
	Up       int    `json:"up"`
	Bookname string `json:"bookname"`
}

type UserDownBook struct {
	gorm.Model
	Username string `json:"username"`
	Down     int    `json:"down"`
	Bookname string `json:"bookname"`
}
