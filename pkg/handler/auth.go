package handler

import (
	"net/http"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func (h *Handler) signUp(c *gin.Context) {
	var user blog.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	session, _ := store.Get(c.Request, "session")
	session.Values[usernameCtx] = user.Username
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"username": user.Username})
}


func (h *Handler) signIn(c *gin.Context) {
	var input blog.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.Authorization.GetUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	session, _ := store.Get(c.Request, "session")
	session.Values[usernameCtx] = user.Username
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"username": user.Username})
}

func(h *Handler) signOut(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	session.Options.MaxAge = -1
	//delete(session.Values, usernameCtx)
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"isOk": session.Options.MaxAge})
}