package item

import (
	"backend/config"
	"backend/internal/controllers/http/middleware"
	request "backend/internal/controllers/http/request/item"
	response "backend/internal/controllers/http/response/item"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"

	"github.com/gin-gonic/gin"
)

type (
	service interface {
		GetItems(ctx context.Context, search string, limit, offset uint) ([]response.Item, config.ServiceCode, error)
		CreateItem(ctx context.Context, body request.Item) (response.Item, config.ServiceCode, error)
		UpdateItem(ctx context.Context, itemID uint, body request.Item) (config.ServiceCode, error)
		DeleteItem(ctx context.Context, itemID uint) (config.ServiceCode, error)
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

	items := g.Group("items")
	{
		items.GET("", handler.getItems)
		items.POST("", handler.createItem)
		items.PATCH("/:item_id", handler.updateItem)
		items.DELETE("/:item_id", handler.deleteItem)
	}
}
