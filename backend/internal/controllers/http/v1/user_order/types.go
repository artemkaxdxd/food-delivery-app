package userorder

import (
	"backend/config"
	"backend/internal/controllers/http/middleware"
	request "backend/internal/controllers/http/request/user_order"
	response "backend/internal/controllers/http/response/user_order"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"

	"github.com/gin-gonic/gin"
)

type (
	service interface {
		GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) ([]response.UserOrder, config.ServiceCode, error)
		CreateOrder(ctx context.Context, userID uint, body request.UserOrder) (response.UserOrder, config.ServiceCode, error)
		UpdateOrder(ctx context.Context, userID, orderID uint, body request.UserOrder) (config.ServiceCode, error)
		DeleteOrder(ctx context.Context, userID, orderID uint) (config.ServiceCode, error)
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

	orders := g.Group("users/:user_id/orders")
	{
		orders.GET("", handler.getOrders)
		orders.POST("", handler.createOrder)
		orders.PATCH("/:order_id", handler.updateOrder)
		orders.DELETE("/:order_id", handler.deleteOrder)
	}
}
