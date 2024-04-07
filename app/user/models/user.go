package models

import (
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Username string
	Password string
	Name     string
	Age      uint
}

func (user *User) CreateNewUserModel(dto *dto.UserRegisterDto) *User {
	user.Username = dto.Username
	user.Password = dto.Password
	user.Name = dto.Name
	user.Email = dto.Email
	user.Age = dto.Age
	return user
}
