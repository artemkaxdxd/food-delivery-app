package item

import (
	"backend/config"
	"backend/internal/controllers/http/request/item"
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

func (h handler) getItems(c *gin.Context) {
	search := c.Query("q")
	limit, offset := h.parseLimitOffset(c)
	if c.IsAborted() {
		return
	}

	items, svcCode, err := h.svc.GetItems(c.Request.Context(), search, limit, offset)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("items", items))
}

func (h handler) createItem(c *gin.Context) {
	var body item.Item
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	item, svcCode, err := h.svc.CreateItem(c.Request.Context(), body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("item", item))
}

func (h handler) updateItem(c *gin.Context) {
	itemID := utils.ParseUintParam(c.Param("item_id"), c)
	if c.IsAborted() {
		return
	}

	var body item.Item
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	svcCode, err := h.svc.UpdateItem(c.Request.Context(), itemID, body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgUpdateOK))
}

func (h handler) deleteItem(c *gin.Context) {
	itemID := utils.ParseUintParam(c.Param("item_id"), c)
	if c.IsAborted() {
		return
	}

	svcCode, err := h.svc.DeleteItem(c.Request.Context(), itemID)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgDeleteOK))
}
