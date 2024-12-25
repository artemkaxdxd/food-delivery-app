package user

import (
	"backend/config"
	"backend/pkg/utils"

	request "backend/internal/controllers/http/request/user"
	response "backend/internal/controllers/http/response/user"
	"context"
)

func (s service) Signup(ctx context.Context, body request.Signup) (string, config.ServiceCode, error) {
	user, err := s.repo.GetUserByEmail(ctx, body.Email)
	if err != nil {
		return "", config.DBErrToServiceCode(err), err
	}

	if user.ID != 0 {
		return "", config.CodeBadRequest, config.ErrUserRegistered
	}

	user = body.ToEntityUser()
	pwdHash, err := utils.HashPassword([]byte(body.Password))
	if err != nil {
		return "", config.CodeUnprocessableEntity, err
	}

	user.Password = pwdHash

	user, err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return "", config.DBErrToServiceCode(err), err
	}
	return s.getToken(user.ID)
}

func (s service) Login(ctx context.Context, body request.Login) (string, config.ServiceCode, error) {
	user, err := s.repo.GetUserByEmail(ctx, body.Email)
	if err != nil {
		return "", config.DBErrToServiceCode(err), err
	}

	if user.ID == 0 {
		return "", config.CodeNotFound, config.ErrInvalidEmail
	}

	if valid := utils.ComparePasswords(user.Password, body.Password); !valid {
		return "", config.CodeForbidden, config.ErrInvalidPwd
	}

	return s.getToken(user.ID)
}

func (s service) GetUserByID(ctx context.Context, userID uint) (response.User, config.ServiceCode, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if user.ID == 0 {
		return response.User{}, config.CodeNotFound, err
	}
	return response.UserToResponse(user), config.DBErrToServiceCode(err), err
}

func (s service) UpdateUser(ctx context.Context, userID uint, body request.UpdateUser) (config.ServiceCode, error) {
	err := s.repo.UpdateUser(ctx, body.ToEntity(userID))
	return config.DBErrToServiceCode(err), err
}

func (s service) DeleteUser(ctx context.Context, userID uint) (config.ServiceCode, error) {
	err := s.repo.DeleteUser(ctx, userID)
	if err != nil {
		return config.DBErrToServiceCode(err), err
	}

	err = s.orderRepo.DeleteOrdersByUserID(ctx, userID)
	return config.DBErrToServiceCode(err), err
}

func (s service) getToken(id uint) (string, config.ServiceCode, error) {
	accessToken, err := utils.GenerateAPIToken(id, s.cfgJWT.Secret, config.JWTExpireTime)
	if err != nil {
		return "", config.CodeUnprocessableEntity, err
	}
	return accessToken, config.CodeOK, nil
}
