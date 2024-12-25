package utils

import (
	"backend/config"
	"backend/internal/controllers/http/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParseUintParam is used for parsing unsinged integer parameters from a request.
//
// Parameters:
//   - param: string which needs to be parsed to uint (e.g. "1")
//   - ctx: gin context, which will be aborted if an error happened
func ParseUintParam(param string, ctx *gin.Context) uint {
	val, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErr(config.CodeBadRequest, err.Error()))
		return 0
	}
	return uint(val)
}
