package user

import (
	"backend/config"
	"backend/internal/controllers/http/middleware"
	"backend/internal/controllers/http/request/user"
	"backend/internal/controllers/http/response"
	"backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) signup(c *gin.Context) {
	var body user.Signup
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	accessToken, svcCode, err := h.svc.Signup(c.Request.Context(), body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("token", accessToken))
}

func (h handler) login(c *gin.Context) {
	var body user.Login
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	accessToken, svcCode, err := h.svc.Login(c.Request.Context(), body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("token", accessToken))
}

func (h handler) getUserByID(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())

	userID := utils.ParseUintParam(c.Param("user_id"), c)
	if c.IsAborted() {
		return
	}

	if claims.ID != userID {
		c.JSON(http.StatusForbidden, response.NewErr(config.CodeForbidden, config.ErrInvalidID.Error()))
		return
	}

	accessToken, svcCode, err := h.svc.GetUserByID(c.Request.Context(), claims.ID)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).AddKey("token", accessToken))
}

func (h handler) updateUser(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())

	var body user.UpdateUser
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	if _, err := h.validator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return
	}

	userID := utils.ParseUintParam(c.Param("user_id"), c)
	if c.IsAborted() {
		return
	}

	if claims.ID != userID {
		c.JSON(http.StatusForbidden, response.NewErr(config.CodeForbidden, config.ErrInvalidID.Error()))
		return
	}

	svcCode, err := h.svc.UpdateUser(c.Request.Context(), claims.ID, body)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgUpdateOK))
}

func (h handler) deleteUser(c *gin.Context) {
	claims := middleware.GetClaims(c.Request.Context())

	userID := utils.ParseUintParam(c.Param("user_id"), c)
	if c.IsAborted() {
		return
	}

	if claims.ID != userID {
		c.JSON(http.StatusForbidden, response.NewErr(config.CodeForbidden, config.ErrInvalidID.Error()))
		return
	}

	svcCode, err := h.svc.DeleteUser(c.Request.Context(), claims.ID)
	if err != nil {
		c.JSON(config.ServiceCodeToHttpStatus(svcCode), response.NewErr(svcCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.New(svcCode).SetDescription(config.MsgDeleteOK))
}
