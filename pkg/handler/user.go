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


func (h *Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.service.User.GetByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}