package handler

import (
	"net/http"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
)


type getAllUsersResponse struct {
	Data []blog.User `json:"data"`
}
// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllUsersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUsersResponse{Data: users,})
}

// @Summary Get User By Username
// @Security ApiKeyAuth
// @Tags users
// @Description get user by username
// @ID get-user-by-username
// @Accept  json
// @Produce  json
// @Param        username   path      string  true  "username"
// @Success 200 {object} blog.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{username} [get]
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