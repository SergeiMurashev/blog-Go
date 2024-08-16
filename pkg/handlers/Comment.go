package handler

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var input models.CommentInputCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	comment, err := h.services.Comment.CreateComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": comment,
	})
}

func (h *Handler) updateComment(c *gin.Context) {
	var input models.CommentInputUpdate
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	output, err := h.services.Comment.UpdateComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": output,
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	var input models.CommentInputDelete

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id

	err = h.services.Comment.DeleteComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{
		Status: "ok",
	})
}
