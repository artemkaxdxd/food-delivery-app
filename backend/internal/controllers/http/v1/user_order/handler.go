package userorder

import (
	"backend/config"
	"backend/internal/controllers/http/middleware"
	userorder "backend/internal/controllers/http/request/user_order"
	"backend/internal/controllers/http/response"
	"backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler) parseLimitOffset(c *gin.Context) (limit, offset uint) {
	limitQuery := c.DefaultQuery("limit", "10")
	limit = utils.ParseUintParam(limitQuery, c)
	if c.IsAborted() {
		return
	}

	offsetQuery := c.DefaultQuery("offset", "0")
	offset = utils.ParseUintParam(offsetQuery, c)
	return
}

func (h handler) getOrders(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())
	limit, offset := h.parseLimitOffset(c)
	if c.IsAborted() {
		return
	}

	orders, svcCode, err := h.svc.GetOrdersByUserID(c.Request.Context(), claims.ID, limit, offset)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("orders", orders))
}

func (h handler) createOrder(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())

	var body userorder.Order
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	order, svcCode, err := h.svc.CreateOrder(c.Request.Context(), claims.ID, body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("order", order))
}

func (h handler) updateOrder(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())
	orderID := utils.ParseUintParam(c.Param("order_id"), c)
	if c.IsAborted() {
		return
	}

	var body userorder.Order
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	svcCode, err := h.svc.UpdateOrder(c.Request.Context(), claims.ID, orderID, body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgUpdateOK))
}

func (h handler) updateOrderStatus(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())
	orderID := utils.ParseUintParam(c.Param("order_id"), c)
	if c.IsAborted() {
		return
	}

	var body userorder.OrderStatus
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	svcCode, err := h.svc.UpdateOrderStatus(c.Request.Context(), claims.ID, orderID, body.Status)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgUpdateOK))
}

func (h handler) deleteOrder(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())
	orderID := utils.ParseUintParam(c.Param("order_id"), c)
	if c.IsAborted() {
		return
	}

	svcCode, err := h.svc.DeleteOrder(c.Request.Context(), claims.ID, orderID)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgDeleteOK))
}
