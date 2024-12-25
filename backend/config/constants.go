package config

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ( // Errors
	ErrRecordNotFound    = gorm.ErrRecordNotFound
	ErrInvalidToken      = errors.New("invalid token")
	ErrUserRegistered    = errors.New("user already registered")
	ErrInvalidID         = errors.New("invalid id")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPwd        = errors.New("invalid password")
	ErrInvalidOrderOwner = errors.New("invalid order owner")
)

type ServiceCode int

const (
	CodeOK                  ServiceCode = 0
	CodeBadRequest          ServiceCode = 1
	CodeUnprocessableEntity ServiceCode = 2
	CodeUnauthorized        ServiceCode = 3
	CodeDatabaseError       ServiceCode = 4
	CodeNotFound            ServiceCode = 5
	CodeConflict            ServiceCode = 6
	CodeForbidden           ServiceCode = 7
)

const ( // Response Messages
	MsgCreateOK = "create success"
	MsgUpdateOK = "update success"
	MsgDeleteOK = "delete success"
)

const ( // Time periods
	JWTExpireTime = time.Hour * 24 * 30 // 30 days
)
