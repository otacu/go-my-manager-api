package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	Name     string
	Avatar   string
	Roles    string
}

// 设置User的表名为`tb_user`
func (User) TableName() string {
	return "tb_user"
}
