package user

import (
	"backend/config"
	"backend/internal/controllers/http/middleware"
	request "backend/internal/controllers/http/request/user"
	response "backend/internal/controllers/http/response/user"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"

	"github.com/gin-gonic/gin"
)

type (
	service interface {
		Signup(ctx context.Context, body request.Signup) (string, config.ServiceCode, error)
		Login(ctx context.Context, body request.Login) (string, config.ServiceCode, error)

		GetUserByID(ctx context.Context, userID uint) (response.User, config.ServiceCode, error)
		UpdateUser(ctx context.Context, userID uint, body request.UpdateUser) (config.ServiceCode, error)
		DeleteUser(ctx context.Context, userID uint) (config.ServiceCode, error)
	}

	handler struct {
		l         logger.Logger
		svc       service
		validator validator.Validator
	}
)

func InitHandler(
	g *gin.Engine,
	l logger.Logger,
	svc service,
	validator validator.Validator,
	middle middleware.Middleware,
) {
	handler := handler{l, svc, validator}

	users := g.Group("users")
	{
		users.POST("/signup", handler.signup)
		users.POST("/login", handler.login)

		users.Use(middle.HandleUser())
		users.GET("/:user_id", handler.getUserByID)
		users.PATCH("/:user_id", handler.updateUser)
		users.DELETE("/:user_id", handler.deleteUser)
	}
}
