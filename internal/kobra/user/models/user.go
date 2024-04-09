package models

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/tools"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"primary_key;comment:用户唯一标识符" validate:"required,uuid"`
	Email    string    `gorm:"unique;comment:用户ID" validate:"required,email"`
	Username string    `gorm:"unique;comment:用户名" validate:"required,min=5,max=32"`
	Password string    `gorm:"unique;comment:用户密码" validate:"required,min=8,max=128"`
	Name     string    `gorm:"comment:用户姓名" validate:"omitempty,min=0,max=32"`
	Age      uint8     `gorm:"comment:用户姓名" validate:"omitempty,min=0,max=100"`
}

func (user *User) BeforeCreate(*gorm.DB) (err error) {
	user.UUID = uuid.New()
	if err := tools.ValidateData(user); err != nil {
		return err
	}
	db := kobra.Env.DB()
	if err := db.Where(user.Email).Find(&User{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该邮箱已存在！")
	}
	if err := db.Where(user.Username).Find(&User{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户名已存在！")
	}
	return
}
