package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	accountContextName  = "account-context"
	accountTokenName    = "account-token"
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	email, err := h.services.User.ParseToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(accountContextName, email)
	c.Set(accountTokenName, header)
}

func (h *Handler) GetUser(c *gin.Context) string {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return ""
	}
	email, err := h.services.User.ParseToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return ""
	}
	return email
}
