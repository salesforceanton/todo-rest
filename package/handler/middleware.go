package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (h *Handler) userIdentitity(c *gin.Context) {
	authHeader := c.GetHeader(AuthorizationHeader)

	if authHeader == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Authorization header is empty")
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "Authorization header is invalid")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(UserCtx, userId)
}

func (h *Handler) getUserId(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get(UserCtx)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("User id is not found")
	}

	return userId.(int), nil
}
