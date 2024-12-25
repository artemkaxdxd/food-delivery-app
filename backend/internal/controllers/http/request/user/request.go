package user

import (
	"backend/internal/entity/user"
	"time"
)

type (
	Signup struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,stringlength(8|16)"`
		Name     string `json:"name" valid:"required,stringlength(3|16)"`
	}

	Login struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,stringlength(8|16)"`
	}

	UpdateUser struct {
		Name string `json:"name" valid:"required,stringlength(3|10)"`
	}
)

func (l Signup) ToEntityUser() user.User {
	return user.User{
		Email:     l.Email,
		Password:  l.Password,
		Name:      l.Name,
		CreatedAt: time.Now().Unix(),
	}
}

func (u UpdateUser) ToEntity(userID uint) user.User {
	return user.User{
		ID:   userID,
		Name: u.Name,
	}
}
