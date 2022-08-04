package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAllUsers writes list of all users to JSON body
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// getUserByUsername writes user's info to JSON body
func (h *Handler) getUserByUsername(c *gin.Context) {
	// Param returns the value of the URL param
	username := c.Param("username")
	user, err := h.service.User.GetByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// password is not required
	user.Password = ""
	c.JSON(http.StatusOK, user)
}