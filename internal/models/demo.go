package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string         `json:"name" gorm:"column:name;type:varchar(100);unique_index;not null;comment:'姓名'"`
	Birthday time.Time      `json:"birthday" gorm:"column:birthday;type:datetime;not null;default:current_timestamp;comment:'生日'"`
	Gender   bool           `json:"gender" gorm:"column:gender;type:bool;not null;comment:'性别'"`
	Tags     datatypes.JSON `json:"tags" gorm:"column:tags;type:json;comment:'标签'"`
}

func (u User) TableName() string {
	return "test_auto_create_table"
}

type Login struct {
	ID       int    `json:"id" gorm:"column:id"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (l Login) TableName() string {
	return "test_user"
}
