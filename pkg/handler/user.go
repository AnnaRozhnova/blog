package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}